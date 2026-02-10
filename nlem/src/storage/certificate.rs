use super::proto;
#[derive(Debug, serde::Deserialize, serde::Serialize)]
pub struct Certificate {
    data: Vec<u8>,
}
impl From<proto::Certificate> for Certificate {
    fn from(value: proto::Certificate) -> Self {
        Self {
            data: prost::Message::encode_to_vec(&value),
        }
    }
}
impl From<Certificate> for proto::Certificate {
    fn from(cert: Certificate) -> Self {
        proto::Certificate { data: cert.data }
    }
}
