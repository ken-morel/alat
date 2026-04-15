mod server;

use crate::{proto, unit};

#[derive(Clone)]
pub struct PairService {
    initialized: bool,
    node: crate::Node,
}

impl PairService {
    pub fn new(node: crate::Node) -> Self {
        Self {
            node,
            initialized: false,
        }
    }
}

#[tonic::async_trait]
impl unit::Unit for PairService {
    fn is_init(&self) -> bool {
        self.initialized
    }
    fn id(&self) -> unit::UnitID {
        unit::UnitID::Service("pair")
    }
    async fn init(&mut self) -> unit::error::UnitResult<()> {
        self.initialized = true;
        Ok(())
    }
    async fn spawn_worker(&self, _: unit::UnitSender) -> unit::SpawnWorkerResult {
        None
    }
    async fn grpc(
        &self,
        server: tonic::transport::server::Router,
    ) -> unit::error::UnitResult<tonic::transport::server::Router> {
        self.ensure_init()?;
        Ok(
            server.add_service(proto::pair_service_server::PairServiceServer::new(
                self.clone(),
            )),
        )
    }
}
