mod error;

pub type ClientError = error::ClientError;

use super::{proto, storage};
use crate::security;
use std::net::SocketAddr;

#[derive(Debug, Clone)]
pub struct Client {
    alat_client: proto::alat_service_client::AlatServiceClient<tonic::transport::Channel>,
    pair_client: proto::pair_service_client::PairServiceClient<tonic::transport::Channel>,
    server_addr: SocketAddr,
}

impl Client {
    pub async fn connect(
        addr: SocketAddr,
    ) -> Result<Self, Box<dyn std::error::Error + Send + Sync>> {
        let alat_client =
            proto::alat_service_client::AlatServiceClient::connect(addr.to_string()).await?;
        let pair_client =
            proto::pair_service_client::PairServiceClient::connect(addr.to_string()).await?;

        Ok(Self {
            server_addr: addr,
            alat_client,
            pair_client,
        })
    }
    pub async fn reconnect(self) -> Result<Self, Box<dyn std::error::Error + Send + Sync>> {
        Client::connect(self.server_addr).await
    }

    pub async fn get_device_info(&mut self) -> Result<storage::DeviceInfo, String> {
        Ok(self
            .alat_client
            .get_device_info(proto::GetDeviceInfoRequest {})
            .await
            .map_err(|e| e.to_string())?
            .into_inner()
            .info
            .ok_or(String::from("Device replied with no response"))?
            .into())
    }
    pub async fn request_pair(
        &mut self,
        this_info: storage::DeviceInfo,
        this_certificate: security::Certificate,
    ) -> Result<
        Result<
            (
                security::PairToken,
                security::Certificate,
                storage::DeviceInfo,
            ),
            String,
        >,
        Box<dyn std::error::Error>,
    > {
        let response = self
            .pair_client
            .request_pair(proto::RequestPairRequest {
                info: Some(this_info.into()),
                certificate: this_certificate,
            })
            .await?
            .into_inner();
        match response.result {
            Some(result) => Ok(match result {
                proto::request_pair_response::Result::Success(success_response) => {
                    _ = 5;
                    Ok((
                        success_response.token,
                        success_response.certificate,
                        success_response
                            .info
                            .ok_or_else(|| {
                                ClientError::MissingItem(String::from("success_response.info"))
                            })?
                            .into(),
                    ))
                }
                proto::request_pair_response::Result::Failure(failure_response) => {
                    Err(failure_response.reason)
                }
            }),
            None => Err(ClientError::EmptyResponse.into()),
        }
    }
}
