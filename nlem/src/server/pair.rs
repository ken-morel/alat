use super::*;

#[derive(Debug)]
pub struct PairService<
    S: storage::Storage,
    P: platform::Platform<S, D>,
    D: discovered::DiscoveryManager,
> {
    device_manager: Arc<RwLock<devicemanager::DeviceManager<S, P, D>>>,
}

impl<S: storage::Storage, P: platform::Platform<S, D>, D: discovered::DiscoveryManager>
    PairService<S, P, D>
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
> proto::pair_service_server::PairService for PairService<S, P, D>
{
    async fn request_pair(
        &self,
        _: Request<proto::RequestPairRequest>,
    ) -> Result<Response<proto::RequestPairResponse>, Status> {
        Ok(Response::new(proto::RequestPairResponse {
            result: Some(proto::request_pair_response::Result::Failure(
                proto::RequestPairResponseFailure {
                    reason: String::from("Pairing not setup"),
                },
            )),
        }))
    }
}
