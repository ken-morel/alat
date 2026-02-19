use crate::proto::RequestPairResponse;

use super::*;

#[derive()]
pub struct PairService {
    device_manager: crate::DeviceManagerC,
}

impl PairService {
    pub fn new(device_manager: crate::DeviceManagerC) -> Self {
        Self { device_manager }
    }
}

#[tonic::async_trait]
impl proto::pair_service_server::PairService for PairService {
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
                req.certificate,
            )
            .await;
        Ok(Response::new(RequestPairResponse {
            result: Some(match result {
                Ok(paired) => proto::request_pair_response::Result::Success(
                    proto::RequestPairResponseSuccess {
                        token: paired.token.into(),
                        certificate: paired.certificate,
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
