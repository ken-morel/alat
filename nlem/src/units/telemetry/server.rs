use crate::proto;
use crate::unit::Unit;

#[tonic::async_trait]
impl proto::telemetry_service_server::TelemetryService for super::TelemetryService {
    async fn get_telemetry_status(
        &self,
        req: tonic::Request<proto::GetTelemetryStatusRequest>,
    ) -> Result<tonic::Response<proto::GetTelemetryStatusResponse>, tonic::Status> {
        let inner = self.inner.read().await;
        self.authenticate(&inner.node.device_manager, &req.into_inner().call.unwrap())
            .await?;
        self.ensure_init().await?;

        Ok(tonic::Response::new(match self.query_telemetry().await {
            Ok(tel) => proto::GetTelemetryStatusResponse {
                reply: Some(proto::ServiceReply {
                    status: proto::ServiceReplyStatus::Ok.into(),
                    message: String::default(),
                }),
                telemetry_status: Some(tel.into()),
            },
            Err(e) => proto::GetTelemetryStatusResponse {
                reply: Some(proto::ServiceReply {
                    status: proto::ServiceReplyStatus::InternalError.into(),
                    message: e.to_string(),
                }),
                telemetry_status: None,
            },
        }))
    }
}
