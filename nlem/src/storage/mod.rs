use super::security::Certificate;

mod error;
pub type StorageError = error::StorageError;

mod color;
pub type Color = color::Color;

mod deviceinfo;
pub type DeviceInfo = deviceinfo::DeviceInfo;

mod devicetype;
pub type DeviceType = devicetype::DeviceType;

mod paireddevice;
pub type PairedDevice = paireddevice::PairedDevice;

#[derive(Debug, Clone, serde::Deserialize, serde::Serialize)]
pub struct StorageData {
    pub certificate: Certificate,
    pub paired_devices: Vec<PairedDevice>,
    pub info: DeviceInfo,
}

use super::proto;

pub type StorageResult<T> = Result<T, StorageError>;

#[tonic::async_trait]
pub trait Storage: Send + Sync {
    async fn init(&self, data: StorageData) -> StorageResult<()>;

    async fn load_certificate(&self) -> StorageResult<Certificate>;
    async fn save_certificate(&self, certificate: Certificate) -> StorageResult<()>;

    async fn load_info(&self) -> StorageResult<DeviceInfo>;
    async fn save_info(&self, info: DeviceInfo) -> StorageResult<()>;

    async fn load_paired(&self) -> StorageResult<Vec<PairedDevice>>;
    async fn save_paired(&self, paired: Vec<PairedDevice>) -> StorageResult<()>;
    async fn add_paired(&self, dev: PairedDevice) -> StorageResult<()> {
        let mut data = self.load_paired().await?;
        data.push(dev);
        self.save_paired(data).await
    }
}
