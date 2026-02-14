use super::*;

#[derive(Debug)]
pub struct AlatService<
    S: storage::Storage,
    P: platform::Platform<S, D>,
    D: discovered::DiscoveryManager,
> {
    device_manager: Arc<RwLock<devicemanager::DeviceManager<S, P, D>>>,
}
impl<S: storage::Storage, P: platform::Platform<S, D>, D: discovered::DiscoveryManager>
    AlatService<S, P, D>
{
    pub fn new(device_manager: Arc<RwLock<devicemanager::DeviceManager<S, P, D>>>) -> Self {
        Self { device_manager }
    }
}

#[tonic::async_trait]
impl<
    S: storage::Storage + 'static,
    P: platform::Platform<S, D> + 'static,
    D: discovered::DiscoveryManager + 'static,
> proto::alat_service_server::AlatService for AlatService<S, P, D>
{
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
