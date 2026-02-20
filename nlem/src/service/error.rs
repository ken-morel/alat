#[derive(Debug, Clone, thiserror::Error)]
pub enum ServiceError {
    #[error("Service initialization error: {0}")]
    Init(String),
    #[error("Service {0} not initialized")]
    NotInitialized(String),

    #[error("Service could not succesfully query it's backend: {0}")]
    BackendQuery(String),

    #[error("{0}")]
    Message(String),
}

impl From<ServiceError> for tonic::Status {
    fn from(val: ServiceError) -> Self {
        match val {
            err => Self::internal(err.to_string()),
        }
    }
}
