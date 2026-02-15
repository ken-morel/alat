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

pub trait Storage: Sized + Send + Sync {
    fn init(
        &self,
        data: StorageData,
    ) -> impl std::future::Future<Output = StorageResult<StorageData>> + Send;

    fn load_certificate(
        &self,
    ) -> impl std::future::Future<Output = StorageResult<Certificate>> + Send;
    fn save_certificate(
        &self,
        certificate: Certificate,
    ) -> impl std::future::Future<Output = StorageResult<()>> + Send;

    fn load_info(&self) -> impl std::future::Future<Output = StorageResult<DeviceInfo>> + Send;
    fn save_info(
        &self,
        info: DeviceInfo,
    ) -> impl std::future::Future<Output = StorageResult<()>> + Send;

    fn load_paired(
        &self,
    ) -> impl std::future::Future<Output = StorageResult<Vec<PairedDevice>>> + Send;
    fn save_paired(
        &self,
        paired: Vec<PairedDevice>,
    ) -> impl std::future::Future<Output = StorageResult<()>> + Send;
    fn add_paired(
        &self,
        dev: PairedDevice,
    ) -> impl std::future::Future<Output = StorageResult<()>> + Send {
        async {
            let mut data = self.load_paired().await?;
            data.push(dev);
            self.save_paired(data).await
        }
    }
}
