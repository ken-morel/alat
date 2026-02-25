use std::path::PathBuf;

#[derive(Debug, Clone)]
pub enum ClipboardContent {
    Text(String),
    Image((usize, usize), Vec<u8>),
    FileList(Vec<PathBuf>),
}

#[tonic::async_trait]
pub trait Clipboard {
    async fn watch(&self) -> Result<tokio::sync::mpsc::Receiver<ClipboardContent>, String>;
    async fn get_content(&self) -> Result<ClipboardContent, String>;
    async fn set_content(&self, content: ClipboardContent) -> Result<(), String>;
}
