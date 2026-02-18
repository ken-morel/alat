use crate::proto::RequestPairResponse;

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
        req: Request<proto::RequestPairRequest>,
    ) -> Result<Response<proto::RequestPairResponse>, Status> {
        let req = req.into_inner();
        let result = self
            .device_manager
            .read()
            .await
            ._handle_pair_request(
                req.info
                    .ok_or(Status::invalid_argument("Device info was blank"))?
                    .into(),
                req.certificate.into(),
            )
            .await;
        Ok(Response::new(RequestPairResponse {
            result: Some(match result {
                Ok(paired) => proto::request_pair_response::Result::Success(
                    proto::RequestPairResponseSuccess {
                        token: paired.token.into(),
                        certificate: paired.certificate.into(),
                        info: Some(paired.info.into()),
                    },
                ),
                Err(reason) => proto::request_pair_response::Result::Failure(
                    proto::RequestPairResponseFailure { reason },
                ),
            }),
        }))
    }
}
