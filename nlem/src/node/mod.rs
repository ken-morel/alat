use std::sync::Arc;

use crate::{
    devicemanager::{DeviceManagerEvent, connected},
    discovery, unit,
};

use super::{client, devicemanager, security, server, storage};

/// A node contains a collection of references to it's components.
/// The node stores no data on it's own, that's why it can easily be cloned
#[derive(Clone, Debug)]
pub struct Node {
    pub storage: crate::StorageC,
    pub platform: crate::PlatformC,
    pub device_manager: crate::DeviceManagerC,
    pub unit_manager: crate::UnitManagerC,
    pub server: crate::ServerC,
}

impl Node {
    pub async fn init(platform: crate::PlatformC) -> Result<Self, crate::ErrorC> {
        let storage = platform.read().await.storage().await?;
        storage
            .lock()
            .await
            .init(storage::StorageData {
                certificate: security::generate_certificate(),
                paired_devices: Vec::new(),
                info: Self::default_device_info(&*platform.read().await).await,
                settings: std::collections::BTreeMap::new(),
            })
            .await?;
        let discovery = platform.write().await.discovery_manager().await?;
        let device_manager = Arc::new(
            devicemanager::DeviceManager::init(
                storage.clone(),
                platform.clone(),
                discovery.clone(),
            )
            .await?,
        );
        let unit_manager = Arc::new(unit::UnitManager::new());

        let server = Arc::new(server::Server::new(
            device_manager.clone(),
            unit_manager.clone(),
        ));
        Ok(Self {
            storage,
            platform,
            device_manager,
            unit_manager,
            server,
        })
    }

    pub async fn default_device_info(p: &crate::Platform) -> storage::DeviceInfo {
        storage::DeviceInfo {
            id: security::generate_id(),
            color: storage::Color::random(),
            name: p.hostname().await.expect("Could not get hostname"),
            device_type: p.device_type().await,
        }
    }
    pub async fn start(&self) -> Result<tokio::sync::mpsc::Receiver<NodeEvent>, crate::ErrorC> {
        self.unit_manager.init().await?;
        let router = self.server.create_router().await?;
        tokio::spawn(async move {
            let addr = std::net::SocketAddr::new(
                std::net::Ipv4Addr::UNSPECIFIED.into(),
                server::ALAT_PORT,
            );
            println!("[node/server] Starting server at {addr}");
            let r = router.serve(addr).await;
            println!("[node/server] Server at {addr} stopped");
            if let Err(e) = r {
                println!("[node/server::error] {e}");
            }
        });

        let managerrx = self.device_manager.start_workers().await;
        let (tx, rx) = tokio::sync::mpsc::channel(1);
        tokio::spawn(node_worker(self.clone(), tx, managerrx));
        Ok(rx)
    }

    pub async fn request_pair(
        &self,
        device_id: &security::DeviceID,
    ) -> Result<Result<storage::PairedDevice, String>, String> {
        let manager = self.device_manager.clone();
        let device = manager
            .discovered_devices
            .get(device_id)
            .ok_or(String::from("Device not found"))?
            .clone();
        let this_info = manager.this_device.read().await.clone().info;
        let this_certificate = manager.device_certificate.read().await.clone();

        let mut cl = client::Client::connect(device.address).await.map_err(|e| {
            format!(
                "Client could not connect to device at {0}: {e}",
                device.address
            )
        })?;

        match cl.request_pair(this_info, this_certificate).await {
            Ok(response) => match response {
                Ok((token, certificate, info)) => {
                    let paired_device = storage::PairedDevice {
                        token,
                        certificate,
                        info,
                    };
                    self.device_manager
                        .add_paired_device(paired_device.clone())
                        .await;
                    Ok(Ok(paired_device))
                }
                Err(message) => Ok(Err(message)),
            },
            Err(err) => Err(format!("Could not send pair request: {err}")),
        }
    }
    pub async fn unit_register(&self, unit: crate::UnitC) {
        self.unit_manager.add_unit(unit).await;
    }
}

#[derive(Debug)]
pub enum NodeEvent {
    DiscoveryStarted,
    DeviceDiscovered(Box<discovery::DiscoveredDevice>),
    DeviceLost(security::DeviceID),
    DiscoveryError(discovery::DiscoveryError),
    DiscoveryStopped,

    DeviceConnected(Box<connected::ConnectedDevice>),
    DeviceDisconnected(security::DeviceID),
    ConnectionError(String),

    DevicePaired(storage::PairedDevice),
    DeviceUnpaired(security::DeviceID),

    DiscoveryServerStarted(Box<discovery::DiscoveredDevice>),
    DiscoveryServerUpdated(Box<discovery::DiscoveredDevice>),
    DiscoveryServerStopped,
    DiscoveryServerError(crate::ErrorC),

    DeviceManagerStarted,
    DeviceManagerStopped,

    NodeStarted,
    NodeStopped,
}

async fn node_worker(
    node: Node,
    events: tokio::sync::mpsc::Sender<NodeEvent>,
    mut manager_events: tokio::sync::mpsc::Receiver<devicemanager::DeviceManagerEvent>,
) {
    let (tx, mut rx) = tokio::sync::mpsc::channel(1);
    
    tokio::spawn(async move {
        while let Some(event) = rx.recv().await {
            events.send(event).await.ok();
        }
    });

    let send = |event| {
        let tx = tx.clone();
        async move {
            tx.send(event).await.ok();
        }
    };

    send(NodeEvent::NodeStarted).await;
    while let Some(event) = manager_events.recv().await {
        match event as DeviceManagerEvent {
            DeviceManagerEvent::Started => send(NodeEvent::DeviceManagerStarted).await,
            DeviceManagerEvent::Stopped => send(NodeEvent::DeviceManagerStopped).await,
            DeviceManagerEvent::Found(device) => {
                send(NodeEvent::DeviceDiscovered(device)).await;
            }
            DeviceManagerEvent::Lost(device_id) => {
                send(NodeEvent::DeviceLost(device_id)).await;
            }
            DeviceManagerEvent::DiscoveryError(err) => {
                send(NodeEvent::DiscoveryError(err)).await;
            }
            DeviceManagerEvent::Connected(device) => {
                send(NodeEvent::DeviceConnected(device)).await;
            }
            DeviceManagerEvent::Disconnected(device_id) => {
                send(NodeEvent::DeviceDisconnected(device_id)).await;
            }
            DeviceManagerEvent::ConnectionError(msg) => {
                send(NodeEvent::ConnectionError(msg)).await;
            }
            DeviceManagerEvent::Paired(device) => {
                send(NodeEvent::DevicePaired(device)).await;
            }
            DeviceManagerEvent::Unpaired(device_id) => {
                send(NodeEvent::DeviceUnpaired(device_id)).await;
            }
            DeviceManagerEvent::DiscoveryServerStarted(info) => {
                send(NodeEvent::DiscoveryServerStarted(info)).await;
            }
            DeviceManagerEvent::DiscoveryServerStartedUpdated(info) => {
                send(NodeEvent::DiscoveryServerUpdated(info)).await;
            }
            DeviceManagerEvent::DiscoveryServerStopped => {
                send(NodeEvent::DiscoveryServerStopped).await;
            }
            DeviceManagerEvent::DiscoveryStarted => {
                send(NodeEvent::DiscoveryStarted).await;
            }
            DeviceManagerEvent::DiscoveryStopped => {
                send(NodeEvent::DiscoveryStopped).await;
            }
            DeviceManagerEvent::InfoLog(msg) => {
                let p = node.platform.read().await;
                p.log_info(msg).await
            }
            DeviceManagerEvent::WarningLog(msg) => {
                let p = node.platform.read().await;
                p.log_warning(msg).await
            }
            DeviceManagerEvent::ErrorLog(msg) => {
                let p = node.platform.read().await;
                p.log_error(msg).await
            }
        }
    }
    send(NodeEvent::NodeStopped).await;
}
