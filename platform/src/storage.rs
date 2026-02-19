use nlem::security::Certificate;
use nlem::storage::*;

pub struct JSONFileStorage {
    path: std::path::PathBuf,
    storage_data: Option<StorageData>,
}

impl JSONFileStorage {
    pub fn new(path: &std::path::Path) -> Self {
        Self {
            path: path.into(),
            storage_data: None,
        }
    }
    pub fn take_data(&mut self) -> StorageResult<StorageData> {
        self.storage_data.take().ok_or(StorageError::Init())
    }
    pub async fn write(&mut self) -> StorageResult<()> {
        let data = self.take_data()?;

        if let Some(parent) = self.path.parent() {
            std::fs::create_dir_all(parent).map_err(|e| {
                StorageError::Other(format!("Could not create config parent directories: {e}"))
            })?;
        }
        println!("writting app data at {:?}", self.path);
        std::fs::write(
            self.path.clone(),
            serde_json::to_string(&data).map_err(|e| {
                StorageError::Other(format!(
                    "Could not serialize JSONFIleStoraage to string: {e}"
                ))
            })?,
        )?;
        self.storage_data = Some(data);
        Ok(())
    }
    pub async fn load(&mut self) -> StorageResult<()> {
        println!("loading app data at {:?}", self.path);
        self.storage_data = Some(
            serde_json::from_slice(std::fs::read(self.path.clone())?.as_slice()).map_err(|e| {
                StorageError::Other(format!("Error parsing config JSONFileStorage: {e}"))
            })?,
        );
        Ok(())
    }
}

#[tonic::async_trait]
impl Storage for JSONFileStorage {
    async fn init(&mut self, data: StorageData) -> StorageResult<()> {
        if let Err(e) = self.load().await {
            println!("[storage] Error loading storage {e} creating new storage");
            self.storage_data = Some(data);
            self.write().await
        } else {
            Ok(())
        }
    }
    async fn get_certificate(&mut self) -> StorageResult<Certificate> {
        let data = self.take_data()?;
        let cert = data.certificate.clone();
        self.storage_data = Some(data);
        Ok(cert)
    }
    async fn set_certificate(&mut self, certificate: Certificate) -> StorageResult<()> {
        let mut data = self.take_data()?;
        data.certificate = certificate;
        self.storage_data = Some(data);
        self.write().await
    }

    async fn get_info(&mut self) -> StorageResult<DeviceInfo> {
        let data = self.take_data()?;
        let info = data.info.clone();
        self.storage_data = Some(data);
        Ok(info)
    }
    async fn set_info(&mut self, info: DeviceInfo) -> StorageResult<()> {
        let mut data = self.take_data()?;
        data.info = info;
        self.storage_data = Some(data);
        self.write().await
    }

    async fn get_paired(&mut self) -> StorageResult<Vec<PairedDevice>> {
        let data = self.take_data()?;
        let paired = data.paired_devices.clone();
        self.storage_data = Some(data);
        Ok(paired)
    }

    async fn set_paired(&mut self, paired: Vec<PairedDevice>) -> StorageResult<()> {
        let mut data = self.take_data()?;
        data.paired_devices = paired;
        self.storage_data = Some(data);
        self.write().await
    }
}
