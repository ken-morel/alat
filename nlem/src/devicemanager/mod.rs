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
    Found(discovered::DiscoveredDevice),
    Lost(security::DeviceID),

    Connected(connected::ConnectedDevice),
    Disconnected(security::DeviceID),
    ConnectionError(String),

    Paired(storage::PairedDevice),
    Unpaired(security::DeviceID),

    DiscoveryServerStarted(discovered::DiscoveredDevice),
    DiscoveryServerUpdated(discovered::DiscoveredDevice),
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
    P: platform::Platform,
    D: discovered::DiscoveryManager,
> {
    pub storage: Arc<RwLock<S>>,
    pub platform: Arc<RwLock<P>>,

    pub paired_devices: Arc<RwLock<HashMap<security::DeviceID, storage::PairedDevice>>>,
    pub this_device: Arc<RwLock<discovered::DiscoveredDevice>>,
    pub device_certificate: Arc<RwLock<security::Certificate>>,

    pub connected_devices: Arc<RwLock<HashMap<security::DeviceID, connected::ConnectedDevice>>>,

    pub discovered_devices: Arc<RwLock<HashMap<security::DeviceID, discovered::DiscoveredDevice>>>,
    pub discovery: Arc<RwLock<D>>,
}

impl<S: storage::Storage, P: platform::Platform, D: discovered::DiscoveryManager>
    DeviceManager<S, P, D>
{
    async fn new(
        store: Arc<RwLock<S>>,
        platform: Arc<RwLock<P>>,
        discovery: Arc<RwLock<D>>,
    ) -> Self {
        Self {
            this_device: Arc::new(RwLock::new(discovered::DiscoveredDevice {
                address: SocketAddr::new(Ipv4Addr::LOCALHOST.into(), server::ALAT_PORT),
                info: DeviceManager::<S, P, D>::default_device_info(&platform).await,
            })),
            storage: store,
            paired_devices: Arc::new(RwLock::new(HashMap::new())),
            device_certificate: Arc::new(RwLock::new(security::generate_certificate())),
            connected_devices: Arc::new(RwLock::new(HashMap::new())),
            discovered_devices: Arc::new(RwLock::new(HashMap::new())),
            discovery,
            platform,
        }
    }
    async fn load(&mut self) -> Result<(), StorageError> {
        let store = self.storage.read().await;
        let mut map = HashMap::new();
        for device in store.load_paired()? {
            map.insert(device.info.id.clone(), device);
        }
        std::mem::replace(&mut *self.paired_devices.write().await, map);
        std::mem::replace(
            &mut *&mut self.this_device.write().await.info,
            store.load_info()?,
        );
        std::mem::replace(
            &mut *self.device_certificate.write().await,
            store.load_certificate()?,
        );
        Ok(())
    }
    async fn save(&self) -> Result<(), StorageError> {
        let store = self.storage.read().await;
        let paired_devices = self
            .paired_devices
            .read()
            .await
            .values()
            .map(|d| d.clone())
            .collect();
        store.save_paired(paired_devices)?;
        store.save_info(self.this_device.read().await.info.clone())?;
        store.save_certificate(self.device_certificate.read().await.clone())?;
        Ok(())
    }
    pub async fn init(
        store: Arc<RwLock<S>>,
        platform: Arc<RwLock<P>>,
        discovery: Arc<RwLock<D>>,
    ) -> Result<Self, StorageError> {
        let mut manager = Self::new(store, platform, discovery).await;
        if manager.load().await.is_err() {
            manager.save().await?;
        }
        Ok(manager)
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
}
