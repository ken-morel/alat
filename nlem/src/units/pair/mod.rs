mod server;

use crate::proto;
use crate::unit;
use std::sync::Arc;
use tokio::sync::RwLock;

#[derive(Debug, Clone)]
pub struct PairService {
    inner: Arc<RwLock<PairServiceInner>>,
}

#[derive(Debug)]
struct PairServiceInner {
    initialized: bool,
    node: crate::Node,
}

impl PairService {
    pub fn new(node: crate::Node) -> Self {
        Self {
            inner: Arc::new(RwLock::new(PairServiceInner {
                node,
                initialized: false,
            })),
        }
    }
}

#[tonic::async_trait]
impl unit::Unit for PairService {
    async fn is_init(&self) -> bool {
        self.inner.read().await.initialized
    }
    fn id(&self) -> unit::UnitID {
        unit::UnitID::Service("pair")
    }
    async fn init(&self) -> unit::error::UnitResult<()> {
        self.inner.write().await.initialized = true;
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
            proto::pair_service_server::PairServiceServer::new(self.clone()),
        ))
    }
}
