mod alat;
mod pair;

use super::{devicemanager, platform, proto, storage};
use std::sync::Arc;
use tokio::sync::RwLock;
use tonic::{Request, Response, Status};

pub const ALAT_PORT: u16 = 1143;

#[derive()]
pub struct Server {
    device_manager: crate::DeviceManagerC,
}
impl Server {
    pub fn new(device_manager: Arc<RwLock<devicemanager::DeviceManager>>) -> Self {
        Self {
            device_manager: device_manager.clone(),
        }
    }
    pub fn create_router(&self) -> tonic::transport::server::Router {
        tonic::transport::Server::builder()
            .add_service(proto::alat_service_server::AlatServiceServer::new(
                alat::AlatService::new(self.device_manager.clone()),
            ))
            .add_service(proto::pair_service_server::PairServiceServer::new(
                pair::PairService::new(self.device_manager.clone()),
            ))
    }
}
