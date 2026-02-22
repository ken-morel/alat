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

#[derive(Debug, Clone, serde::Deserialize, serde::Serialize, Default)]
pub struct StorageData {
    pub certificate: Certificate,
    pub paired_devices: Vec<PairedDevice>,
    pub info: DeviceInfo,
    pub settings: std::collections::BTreeMap<String, String>,
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

    async fn _load_settings(&mut self, key: &str) -> StorageResult<Option<String>>;
    async fn _save_settings(&mut self, key: &str, value: String) -> StorageResult<()>;

    async fn load_settings(&mut self, key: &str) -> StorageResult<Option<serde_json::Value>> {
        if let Some(raw_data) = self._load_settings(key).await? {
            serde_json::from_str(&raw_data).map_err(|e| StorageError::Deserialize(e.to_string()))
        } else {
            Ok(None)
        }
    }
    async fn save_settings(&mut self, key: &str, data: &serde_json::Value) -> StorageResult<()> {
        let raw_data =
            serde_json::to_string(data).map_err(|e| StorageError::Deserialize(e.to_string()))?;
        self._save_settings(key, raw_data).await
    }
}
