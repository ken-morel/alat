mod server;

use crate::proto;

#[derive(Default, Clone)]
pub struct PairService {
    node: Option<crate::Node>,
}

impl PairService {
    pub fn new() -> Self {
        Self::default()
    }
}

#[tonic::async_trait]
impl super::Service for PairService {
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
    async fn register_grpc_service_server(
        &self,
        server: tonic::transport::server::Router,
    ) -> Result<tonic::transport::server::Router, super::error::ServiceError> {
        Ok(
            server.add_service(proto::pair_service_server::PairServiceServer::new(
                self.clone(),
            )),
        )
    }
}
