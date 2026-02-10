use super::proto;
#[derive(Debug, serde::Deserialize, serde::Serialize)]
pub enum DeviceType {
    Unspecified = 0,
    Mobile = 1,
    Desktop = 2,
    Tv = 3,
    Embedded = 4,
}
impl From<proto::DeviceType> for DeviceType {
    fn from(t: proto::DeviceType) -> Self {
        match t {
            proto::DeviceType::Unspecified => Self::Unspecified,
            proto::DeviceType::Mobile => Self::Mobile,
            proto::DeviceType::Desktop => Self::Desktop,
            proto::DeviceType::Tv => Self::Tv,
            proto::DeviceType::Embedded => Self::Embedded,
        }
    }
}
impl From<DeviceType> for proto::DeviceType {
    fn from(t: DeviceType) -> Self {
        match t {
            DeviceType::Unspecified => Self::Unspecified,
            DeviceType::Mobile => Self::Mobile,
            DeviceType::Desktop => Self::Desktop,
            DeviceType::Tv => Self::Tv,
            DeviceType::Embedded => Self::Embedded,
        }
    }
}
