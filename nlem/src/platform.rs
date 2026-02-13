use super::{devicemanager, storage};

pub trait Platform: Sized + Send + Sync {
    fn hostname(&self) -> Result<String, String>;
    fn device_type(&self) -> storage::DeviceType;
    fn discovery_manager(&self) -> impl devicemanager::discovered::DiscoveryManager;
}
