use super::{proto, storage};
use std::net::SocketAddr;

#[derive(Debug, Clone)]
pub struct Client {
    alat_client: proto::alat_service_client::AlatServiceClient<tonic::transport::Channel>,
    server_addr: SocketAddr,
}

impl Client {
    pub async fn connect(addr: SocketAddr) -> Result<Self, String> {
        let client = proto::alat_service_client::AlatServiceClient::connect(addr.to_string())
            .await
            .map_err(|e| e.to_string())?;
        Ok(Self {
            server_addr: addr,
            alat_client: client,
        })
    }
    pub async fn reconnect(self) -> Result<Self, String> {
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
}
