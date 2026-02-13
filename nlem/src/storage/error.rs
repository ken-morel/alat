use thiserror::Error;

#[derive(Error, Debug)]
pub enum StorageError {
    #[error("Error loading storage file {0}: {1}")]
    Load(String, String),
    #[error("Could not convert object from pbuf object {0}")]
    PbufConvert(String),
    #[error("serde_yaml had an error processing yaml data: {0}")]
    SerdeYaml(serde_yaml::Error),
    #[error("input/output error saving/loading data: {0}")]
    Io(std::io::Error),
}

impl From<serde_yaml::Error> for StorageError {
    fn from(err: serde_yaml::Error) -> Self {
        Self::SerdeYaml(err)
    }
}

impl From<std::io::Error> for StorageError {
    fn from(err: std::io::Error) -> Self {
        Self::Io(err)
    }
}

impl From<StorageError> for String {
    fn from(error: StorageError) -> Self {
        format!("{error}")
    }
}
