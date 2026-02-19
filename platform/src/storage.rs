use nlem::security::Certificate;
use nlem::storage::*;

pub struct JSONFileStorage {
    path: std::path::PathBuf,
}

impl JSONFileStorage {
    pub fn new(path: &std::path::Path) -> Self {
        Self { path: path.into() }
    }
    pub async fn write(&self, data: StorageData) -> StorageResult<()> {
        if let Some(parent) = self.path.parent() {
            std::fs::create_dir_all(parent).map_err(|e| {
                StorageError::Other(format!("Could not create config parent directories: {e}"))
            })?;
        }
        println!("writting app data at {:?}", self.path);
        Ok(std::fs::write(
            self.path.clone(),
            serde_json::to_string(&data).map_err(|e| {
                StorageError::Other(format!(
                    "Could not serialize JSONFIleStoraage to string: {e}"
                ))
            })?,
        )?)
    }
    pub async fn load(&self) -> StorageResult<StorageData> {
        println!("loading app data at {:?}", self.path);
        let ret = serde_json::from_slice(std::fs::read(self.path.clone())?.as_slice())
            .map_err(|e| StorageError::Other(format!("Error parsing config JSONFileStorage: {e}")));
        println!("  ... data loaded");
        ret
    }
}

#[tonic::async_trait]
impl Storage for JSONFileStorage {
    async fn init(&self, data: StorageData) -> StorageResult<()> {
        if let Err(e) = self.load().await {
            println!("[storage] Error loading storage {e} creating new storage");
            self.write(data).await
        } else {
            Ok(())
        }
    }
    async fn load_certificate(&self) -> StorageResult<Certificate> {
        Ok(self.load().await?.certificate)
    }
    async fn save_certificate(&self, certificate: Certificate) -> StorageResult<()> {
        let mut data = self.load().await?;
        data.certificate = certificate;
        self.write(data).await
    }

    async fn load_info(&self) -> StorageResult<DeviceInfo> {
        Ok(self.load().await?.info)
    }
    async fn save_info(&self, info: DeviceInfo) -> StorageResult<()> {
        let mut data = self.load().await?;
        data.info = info;
        self.write(data).await
    }

    async fn load_paired(&self) -> StorageResult<Vec<PairedDevice>> {
        Ok(self.load().await?.paired_devices)
    }

    async fn save_paired(&self, paired: Vec<PairedDevice>) -> StorageResult<()> {
        let mut data = self.load().await?;
        data.paired_devices = paired;
        self.write(data).await
    }
}
