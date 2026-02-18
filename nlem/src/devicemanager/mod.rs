pub mod connected;
pub mod discovered;
mod workers;

use tokio::sync::RwLock;

use crate::{server, storage::StorageError};

use super::{platform, security, storage};
use std::{
    collections::HashMap,
    net::{Ipv4Addr, SocketAddr},
    sync::Arc,
};

#[derive(Debug, Clone)]
pub enum DeviceManagerEvent {
    Found(Box<discovered::DiscoveredDevice>),
    Lost(security::DeviceID),

    Connected(Box<connected::ConnectedDevice>),
    Disconnected(security::DeviceID),
    ConnectionError(String),

    Paired(storage::PairedDevice),
    Unpaired(security::DeviceID),

    DiscoveryServerStarted(Box<discovered::DiscoveredDevice>),
    DiscoveryServerUpdated(Box<discovered::DiscoveredDevice>),
    DiscoveryServerStopped,

    DiscoveryStarted,
    DiscoveryError(discovered::DiscoveryError),
    DiscoveryStopped,

    InfoLog(String),
    WarningLog(String),

    Started,
    Stopped,
}

#[derive(Debug)]
pub struct DeviceManager<
    S: storage::Storage,
    P: platform::Platform<S, D>,
    D: discovered::DiscoveryManager,
> {
    pub storage: Arc<RwLock<S>>,
    pub platform: Arc<RwLock<P>>,

    pub paired_devices: Arc<RwLock<HashMap<security::DeviceID, storage::PairedDevice>>>,
    pub this_device: Arc<RwLock<discovered::DiscoveredDevice>>,
    pub device_certificate: Arc<RwLock<security::Certificate>>,

    pub connected_devices: Arc<RwLock<HashMap<security::DeviceID, connected::ConnectedDevice>>>,

    pub discovered_devices: Arc<RwLock<HashMap<security::DeviceID, discovered::DiscoveredDevice>>>,
    discovery: Arc<RwLock<D>>,
}

impl<S: storage::Storage, P: platform::Platform<S, D>, D: discovered::DiscoveryManager>
    DeviceManager<S, P, D>
{
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
        store: Arc<RwLock<S>>,
        platform: Arc<RwLock<P>>,
        discovery: Arc<RwLock<D>>,
    ) -> Result<Self, StorageError> {
        let data = store
            .write()
            .await
            .init(storage::StorageData {
                certificate: security::generate_certificate(),
                paired_devices: Vec::new(),
                info: Self::default_device_info(&platform).await,
            })
            .await?;
        let mut paired_devices = HashMap::new();
        for paired_device in data.paired_devices {
            paired_devices.insert(paired_device.info.id, paired_device);
        }

        Ok(Self {
            this_device: Arc::new(RwLock::new(discovered::DiscoveredDevice {
                address: SocketAddr::new(Ipv4Addr::LOCALHOST.into(), server::ALAT_PORT),
                info: data.info,
            })),
            storage: store.clone(),
            paired_devices: Arc::new(RwLock::new(paired_devices)),
            device_certificate: Arc::new(RwLock::new(data.certificate)),
            connected_devices: Arc::new(RwLock::new(HashMap::new())),
            discovered_devices: Arc::new(RwLock::new(HashMap::new())),
            discovery,
            platform,
        })
    }
    pub async fn default_device_info(p: &Arc<RwLock<P>>) -> storage::DeviceInfo {
        let lck = p.read().await;
        storage::DeviceInfo {
            id: security::generate_id(),
            color: storage::Color::random(),
            name: lck.hostname().expect("Could not get hostname"),
            device_type: lck.device_type(),
        }
    }

    pub async fn add_paired_device(&self, device: storage::PairedDevice) {
        self.paired_devices
            .write()
            .await
            .insert(device.info.id, device);
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
