use super::{proto, Certificate, DeviceInfo, StorageError, StorageResult};
#[derive(Debug, serde::Deserialize, serde::Serialize)]
pub struct PairedDevice {
    token: Vec<u8>,
    certificate: Certificate,
    info: DeviceInfo,
}
impl From<proto::PairedDevice> for StorageResult<PairedDevice> {
    fn from(pd: proto::PairedDevice) -> Self {
        Ok(PairedDevice {
            token: match pd.token {
                Some(tk) => tk.data,
                None => {
                    return Err(StorageError::PbufConvertError(String::from(
                        "paired device record has no token",
                    )))
                }
            },
            certificate: match pd.certificate {
                Some(c) => c.into(),
                None => {
                    return Err(StorageError::PbufConvertError(String::from(
                        "paired device record has no certificate",
                    )))
                }
            },
            info: match pd.info {
                Some(i) => i.into(),
                None => {
                    return Err(StorageError::PbufConvertError(String::from(
                        "paired device record has no device info",
                    )))
                }
            },
        })
    }
}
impl From<PairedDevice> for proto::PairedDevice {
    fn from(pd: PairedDevice) -> Self {
        proto::PairedDevice {
            token: Some(proto::PairToken { data: pd.token }),
            certificate: Some(pd.certificate.into()),
            info: Some(pd.info.into()),
        }
    }
}
