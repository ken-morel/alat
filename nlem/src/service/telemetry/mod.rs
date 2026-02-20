mod config;

use crate::proto;

#[derive(Default, Clone)]
pub struct TelemetryService {
    node: Option<crate::Node>,
    config: Option<super::config::ServiceConfig>,
}

impl TelemetryService {
    pub fn new() -> Self {
        Self::default()
    }
}

#[tonic::async_trait]
impl super::Service for TelemetryService {
    fn name(&self) -> super::ServiceID {
        "telemetry"
    }
    async fn init(&mut self, node: crate::Node) -> Result<(), crate::ErrorC> {
        self.config = Some(super::config::ServiceConfig::new(
            node.storage.clone(),
            self.name(),
        ));
        self.node = Some(node);
        Ok(())
    }
    async fn worker(&mut self) -> Result<(), super::error::ServiceError> {
        Ok(())
    }
    async fn register_grpc_service_server(
        &self,
        server: tonic::transport::server::Router,
    ) -> Result<tonic::transport::server::Router, super::error::ServiceError> {
        Ok(server)
    }
}
