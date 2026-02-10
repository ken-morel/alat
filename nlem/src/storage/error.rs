use thiserror::Error;

#[derive(Error, Debug)]
pub enum StorageError {
    #[error("Error loading storage file {0}: {1}")]
    LoadError(String, String),
    #[error("Could not convert object from pbuf object {0}")]
    PbufConvertError(String),
    #[error("serde_yaml had an error processing yaml data: {0}")]
    SerdeYamlError(serde_yaml::Error),
    #[error("input/output error saving/loading data: {0}")]
    IoError(std::io::Error),
}

impl From<serde_yaml::Error> for StorageError {
    fn from(err: serde_yaml::Error) -> Self {
        Self::SerdeYamlError(err)
    }
}

impl From<std::io::Error> for StorageError {
    fn from(err: std::io::Error) -> Self {
        Self::IoError(err)
    }
}
