use super::{security, storage};

#[derive(Debug, Clone, PartialEq, Eq, serde::Serialize, serde::Deserialize)]
pub struct DiscoveredDevice {
    pub address: std::net::SocketAddr,
    pub info: storage::DeviceInfo,
}

pub enum DiscoveryEvent {
    Found(DiscoveredDevice),
    Lost(security::DeviceID),
}

#[derive(Debug, Clone, thiserror::Error)]
pub enum DiscoveryError {
    #[error("Failed to start advertising: {0}")]
    AdvertiseError(String),
    #[error("Failed to cease advertising: {0}")]
    CeaseAdvertiseError(String),
    #[error("Failed to start scanning: {0}")]
    ScanError(String),
    #[error("Failed to stop scanning: {0}")]
    StopScanError(String),
    #[error("Platform-specific error: {0}")]
    PlatformSpecific(String),
}

#[tonic::async_trait]
pub trait DiscoveryManager: Send + Sync {
    async fn advertise(&mut self, device: DiscoveredDevice) -> Result<(), DiscoveryError>;
    async fn cease_advertising(&mut self) -> Result<(), DiscoveryError>;
    async fn scan(
        &mut self,
        sender: tokio::sync::mpsc::Sender<DiscoveryEvent>,
    ) -> Result<(), DiscoveryError>;
    async fn cease_scan(&mut self) -> Result<(), DiscoveryError>;
}
