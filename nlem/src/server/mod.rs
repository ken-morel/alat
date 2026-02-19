mod alat;

use crate::service;

use super::{devicemanager, proto};
use std::sync::Arc;
use tokio::sync::RwLock;
use tonic::{Request, Response, Status};

pub const ALAT_PORT: u16 = 1143;

#[derive()]
pub struct Server {
    device_manager: crate::DeviceManagerC,
    service_manager: crate::ServiceManagerC,
}
impl Server {
    pub fn new(
        device_manager: crate::DeviceManagerC,
        service_manager: crate::ServiceManagerC,
    ) -> Self {
        Self {
            device_manager,
            service_manager,
        }
    }
    pub async fn create_router(
        &self,
    ) -> Result<tonic::transport::server::Router, service::error::ServiceError> {
        let mut router = tonic::transport::Server::builder();
        let router = router.add_service(proto::alat_service_server::AlatServiceServer::new(
            alat::AlatService::new(self.device_manager.clone()),
        ));
        self.service_manager
            .read()
            .await
            .register_grpc_service_servers(router)
            .await
    }
}
