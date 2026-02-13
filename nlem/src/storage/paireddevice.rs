use super::{Certificate, DeviceInfo, StorageError, StorageResult, proto};
#[derive(Debug, Clone, serde::Deserialize, serde::Serialize)]
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
                    return Err(StorageError::PbufConvert(String::from(
                        "paired device record has no token",
                    )));
                }
            },
            certificate: pd.certificate,
            info: match pd.info {
                Some(i) => i.into(),
                None => {
                    return Err(StorageError::PbufConvert(String::from(
                        "paired device record has no device info",
                    )));
                }
            },
        })
    }
}
impl From<PairedDevice> for proto::PairedDevice {
    fn from(pd: PairedDevice) -> Self {
        proto::PairedDevice {
            token: Some(proto::PairToken { data: pd.token }),
            certificate: pd.certificate,
            info: Some(pd.info.into()),
        }
    }
}
