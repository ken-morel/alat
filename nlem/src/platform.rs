use super::{devicemanager, security, storage};

pub trait Platform<S: storage::Storage, D: devicemanager::discovered::DiscoveryManager>:
    Sized + Send + Sync
{
    fn hostname(&self) -> Result<String, String>;
    fn device_type(&self) -> storage::DeviceType;
    fn discovery_manager(&self) -> impl std::future::Future<Output = Result<D, String>> + Send;
    fn storage(&self) -> impl std::future::Future<Output = Result<S, String>> + Send;

    async fn prompt_pair_request(
        &self,
        info: storage::DeviceInfo,
        certificate: security::Certificate,
    ) -> Result<(), String>;
}
