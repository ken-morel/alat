mod server;

use crate::proto;

#[derive(Default, Clone)]
pub struct PairService {
    initialized: bool,
    node: Option<crate::Node>,
}

impl PairService {
    pub fn new() -> Self {
        Self::default()
    }
}

#[tonic::async_trait]
impl super::Service for PairService {
    fn is_init(&self) -> bool {
        self.initialized
    }
    fn name(&self) -> super::ServiceID {
        "pair"
    }
    async fn init(&mut self, node: crate::Node) -> Result<(), crate::ErrorC> {
        self.node = Some(node);
        Ok(())
    }
    async fn worker(&mut self) -> Result<(), super::error::ServiceError> {
        Ok(())
    }
    async fn grpc(
        &self,
        server: tonic::transport::server::Router,
    ) -> Result<tonic::transport::server::Router, super::error::ServiceError> {
        self.ensure_init()?;
        Ok(
            server.add_service(proto::pair_service_server::PairServiceServer::new(
                self.clone(),
            )),
        )
    }
}
