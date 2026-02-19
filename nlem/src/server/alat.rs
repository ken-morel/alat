use super::*;

#[derive()]
pub struct AlatService {
    device_manager: Arc<RwLock<devicemanager::DeviceManager>>,
}
impl AlatService {
    pub fn new(device_manager: Arc<RwLock<devicemanager::DeviceManager>>) -> Self {
        Self { device_manager }
    }
}

#[tonic::async_trait]
impl proto::alat_service_server::AlatService for AlatService {
    async fn get_device_info(
        &self,
        _: Request<proto::GetDeviceInfoRequest>,
    ) -> Result<Response<proto::GetDeviceInfoResponse>, Status> {
        Ok(Response::new(proto::GetDeviceInfoResponse {
            info: Some(
                self.device_manager
                    .read()
                    .await
                    .this_device
                    .read()
                    .await
                    .info
                    .clone()
                    .into(),
            ),
        }))
    }
}
