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
    async fn init(&mut self, data: StorageData) -> StorageResult<()>;

    async fn get_certificate(&mut self) -> StorageResult<Certificate>;
    async fn set_certificate(&mut self, certificate: Certificate) -> StorageResult<()>;

    async fn get_info(&mut self) -> StorageResult<DeviceInfo>;
    async fn set_info(&mut self, info: DeviceInfo) -> StorageResult<()>;

    async fn get_paired(&mut self) -> StorageResult<Vec<PairedDevice>>;
    async fn set_paired(&mut self, paired: Vec<PairedDevice>) -> StorageResult<()>;
    async fn add_paired(&mut self, dev: PairedDevice) -> StorageResult<()> {
        let mut data = self.get_paired().await?;
        data.push(dev);
        self.set_paired(data).await
    }
}
