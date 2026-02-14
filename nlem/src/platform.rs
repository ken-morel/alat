use super::{devicemanager, storage};

pub trait Platform<D: devicemanager::discovered::DiscoveryManager>: Sized + Send + Sync {
    fn hostname(&self) -> Result<String, String>;
    fn device_type(&self) -> storage::DeviceType;
    async fn discovery_manager(&self) -> Result<D, String>;
}
