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

mod yaml;

#[derive(Debug, serde::Deserialize, serde::Serialize)]
pub struct StorageData {
    certificate: Certificate,
    paired_devices: Vec<PairedDevice>,
    info: DeviceInfo,
}

use super::proto;

type StorageResult<T> = Result<T, StorageError>;

pub trait Storage: Send + Sync {
    fn load_certificate(&self) -> StorageResult<Certificate>;
    fn save_certificate(&self, certificate: Certificate) -> StorageResult<()>;

    fn load_info(&self) -> StorageResult<DeviceInfo>;
    fn save_info(&self, info: DeviceInfo) -> StorageResult<()>;

    fn load_paired(&self) -> StorageResult<Vec<PairedDevice>>;
    fn save_paired(&self, paired: Vec<PairedDevice>) -> StorageResult<()>;
    fn add_paired(&self, dev: PairedDevice) -> StorageResult<()> {
        let mut data = self.load_paired()?;
        data.push(dev);
        self.save_paired(data)
    }
}

pub trait FileStorage {
    fn write(&self, data: StorageData) -> StorageResult<()>;
    fn load(&self) -> StorageResult<StorageData>;
    //TODO: Remove if unnneded or disturbs
    fn new(path: &std::path::Path) -> Self;
    fn load_certificate(&self) -> StorageResult<Certificate> {
        Ok(self.load()?.certificate)
    }
    fn save_certificate(&self, certificate: Certificate) -> StorageResult<()> {
        let mut data = self.load()?;
        data.certificate = certificate;
        self.write(data)
    }

    fn load_info(&self) -> StorageResult<DeviceInfo> {
        Ok(self.load()?.info)
    }
    fn save_info(&self, info: DeviceInfo) -> StorageResult<()> {
        let mut data = self.load()?;
        data.info = info;
        self.write(data)
    }

    fn load_paired(&self) -> StorageResult<Vec<PairedDevice>> {
        Ok(self.load()?.paired_devices)
    }

    fn save_paired(&self, paired: Vec<PairedDevice>) -> StorageResult<()> {
        let mut data = self.load()?;
        data.paired_devices = paired;
        self.write(data)
    }
}
