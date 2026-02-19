#[derive(Debug, Clone, thiserror::Error)]
pub enum ServiceError {
    #[error("Service initialization error: {0}")]
    Init(String),
    #[error("Service {0} not initialized")]
    NotInitialized(String),
}
