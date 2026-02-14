#[derive(Debug, Clone, thiserror::Error)]
pub enum ClientError {
    #[error("Missing item in response: {0}")]
    MissingItem(String),
    #[error("Server replied with an empty response")]
    EmptyResponse,
}
