use super::{devicemanager, storage};

pub trait Platform<D: devicemanager::discovered::DiscoveryManager>: Sized + Send + Sync {
    fn hostname(&self) -> Result<String, String>;
    fn device_type(&self) -> storage::DeviceType;
    fn discovery_manager(&self) -> impl std::future::Future<Output = Result<D, String>> + Send;
}
