#[derive(Debug, Clone, thiserror::Error)]
pub enum ServiceError {
    #[error("Service {0} initialization error: {1}")]
    Init(super::ServiceID, String),
    #[error("Service {0} not initialized")]
    NotInitialized(super::ServiceID),

    #[error("Service {0} could not succesfully query it's backend: {1}")]
    BackendQuery(super::ServiceID, String),

    #[error("Service {0} error querying service storage: {1}")]
    StorageError(super::ServiceID, crate::storage::StorageError),

    #[error("Service {0} had an error: {1}")]
    Message(super::ServiceID, String),
}

pub type ServiceResult<T> = Result<T, ServiceError>;

impl From<ServiceError> for tonic::Status {
    fn from(val: ServiceError) -> Self {
        match val {
            err => Self::internal(err.to_string()),
        }
    }
}
