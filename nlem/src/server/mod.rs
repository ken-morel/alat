mod alat;
mod pair;

use super::{devicemanager, devicemanager::discovered, platform, proto, storage};
use std::sync::Arc;
use tokio::sync::Mutex;
use tonic::{Request, Response, Status};

pub const ALAT_PORT: u16 = 1143;

#[derive(Debug)]
pub struct Server {
    server: tonic::transport::server::Router,
}
impl Server {
    pub fn new<
        S: storage::Storage + 'static,
        P: platform::Platform + 'static,
        D: discovered::DiscoveryManager + 'static,
    >(
        device_manager: Arc<Mutex<devicemanager::DeviceManager<S, P, D>>>,
    ) -> Self {
        let server = tonic::transport::Server::builder()
            .add_service(proto::alat_service_server::AlatServiceServer::new(
                alat::AlatService::new(device_manager.clone()),
            ))
            .add_service(proto::pair_service_server::PairServiceServer::new(
                pair::PairService::new(device_manager.clone()),
            ));
        Self { server }
    }
    pub async fn serve(self, addr: std::net::SocketAddr) {
        self.server.serve(addr).await.expect("Server failed");
    }
}
