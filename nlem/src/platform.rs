use super::{security, service, storage};

#[tonic::async_trait]
pub trait Platform: Send + Sync {
    async fn hostname(&self) -> Result<String, String>;
    async fn device_type(&self) -> storage::DeviceType;
    async fn discovery_manager(&self) -> Result<crate::DiscoveryC, String>;
    async fn storage(&self) -> Result<crate::StorageC, String>;

    async fn prompt_pair_request(
        &self,
        info: storage::DeviceInfo,
        certificate: security::Certificate,
    ) -> Result<(), String>;

    async fn log_info(&self, msg: String);
    async fn log_error(&self, msg: String);
    async fn log_warning(&self, msg: String);

    async fn query_telemetry(&self) -> Result<service::telemetry::TelemetryInfo, String>;
}
