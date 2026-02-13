use crate::storage;

pub trait Platform: Send + Sync {
    fn hostname(&self) -> Result<String, String>;
    fn device_type(&self) -> storage::DeviceType;
}
