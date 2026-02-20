use crate::{proto, service::Service};

#[tonic::async_trait]
impl proto::telemetry_service_server::TelemetryService for super::TelemetryService {
    async fn get_telemetry_status(
        &self,
        req: tonic::Request<proto::GetTelemetryStatusRequest>,
    ) -> Result<tonic::Response<proto::GetTelemetryStatusResponse>, tonic::Status> {
        let req = req.into_inner();
        _ = req;
        self.ensure_init()?;

        Ok(
            if let Some(info) = self.info.clone().unwrap().read().await.clone() {
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
            .into(),
        )
    }
}
