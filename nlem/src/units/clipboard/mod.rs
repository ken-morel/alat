mod clipboardcontent;
mod server;
pub use clipboardcontent::ClipboardContent;

use crate::proto;
use crate::unit;
use std::sync::Arc;
use tokio::sync::RwLock;

#[tonic::async_trait]
pub trait Clipboard: std::fmt::Debug {
    async fn get_content(&self) -> Result<ClipboardContent, String>;
    async fn set_content(&self, content: ClipboardContent) -> Result<(), String>;
}

pub type ClipboardC = crate::MContainer<dyn Clipboard + Send + Sync + 'static>;

#[derive(Debug, Clone)]
pub struct ClipboardService {
    inner: Arc<RwLock<ClipboardServiceInner>>,
}

#[derive(Debug)]
struct ClipboardServiceInner {
    initialized: bool,
    node: crate::Node,
    clipboard: Option<ClipboardC>,
}

impl ClipboardService {
    pub fn new(node: crate::Node) -> Self {
        Self {
            inner: Arc::new(RwLock::new(ClipboardServiceInner {
                node,
                clipboard: None,
                initialized: false,
            })),
        }
    }
}

#[tonic::async_trait]
impl unit::Unit for ClipboardService {
    async fn is_init(&self) -> bool {
        self.inner.read().await.initialized
    }
    fn id(&self) -> unit::UnitID {
        unit::UnitID::Service("clipboard")
    }
    async fn init(&self) -> unit::error::UnitResult<()> {
        let node = self.inner.read().await.node.clone();
        let clipboard = node
            .platform
            .read()
            .await
            .clipboard()
            .await
            .map_err(|e| unit::error::UnitError::BackendQuery(self.id(), e))?;

        let mut inner = self.inner.write().await;
        inner.clipboard = Some(clipboard);
        inner.initialized = true;
        Ok(())
    }
    async fn spawn_worker(&self, _: unit::UnitSender) -> unit::SpawnWorkerResult {
        None
    }
    async fn grpc(
        &self,
        server: tonic::transport::server::Router,
    ) -> unit::error::UnitResult<tonic::transport::server::Router> {
        self.ensure_init().await?;
        Ok(server.add_service(
            proto::clipboard_service_server::ClipboardServiceServer::new(self.clone()),
        ))
    }
}
