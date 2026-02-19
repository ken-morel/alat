pub mod connected;
mod workers;

use tokio::sync::RwLock;

use crate::{discovery, server, storage::StorageError};

use super::{security, storage};
use std::{
    collections::HashMap,
    net::{Ipv4Addr, SocketAddr},
    sync::Arc,
};

#[derive(Debug, Clone)]
pub enum DeviceManagerEvent {
    Found(Box<discovery::DiscoveredDevice>),
    Lost(security::DeviceID),

    Connected(Box<connected::ConnectedDevice>),
    Disconnected(security::DeviceID),
    ConnectionError(String),

    Paired(storage::PairedDevice),
    Unpaired(security::DeviceID),

    DiscoveryServerStarted(Box<discovery::DiscoveredDevice>),
    DiscoveryServerUpdated(Box<discovery::DiscoveredDevice>),
    DiscoveryServerStopped,

    DiscoveryStarted,
    DiscoveryError(discovery::DiscoveryError),
    DiscoveryStopped,

    InfoLog(String),
    WarningLog(String),

    Started,
    Stopped,
}

#[derive()]
pub struct DeviceManager {
    pub storage: crate::StorageC,
    pub platform: crate::PlatformC,

    pub paired_devices: Arc<RwLock<HashMap<security::DeviceID, storage::PairedDevice>>>,
    pub this_device: Arc<RwLock<discovery::DiscoveredDevice>>,
    pub device_certificate: Arc<RwLock<security::Certificate>>,

    pub connected_devices: Arc<RwLock<HashMap<security::DeviceID, connected::ConnectedDevice>>>,

    pub discovered_devices: Arc<RwLock<HashMap<security::DeviceID, discovery::DiscoveredDevice>>>,

    worker: Option<tokio::sync::mpsc::Sender<workers::WorkerEvent>>,
    discovery: crate::DiscoveryC,
}

impl DeviceManager {
    async fn load(&mut self) -> Result<(), StorageError> {
        let store = self.storage.read().await;
        let mut map = HashMap::new();
        for device in store.load_paired().await? {
            map.insert(device.info.id, device);
        }
        *self.paired_devices.write().await = map;
        self.this_device.write().await.info = store.load_info().await?;
        *self.device_certificate.write().await = store.load_certificate().await?;
        Ok(())
    }

    async fn save(&self) -> Result<(), StorageError> {
        let store = self.storage.read().await;
        let paired_devices = self.paired_devices.read().await.values().cloned().collect();
        store.save_paired(paired_devices).await?;
        store
            .save_info(self.this_device.read().await.info.clone())
            .await?;
        store
            .save_certificate(self.device_certificate.read().await.clone())
            .await?;
        Ok(())
    }
    pub async fn init(
        store: crate::StorageC,
        platform: crate::PlatformC,
        discovery: crate::DiscoveryC,
    ) -> Result<Self, StorageError> {
        let mut manager = Self {
            this_device: Arc::new(RwLock::new(discovery::DiscoveredDevice {
                address: SocketAddr::new(Ipv4Addr::LOCALHOST.into(), server::ALAT_PORT),
                info: storage::DeviceInfo::default(),
            })),
            paired_devices: Arc::new(RwLock::new(HashMap::default())),
            device_certificate: Arc::new(RwLock::new(security::Certificate::default())),
            connected_devices: Arc::new(RwLock::new(HashMap::default())),
            discovered_devices: Arc::new(RwLock::new(HashMap::default())),
            worker: None,
            storage: store.clone(),
            discovery,
            platform,
        };
        manager.load().await?;
        Ok(manager)
    }

    pub async fn add_paired_device(&self, device: storage::PairedDevice) {
        self.paired_devices
            .write()
            .await
            .insert(device.info.id, device.clone());
        if let Some(worker) = self.worker.clone() {
            _ = worker
                .send(workers::WorkerEvent::Wrapper(DeviceManagerEvent::Paired(
                    device,
                )))
                .await;
        }

        if let Err(e) = self.save().await {
            println!("Error saving paired devices after add_paired_device: {e}");
        }
    }

    pub async fn _handle_pair_request(
        &self,
        info: storage::DeviceInfo,
        certificate: security::Certificate,
    ) -> Result<storage::PairedDevice, String> {
        self.platform
            .read()
            .await
            .prompt_pair_request(info.clone(), certificate.clone())
            .await?;
        let paired = storage::PairedDevice {
            token: security::generate_pair_token(),
            certificate,
            info,
        };
        self.add_paired_device(paired.clone()).await;
        Ok(paired)
    }
}
