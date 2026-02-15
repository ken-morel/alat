use crate::security;

use super::{Certificate, DeviceInfo, StorageError, StorageResult, proto};
#[derive(Debug, Clone, serde::Deserialize, serde::Serialize)]
pub struct PairedDevice {
    pub token: security::PairToken,
    pub certificate: Certificate,
    pub info: DeviceInfo,
}
impl From<proto::PairedDevice> for StorageResult<PairedDevice> {
    fn from(pd: proto::PairedDevice) -> Self {
        Ok(PairedDevice {
            token: security::array_from_vec(pd.token),
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
            token: pd.token.to_vec(),
            certificate: pd.certificate,
            info: Some(pd.info.into()),
        }
    }
}
