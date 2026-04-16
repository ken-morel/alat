mod clipboardcontent;
mod server;
pub use clipboardcontent::ClipboardContent;

use crate::proto;
use crate::unit;

#[tonic::async_trait]
pub trait Clipboard: std::fmt::Debug {
    async fn get_content(&self) -> Result<ClipboardContent, String>;
    async fn set_content(&self, content: ClipboardContent) -> Result<(), String>;
}

pub type ClipboardC = crate::MContainer<dyn Clipboard + Send + Sync + 'static>;

#[derive(Debug, Clone)]
pub struct ClipboardService {
    initialized: bool,
    node: crate::Node,
    clipboard: Option<ClipboardC>,
}

impl ClipboardService {
    pub fn new(node: crate::Node) -> Self {
        Self {
            node,
            clipboard: None,
            initialized: false,
        }
    }
}

#[tonic::async_trait]
impl unit::Unit for ClipboardService {
    fn is_init(&self) -> bool {
        self.initialized
    }
    fn id(&self) -> unit::UnitID {
        unit::UnitID::Service("clipboard")
    }
    async fn init(&mut self) -> unit::error::UnitResult<()> {
        self.clipboard = Some(
            self.node
                .platform
                .read()
                .await
                .clipboard()
                .await
                .map_err(|e| unit::error::UnitError::BackendQuery(self.id(), e))?,
        );
        self.initialized = true;
        Ok(())
    }
    async fn spawn_worker(&mut self, _: unit::UnitSender) -> unit::SpawnWorkerResult {
        None
    }
    async fn grpc(
        &mut self,
        server: tonic::transport::server::Router,
    ) -> unit::error::UnitResult<tonic::transport::server::Router> {
        self.ensure_init()?;
        Ok(server.add_service(
            proto::clipboard_service_server::ClipboardServiceServer::new(self.clone()),
        ))
    }
}
