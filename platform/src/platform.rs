use super::discovery;

pub struct Platform {}

impl Platform {
    pub fn init() {}
}

impl nlem::platform::Platform<discovery::DiscoveryManager> for Platform {
    fn hostname(&self) -> Result<String, String> {
        hostname::get()
            .map_err(|e| e.to_string())?
            .into_string()
            .map_err(|e| format!("Could not convert osstring to string: {e:#?}"))
    }
    fn device_type(&self) -> nlem::storage::DeviceType {
        nlem::storage::DeviceType::Desktop
    }
    async fn discovery_manager(&self) -> Result<discovery::DiscoveryManager, String> {
        discovery::DiscoveryManager::init().await
    }
}
