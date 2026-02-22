use nlem::security::Certificate;
use nlem::storage::*;

pub struct JSONFileStorage {
    path: std::path::PathBuf,
    data: StorageData,
}

impl JSONFileStorage {
    pub fn new(path: &std::path::Path) -> Self {
        Self {
            path: path.into(),
            data: StorageData::default(),
        }
    }
    pub async fn write(&mut self) -> StorageResult<()> {
        if let Some(parent) = self.path.parent() {
            std::fs::create_dir_all(parent).map_err(|e| {
                StorageError::Other(format!("Could not create config parent directories: {e}"))
            })?;
        }
        println!("writting app data at {:?}", self.path);
        std::fs::write(
            self.path.clone(),
            serde_json::to_string_pretty(&self.data).map_err(|e| {
                StorageError::Other(format!(
                    "Could not serialize JSONFIleStoraage to string: {e}"
                ))
            })?,
        )?;
        Ok(())
    }
    pub async fn load(&mut self) -> StorageResult<()> {
        println!("loading app data at {:?}", self.path);
        self.data =
            serde_json::from_slice(std::fs::read(self.path.clone())?.as_slice()).map_err(|e| {
                StorageError::Other(format!("Error parsing config JSONFileStorage: {e}"))
            })?;
        Ok(())
    }
}

#[tonic::async_trait]
impl Storage for JSONFileStorage {
    async fn init(&mut self, data: StorageData) -> StorageResult<()> {
        if let Err(e) = self.load().await {
            println!("[storage] Error loading storage {e} creating new storage");
            self.data = data;
            self.write().await
        } else {
            Ok(())
        }
    }
    async fn get_certificate(&mut self) -> StorageResult<Certificate> {
        Ok(self.data.certificate.clone())
    }
    async fn set_certificate(&mut self, certificate: Certificate) -> StorageResult<()> {
        self.data.certificate = certificate;
        self.write().await
    }

    async fn get_info(&mut self) -> StorageResult<DeviceInfo> {
        Ok(self.data.info.clone())
    }
    async fn set_info(&mut self, info: DeviceInfo) -> StorageResult<()> {
        self.data.info = info;
        self.write().await
    }

    async fn get_paired(&mut self) -> StorageResult<Vec<PairedDevice>> {
        Ok(self.data.paired_devices.clone())
    }

    async fn set_paired(&mut self, paired: Vec<PairedDevice>) -> StorageResult<()> {
        self.data.paired_devices = paired;
        self.write().await
    }
    async fn _load_settings(&mut self, key: &str) -> StorageResult<Option<String>> {
        Ok(self.data.settings.get(key).cloned())
    }
    async fn _save_settings(&mut self, key: &str, data: String) -> StorageResult<()> {
        self.data.settings.insert(key.to_string(), data);
        self.write().await
    }
}
