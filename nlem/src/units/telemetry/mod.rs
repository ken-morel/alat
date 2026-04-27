mod config;
mod info;
mod server;

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
    config: config::TelemetryServiceConfig,
    info: Option<(std::time::Instant, TelemetryInfo)>,
}

impl TelemetryService {
    pub fn new(node: crate::Node) -> Self {
        Self {
            inner: Arc::new(RwLock::new(TelemetryServiceInner {
                node,
                initialized: false,
                config: config::TelemetryServiceConfig::default(),
                info: None,
            })),
        }
    }
    pub async fn _query_telemetry(&self) -> unit::error::UnitResult<TelemetryInfo> {
        self.inner
            .read()
            .await
            .node
            .platform
            .read()
            .await
            .query_telemetry()
            .await
            .map_err(|e| unit::error::UnitError::Message(unit::Unit::id(self), e.to_string()))
    }
    pub async fn query_telemetry(&self) -> unit::error::UnitResult<TelemetryInfo> {
        let inner = self.inner.read().await;
        if let Some((timestamp, info)) = &self.inner.read().await.info
            && std::time::Instant::now().duration_since(*timestamp)
                < std::time::Duration::from_secs(inner.config.poll_interval_secs.into())
        {
            return Ok(info.clone());
        }
        drop(inner);

        let info = self._query_telemetry().await?;
        self.inner.write().await.info = Some((std::time::Instant::now(), info.clone()));
        Ok(info)
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
        let mut inner = self.inner.write().await;
        inner.initialized = true;
        if let Some(conf) = inner
            .node
            .clone()
            .storage
            .lock()
            .await
            .load_settings(&self.id().to_string())
            .await
            .map_err(|e| unit::error::UnitError::Init(self.id(), e.to_string()))?
        {
            inner.config = serde_json::from_value(conf)
                .map_err(|e| unit::error::UnitError::Message(self.id(), e.to_string()))?;
        }
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
