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

pub trait DiscoveryManager: Sized + Send + Sync {
    fn advertise(
        &mut self,
        device: DiscoveredDevice,
    ) -> impl std::future::Future<Output = Result<(), DiscoveryError>> + std::marker::Send;
    fn cease_advertising(
        &mut self,
    ) -> impl std::future::Future<Output = Result<(), DiscoveryError>> + std::marker::Send;
    fn scan(
        &mut self,
        sender: tokio::sync::mpsc::Sender<DiscoveryEvent>,
    ) -> impl std::future::Future<Output = Result<(), DiscoveryError>> + std::marker::Send;
    fn cease_scan(
        &mut self,
    ) -> impl std::future::Future<Output = Result<(), DiscoveryError>> + std::marker::Send;
}
