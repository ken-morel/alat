use crate::proto;
use crate::unit::Unit;

#[tonic::async_trait]
impl proto::clipboard_service_server::ClipboardService for super::ClipboardService {
    async fn get_clipboard(
        &self,
        req: tonic::Request<proto::GetClipboardRequest>,
    ) -> Result<tonic::Response<proto::GetClipboardResponse>, tonic::Status> {
        self.ensure_init().await?;
        let inner = self.inner.read().await;
        self.authenticate(
            &inner.node.device_manager,
            &req.into_inner().call.unwrap(),
        )
        .await?;

        match inner
            .clipboard
            .clone()
            .unwrap()
            .lock()
            .await
            .get_content()
            .await
        {
            Ok(content) => Ok(tonic::Response::new(proto::GetClipboardResponse {
                reply: Some(proto::ServiceReply {
                    status: proto::ServiceReplyStatus::Ok.into(),
                    message: String::default(),
                }),
                content: Some(content.into()),
            })),
            Err(e) => Ok(tonic::Response::new(proto::GetClipboardResponse {
                reply: Some(proto::ServiceReply {
                    status: proto::ServiceReplyStatus::InternalError.into(),
                    message: format!("Error reading clipboard from platform backend: {e}"),
                }),
                content: None,
            })),
        }
    }
    async fn send_clipboard(
        &self,
        req: tonic::Request<proto::SendClipboardRequest>,
    ) -> Result<tonic::Response<proto::SendClipboardResponse>, tonic::Status> {
        self.ensure_init().await?;
        let req = req.into_inner();
        let inner = self.inner.read().await;
        self.authenticate(
            &inner.node.device_manager,
            &req.call.unwrap(),
        )
        .await?;

        match req.content {
            Some(data) => match data.data {
                Some(data) => {
                    //TODO: Do something with new clipboard data
                    _ = data;
                    Ok(tonic::Response::new(proto::SendClipboardResponse {
                        reply: Some(proto::ServiceReply {
                            status: proto::ServiceReplyStatus::Ok.into(),
                            message: format!("Ok"),
                        }),
                    }))
                }
                None => Ok(tonic::Response::new(proto::SendClipboardResponse {
                    reply: Some(proto::ServiceReply {
                        status: proto::ServiceReplyStatus::Ok.into(),
                        message: format!("Ok"),
                    }),
                })), //Clipboard was cleared
            },
            None => {
                Ok(tonic::Response::new(proto::SendClipboardResponse {
                    reply: Some(proto::ServiceReply {
                        status: proto::ServiceReplyStatus::InternalError.into(), //TODO: Use InvalidArguments
                        message: format!("No clipboard data was sent"),
                    }),
                }))
            }
        }
    }
}
