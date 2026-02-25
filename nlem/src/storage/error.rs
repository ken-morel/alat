use thiserror::Error;

#[derive(Error, Debug, Clone)]
pub enum StorageError {
    #[error("Error loading storage {0}: {1}")]
    Load(String, String),

    #[error("Could not serialize data: {0}")]
    Serialize(String),

    #[error("Could not deserialize data: {0}")]
    Deserialize(String),

    #[error("Could not convert object from pbuf object {0}")]
    PbufConvert(String),

    #[error("{0}")]
    Other(String),
    #[error("input/output error saving/loading node storage data: {0}")]
    Io(String),
    #[error("Storage was not initialized")]
    Init(),
}

impl From<std::io::Error> for StorageError {
    fn from(err: std::io::Error) -> Self {
        Self::Io(err.to_string())
    }
}

impl From<StorageError> for String {
    fn from(error: StorageError) -> Self {
        format!("{error}")
    }
}
