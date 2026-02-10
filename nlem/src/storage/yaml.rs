use super::{FileStorage, StorageData, StorageResult};
use serde_yaml;

pub struct YamlFileStorage {
    path: std::path::PathBuf,
}

impl FileStorage for YamlFileStorage {
    fn new(path: &std::path::Path) -> Self {
        Self { path: path.into() }
    }
    fn write(&self, data: StorageData) -> StorageResult<()> {
        Ok(std::fs::write(
            self.path.clone(),
            serde_yaml::to_string(&data)?,
        )?)
    }
    fn load(&self) -> StorageResult<StorageData> {
        Ok(serde_yaml::from_slice(
            std::fs::read(self.path.clone())?.as_slice(),
        )?)
    }
}
