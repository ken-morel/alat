use crate::{proto, unit::Unit};

#[tonic::async_trait]
impl proto::telemetry_service_server::TelemetryService for super::TelemetryService {
    async fn get_telemetry_status(
        &self,
        req: tonic::Request<proto::GetTelemetryStatusRequest>,
    ) -> Result<tonic::Response<proto::GetTelemetryStatusResponse>, tonic::Status> {
        self.authenticate(
            &*self.node.device_manager.read().await,
            &req.into_inner().call.unwrap(),
        )
        .await?;

        self.ensure_init()?;

        Ok(if let Some(info) = self.info.read().await.clone() {
            proto::GetTelemetryStatusResponse {
                reply: Some(proto::ServiceReply {
                    status: proto::ServiceReplyStatus::Retry.into(),
                    message: "Telemetry info not available at the moment".to_string(),
                }),
                telemetry_status: Some(info.into()),
            }
        } else {
            proto::GetTelemetryStatusResponse {
                reply: Some(proto::ServiceReply {
                    status: proto::ServiceReplyStatus::Retry.into(),
                    message: "Telemetry info not available at the moment".to_string(),
                }),
                telemetry_status: None,
            }
        }
        .into())
    }
}
