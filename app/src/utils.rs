use super::ui;

pub fn slint_col(col: &nlem::storage::Color) -> slint::Color {
    slint::Color::from_rgb_u8(col.0, col.1, col.2)
}

#[derive(Debug, Clone)]
pub enum DeviceRelationship {
    Paired,
    Connected,
    Found,
}
impl From<DeviceRelationship> for ui::DeviceRelationship {
    fn from(value: DeviceRelationship) -> Self {
        match value {
            DeviceRelationship::Paired => ui::DeviceRelationship::Paired,
            DeviceRelationship::Connected => ui::DeviceRelationship::Connected,
            DeviceRelationship::Found => ui::DeviceRelationship::Found,
        }
    }
}
unsafe impl Sync for DeviceRelationship {}
unsafe impl Send for DeviceRelationship {}

#[derive(Debug, Clone)]
pub struct Device {
    pub name: String,
    pub color: nlem::storage::Color,
    pub address: String,
    pub port: i32,
    pub id: nlem::security::DeviceID,
    pub relationship: DeviceRelationship,
}

unsafe impl Sync for Device {}

unsafe impl Send for Device {}
