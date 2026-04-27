use crate::proto;
use crate::unit::Unit;

#[tonic::async_trait]
impl proto::pair_service_server::PairService for super::PairService {
    async fn request_pair(
        &self,
        req: tonic::Request<proto::RequestPairRequest>,
    ) -> Result<tonic::Response<proto::RequestPairResponse>, tonic::Status> {
        self.ensure_init().await?;
        let req = req.into_inner();
        let self = self.inner.read().await;

        match self
            .node
            .device_manager
            ._handle_pair_request(
                req.info
                    .ok_or(tonic::Status::invalid_argument("Missing argument"))?
                    .into(),
                req.certificate.into(),
            )
            .await
        {
            Ok(dev) => {
                return Ok(tonic::Response::new(proto::RequestPairResponse {
                    result: Some(proto::request_pair_response::Result::Success(
                        proto::RequestPairResponseSuccess {
                            token: dev.token.into(),
                            certificate: self
                                .node
                                .device_manager
                                .device_certificate
                                .read()
                                .await
                                .clone(),
                            info: Some(
                                self.node
                                    .device_manager
                                    .this_device
                                    .read()
                                    .await
                                    .info
                                    .clone()
                                    .into(),
                            ),
                        },
                    )),
                }));
            }
            Err(val) => Ok(tonic::Response::new(proto::RequestPairResponse {
                result: Some(proto::request_pair_response::Result::Failure(
                    proto::RequestPairResponseFailure { reason: val },
                )),
            })),
        }
    }
}
