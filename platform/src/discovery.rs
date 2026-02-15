use log::{error, warn};
use nlem::devicemanager::discovered::{self, DiscoveredDevice, DiscoveryError, DiscoveryEvent};
use nlem::security::DeviceID;
use nlem::storage::DeviceInfo;
use std::collections::HashMap;
use std::sync::Arc;
use tokio::net::UdpSocket;
use tokio::sync::{Mutex, mpsc};
use tokio::task::JoinHandle;
use tokio::time::{self, Duration};

const DISCOVERY_PORT: u16 = 4147;
const BROADCAST_ADDR: &str = "255.255.255.255";
const ADVERTISEMENT_INTERVAL: Duration = Duration::from_secs(5);
const DEVICE_TIMEOUT: Duration = Duration::from_secs(15);
const BROADCAST_DATA_BUFFER_SIZE: usize = 256;

pub struct DiscoveryManager {
    advertising_task: Option<JoinHandle<()>>,
    scan_task: Option<JoinHandle<()>>,
    advertised_device_id: Arc<Mutex<Option<DeviceID>>>,
}

impl DiscoveryManager {
    pub async fn init() -> Result<Self, String> {
        Ok(Self {
            advertising_task: None,
            scan_task: None,
            advertised_device_id: Arc::new(Mutex::new(None)),
        })
    }
}

async fn run_advertiser(device_info: DeviceInfo, socket: Arc<UdpSocket>) {
    let mut interval = time::interval(ADVERTISEMENT_INTERVAL);
    let broadcast_addr = format!("{}:{}", BROADCAST_ADDR, DISCOVERY_PORT);

    // usuall about 130B
    let message_bytes = match serde_json::to_vec(&device_info) {
        Ok(bytes) => bytes,
        Err(e) => {
            error!("Failed to serialize DeviceInfo for advertising: {}", e);
            return;
        }
    };

    log::info!(
        "[platform/discovery.rs] Starting to advertise device '{}' on port {}",
        device_info.name,
        DISCOVERY_PORT
    );

    loop {
        interval.tick().await;
        if let Err(e) = socket.send_to(&message_bytes, &broadcast_addr).await {
            warn!("Failed to send advertisement broadcast: {}", e);
        }
    }
}

async fn run_scanner(
    own_device_id_arc: Arc<Mutex<Option<DeviceID>>>,
    socket: Arc<UdpSocket>,
    sender: mpsc::Sender<DiscoveryEvent>,
) {
    let mut discovered_devices: HashMap<DeviceID, (DiscoveredDevice, time::Instant)> =
        HashMap::new();
    let mut buffer = [0u8; BROADCAST_DATA_BUFFER_SIZE];
    let mut cleanup_interval = time::interval(DEVICE_TIMEOUT / 2);

    println!("Starting to scan for devices on port {}", DISCOVERY_PORT);

    loop {
        tokio::select! {
            Ok((len, remote_addr)) = socket.recv_from(&mut buffer) => {
                match serde_json::from_slice::<DeviceInfo>(&buffer[..len]) {
                    Ok(info) => {
                        let own_id_lock = own_device_id_arc.lock().await;
                        if let Some(own_id) = &*own_id_lock
                            && &info.id == own_id {
                                //DEBUG: Let's just consider it for tests
                                // continue;
                            }

                        let discovered_device = DiscoveredDevice { address: remote_addr, info };

                        if !discovered_devices.contains_key(&discovered_device.info.id)
                             && sender.send(DiscoveryEvent::Found(discovered_device.clone())).await.is_err() {
                                error!("Failed to send DiscoveryEvent::Found. Receiver closed.");
                                break;
                             }
                        discovered_devices.insert(discovered_device.info.id, (discovered_device, time::Instant::now()));
                    }
                    Err(e) => {
                        warn!("Failed to deserialize discovery packet from {}: {}", remote_addr, e);
                    }
                }
            }

            _ = cleanup_interval.tick() => {
                let now = time::Instant::now();
                let mut lost_devices = Vec::new();

                discovered_devices.retain(|_id, (device, last_seen)| {
                    if now.duration_since(*last_seen) > DEVICE_TIMEOUT {
                        lost_devices.push(device.info.id);
                        false
                    } else {
                        true
                    }
                });

                for device_id in lost_devices {
                    if sender.send(DiscoveryEvent::Lost(device_id)).await.is_err() {
                        error!("Failed to send DiscoveryEvent::Lost. Receiver closed.");
                        break;
                    }
                }
            }
        }
    }
    log::warn!("Scanner task stopped.");
}

impl discovered::DiscoveryManager for DiscoveryManager {
    async fn advertise(&mut self, device: DiscoveredDevice) -> Result<(), DiscoveryError> {
        if self.advertising_task.is_some() {
            return Err(DiscoveryError::AdvertiseError(
                "Advertising already in progress".to_string(),
            ));
        }

        let socket = UdpSocket::bind("0.0.0.0:0").await.map_err(|e| {
            DiscoveryError::AdvertiseError(format!("Failed to bind UDP socket: {}", e))
        })?;
        socket.set_broadcast(true).map_err(|e| {
            DiscoveryError::AdvertiseError(format!("Failed to set broadcast on UDP socket: {}", e))
        })?;
        let socket = Arc::new(socket);

        let device_info = device.info;
        *self.advertised_device_id.lock().await = Some(device_info.id);

        let handle = tokio::spawn(async move {
            run_advertiser(device_info, socket).await;
        });

        self.advertising_task = Some(handle);
        Ok(())
    }

    async fn cease_advertising(&mut self) -> Result<(), DiscoveryError> {
        if let Some(handle) = self.advertising_task.take() {
            log::info!("Ceasing advertising.");
            handle.abort();
            *self.advertised_device_id.lock().await = None;
        }
        Ok(())
    }

    async fn scan(
        &mut self,
        sender: tokio::sync::mpsc::Sender<DiscoveryEvent>,
    ) -> Result<(), DiscoveryError> {
        if self.scan_task.is_some() {
            return Err(DiscoveryError::ScanError(
                "Scan already in progress".to_string(),
            ));
        }

        let listen_addr = format!("0.0.0.0:{}", DISCOVERY_PORT);
        let socket = UdpSocket::bind(&listen_addr)
            .await
            .map_err(|e| DiscoveryError::ScanError(format!("Failed to bind UDP socket: {}", e)))?;
        let socket = Arc::new(socket);

        let own_device_id_arc = self.advertised_device_id.clone();

        let handle = tokio::spawn(async move {
            run_scanner(own_device_id_arc, socket, sender).await;
        });

        self.scan_task = Some(handle);
        Ok(())
    }

    async fn cease_scan(&mut self) -> Result<(), DiscoveryError> {
        if let Some(handle) = self.scan_task.take() {
            println!("Ceasing scan.");
            handle.abort();
        }
        Ok(())
    }
}
