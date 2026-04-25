use super::*;
use std::sync::Arc;

#[derive()]
pub struct AlatService {
    device_manager: Arc<devicemanager::DeviceManager>,
}
impl AlatService {
    pub fn new(device_manager: Arc<devicemanager::DeviceManager>) -> Self {
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
