use super::proto::Certificate;
use thiserror::Error;

#[derive(Error, Debug)]
pub enum StorageError {
    #[error("Error loading storage file {0}")]
    LoadError(String),
}

type StorageResult<T> = Result<T, StorageError>;

pub trait Storage {
    fn load_certificate(&self) -> StorageResult<Certificate>;
    fn save_certificate(&self, certificate: Certificate) -> StorageResult<()>;
}
