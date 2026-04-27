#[derive(Debug, Clone, thiserror::Error)]
pub enum UnitError {
    #[error("Unit {0} initialization error: {1}")]
    Init(super::UnitID, String),
    #[error("Unit {0} not initialized")]
    NotInitialized(super::UnitID),

    #[error("Unit {0} could not succesfully query it's backend: {1}")]
    BackendQuery(super::UnitID, String),

    #[error("Unit {0} error querying service storage: {1}")]
    StorageError(super::UnitID, crate::storage::StorageError),

    #[error("Unit {0} had an error: {1}")]
    Message(super::UnitID, String),
    #[error("Request not authenticated")]
    Unauthenticated(),

    #[error("Unit manager not initialized")]
    ManagerNotInitialized(),
    #[error("Unit {0} already registered")]
    UnitAlreadyRegistered(super::UnitID),
}

pub type UnitResult<T> = Result<T, UnitError>;

impl From<UnitError> for tonic::Status {
    fn from(val: UnitError) -> Self {
        match val {
            err => Self::internal(err.to_string()),
        }
    }
}
