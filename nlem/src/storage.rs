use super::proto::Certificate;

#[derive(Debug)]
pub enum StorageError {
    LoadError,
}

type StorageResult<T> = Result<T, StorageError>;

pub trait Storage {
    fn load_certificate(&self) -> StorageResult<Certificate>;
    fn save_certificate(&self, certificate: Certificate) -> StorageResult<()>;
}
