#[derive(serde::Deserialize, serde::Serialize, Clone)]
pub struct TelemetryServiceConfig {
    pub enabled: bool,
    pub poll_interval_secs: u32,
}
impl Default for TelemetryServiceConfig {
    fn default() -> Self {
        Self {
            enabled: true,
            poll_interval_secs: 10,
        }
    }
}
