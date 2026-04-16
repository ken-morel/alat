mod alat;

use crate::unit;

use super::{devicemanager, proto};
use std::sync::Arc;
use tokio::sync::RwLock;
use tonic::{Request, Response, Status};

pub const ALAT_PORT: u16 = 1143;

#[derive(Debug)]
pub struct Server {
    device_manager: crate::DeviceManagerC,
    unit_manager: crate::UnitManagerC,
}
impl Server {
    pub fn new(device_manager: crate::DeviceManagerC, unit_manager: crate::UnitManagerC) -> Self {
        Self {
            device_manager,
            unit_manager,
        }
    }
    pub async fn create_router(
        &self,
    ) -> Result<tonic::transport::server::Router, unit::error::UnitError> {
        let mut router = tonic::transport::Server::builder();
        let router = router.add_service(proto::alat_service_server::AlatServiceServer::new(
            alat::AlatService::new(self.device_manager.clone()),
        ));
        self.unit_manager
            .read()
            .await
            .register_grpc_unit_servers(router)
            .await
    }
}
