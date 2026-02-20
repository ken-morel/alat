mod config;
mod info;
mod server;

pub type TelemetryInfo = info::TelemetryInfo;

use crate::{proto, service::Service};

#[derive(Default, Clone)]
pub struct TelemetryService {
    initialized: bool,
    node: Option<crate::Node>,
    config: Option<crate::RWContainer<config::TelemetryServiceConfig>>,
    storage: Option<crate::RWContainer<super::config::ServiceConfig>>,
    info: Option<crate::RWContainer<Option<info::TelemetryInfo>>>,
}

impl TelemetryService {
    pub fn new() -> Self {
        Self::default()
    }
    pub async fn query_info(&self) -> Result<info::TelemetryInfo, crate::ErrorC> {
        Service::ensure_init(self)?;
        Ok(self
            .node
            .clone()
            .unwrap()
            .platform
            .read()
            .await
            .query_telemetry()
            .await
            .map_err(super::error::ServiceError::BackendQuery)?)
    }
}

#[tonic::async_trait]
impl super::Service for TelemetryService {
    fn is_init(&self) -> bool {
        self.initialized
    }
    fn name(&self) -> super::ServiceID {
        "telemetry"
    }
    async fn init(&mut self, node: crate::Node) -> Result<(), crate::ErrorC> {
        let mut storage = super::config::ServiceConfig::new(node.storage.clone(), self.name());
        self.config = Some(crate::contain(
            storage
                .init(config::TelemetryServiceConfig::default())
                .await?,
        ));
        self.storage = Some(crate::contain(storage));
        self.node = Some(node);
        self.info = Some(crate::contain(None));
        self.initialized = true;
        Ok(())
    }
    async fn worker(&mut self) -> Result<(), super::error::ServiceError> {
        let info = self.info.clone().unwrap();
        let config = self.config.clone().unwrap();

        loop {
            match self.query_info().await {
                Ok(data) => {
                    *info.write().await = Some(data);
                }
                Err(e) => {
                    println!("[service/telemetry] ERROR {e}");
                }
            }

            tokio::time::sleep(std::time::Duration::from_secs(
                config.read().await.poll_interval_secs.into(),
            ))
            .await;
        }
    }
    async fn grpc(
        &self,
        server: tonic::transport::server::Router,
    ) -> Result<tonic::transport::server::Router, super::error::ServiceError> {
        let server = server.add_service(
            proto::telemetry_service_server::TelemetryServiceServer::new(self.clone()),
        );
        Ok(server)
    }
}
