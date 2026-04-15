mod server;

use crate::proto;
use std::path::PathBuf;

#[derive(Debug, Clone, Default, PartialEq)]
pub enum ClipboardContent {
    Text(String),
    Image((u32, u32), Vec<u8>),
    FileList(Vec<PathBuf>),
    #[default]
    Empty,
}
impl From<ClipboardContent> for proto::ClipboardContent {
    fn from(c: ClipboardContent) -> Self {
        Self {
            data: Some(match c {
                ClipboardContent::Text(txt) => {
                    proto::clipboard_content::Data::Text(proto::ClipboardContentText { data: txt })
                }
                ClipboardContent::Image((width, height), data) => {
                    proto::clipboard_content::Data::Image(proto::ClipboardContentImage {
                        width,
                        height,
                        data,
                    })
                }
                ClipboardContent::FileList(files) => {
                    proto::clipboard_content::Data::Files(proto::ClipboardContentFiles {
                        files: files
                            .into_iter()
                            .map(|b| b.to_str().unwrap_or_default().to_string())
                            .collect(),
                    })
                }
                ClipboardContent::Empty => {
                    proto::clipboard_content::Data::Empty(proto::ClipboardContentEmpty {})
                }
            }),
        }
    }
}

#[tonic::async_trait]
pub trait Clipboard {
    async fn get_content(&self) -> Result<ClipboardContent, String>;
    async fn set_content(&self, content: ClipboardContent) -> Result<(), String>;
}

pub type ClipboardC = crate::MContainer<dyn Clipboard + Send + Sync + 'static>;

#[derive(Default, Clone)]
pub struct ClipboardService {
    initialized: bool,
    node: Option<crate::Node>,
    clipboard: Option<ClipboardC>,
}

impl ClipboardService {
    pub fn new() -> Self {
        Self::default()
    }
}

#[tonic::async_trait]
impl super::Service for ClipboardService {
    fn is_init(&self) -> bool {
        self.initialized
    }
    fn name(&self) -> super::ServiceID {
        "clipboard"
    }
    async fn init(&mut self, node: crate::Node) -> super::error::ServiceResult<()> {
        self.clipboard = Some(
            node.platform
                .read()
                .await
                .clipboard()
                .await
                .map_err(|e| super::error::ServiceError::BackendQuery(self.name(), e))?,
        );
        self.node = Some(node);
        self.initialized = true;
        Ok(())
    }
    async fn spawn_worker(&self, _: super::ServiceChannel) -> super::SpawnWorkerResult {
        None
    }
    async fn grpc(
        &self,
        server: tonic::transport::server::Router,
    ) -> super::error::ServiceResult<tonic::transport::server::Router> {
        self.ensure_init()?;
        Ok(server.add_service(
            proto::clipboard_service_server::ClipboardServiceServer::new(self.clone()),
        ))
    }
}
