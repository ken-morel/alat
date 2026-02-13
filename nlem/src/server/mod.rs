use super::{devicemanager, platform, proto, storage};
use std::sync::Arc;
use tokio::sync::Mutex;
use tonic::{Request, Response, Status};

#[derive(Debug)]
pub struct Server {
    server: tonic::transport::server::Router,
}
impl Server {
    pub fn new<S, P>(device_manager: Arc<Mutex<devicemanager::DeviceManager<S, P>>>) -> Self
    where
        S: storage::Storage + 'static,
        P: platform::Platform + 'static,
    {
        let server = tonic::transport::Server::builder().add_service(
            proto::alat_service_server::AlatServiceServer::new(AlatService::new(device_manager)),
        );
        Self { server }
    }
    pub async fn serve(self, addr: std::net::SocketAddr) {
        self.server.serve(addr).await.expect("Server failed");
    }
}

#[derive(Debug)]
pub struct AlatService<S: storage::Storage, P: platform::Platform> {
    device_manager: Arc<Mutex<devicemanager::DeviceManager<S, P>>>,
}
impl<S: storage::Storage, P: platform::Platform> AlatService<S, P> {
    pub fn new(device_manager: Arc<Mutex<devicemanager::DeviceManager<S, P>>>) -> Self {
        Self { device_manager }
    }
}

#[tonic::async_trait]
impl<S: storage::Storage + 'static, P: platform::Platform + 'static>
    proto::alat_service_server::AlatService for AlatService<S, P>
{
    async fn get_device_info(
        &self,
        _: Request<proto::GetDeviceInfoRequest>,
    ) -> Result<Response<proto::GetDeviceInfoResponse>, Status> {
        Ok(Response::new(proto::GetDeviceInfoResponse {
            info: Some(self.device_manager.lock().await.device_info.clone().into()),
        }))
    }
}
