mod info;
mod server;
mod config;

pub use info::TelemetryInfo;

use crate::proto;
use crate::unit;
use std::sync::Arc;
use tokio::sync::RwLock;

#[derive(Debug, Clone)]
pub struct TelemetryService {
    inner: Arc<RwLock<TelemetryServiceInner>>,
}

#[derive(Debug)]
struct TelemetryServiceInner {
    initialized: bool,
    node: crate::Node,
}

impl TelemetryService {
    pub fn new(node: crate::Node) -> Self {
        Self {
            inner: Arc::new(RwLock::new(TelemetryServiceInner {
                node,
                initialized: false,
            })),
        }
    }
}

#[tonic::async_trait]
impl unit::Unit for TelemetryService {
    async fn is_init(&self) -> bool {
        self.inner.read().await.initialized
    }
    fn id(&self) -> unit::UnitID {
        unit::UnitID::Service("telemetry")
    }
    async fn init(&self) -> unit::error::UnitResult<()> {
        self.inner.write().await.initialized = true;
        Ok(())
    }
    async fn spawn_worker(&self, _channel: unit::UnitSender) -> unit::SpawnWorkerResult {
        None
    }
    async fn grpc(
        &self,
        server: tonic::transport::server::Router,
    ) -> unit::error::UnitResult<tonic::transport::server::Router> {
        self.ensure_init().await?;
        Ok(server.add_service(
            proto::telemetry_service_server::TelemetryServiceServer::new(self.clone()),
        ))
    }
}
