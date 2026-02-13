use super::{Color, DeviceType, proto};

#[derive(Debug, Clone, serde::Deserialize, serde::Serialize)]
pub struct DeviceInfo {
    pub id: Vec<u8>,
    pub name: String,
    pub color: Color,
    pub device_type: DeviceType,
}
impl From<proto::DeviceInfo> for DeviceInfo {
    fn from(inf: proto::DeviceInfo) -> Self {
        Self {
            device_type: inf.device_type().into(),
            id: inf.id,
            color: inf
                .color
                .unwrap_or(proto::Color { r: 0, g: 0, b: 0 })
                .into(),
            name: inf.name,
        }
    }
}
impl From<DeviceInfo> for proto::DeviceInfo {
    fn from(inf: DeviceInfo) -> Self {
        Self {
            device_type: proto::DeviceType::from(inf.device_type).into(),
            id: inf.id,
            color: Some(inf.color.into()),
            name: inf.name,
        }
    }
}
