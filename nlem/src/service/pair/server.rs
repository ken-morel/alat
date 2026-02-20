use crate::proto;

#[tonic::async_trait]
impl proto::pair_service_server::PairService for super::PairService {
    async fn request_pair(
        &self,
        req: tonic::Request<proto::RequestPairRequest>,
    ) -> Result<tonic::Response<proto::RequestPairResponse>, tonic::Status> {
        let req = req.into_inner();
        let result = self
            .node
            .clone()
            .ok_or(tonic::Status::unavailable("Service uninitialized"))?
            .device_manager
            .read()
            .await
            ._handle_pair_request(
                req.info
                    .ok_or(tonic::Status::invalid_argument("Device info was blank"))?
                    .into(),
                req.certificate,
            )
            .await;
        Ok(tonic::Response::new(proto::RequestPairResponse {
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
