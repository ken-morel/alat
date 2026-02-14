use nlem::devicemanager::discovered;

pub struct DiscoveryManager {}

impl DiscoveryManager {
    pub async fn init() -> Result<Self, String> {
        Ok(Self {})
    }
}

impl discovered::DiscoveryManager for DiscoveryManager {
    async fn advertise(
        &mut self,
        _: discovered::DiscoveredDevice,
    ) -> Result<(), discovered::DiscoveryError> {
        Ok(())
    }
    async fn cease_advertising(&mut self) -> Result<(), discovered::DiscoveryError> {
        Ok(())
    }

    async fn scan(
        &mut self,
        _: tokio::sync::mpsc::Sender<discovered::DiscoveryEvent>,
    ) -> Result<(), discovered::DiscoveryError> {
        Ok(())
    }

    async fn cease_scan(&mut self) -> Result<(), discovered::DiscoveryError> {
        Ok(())
    }
}
