#[derive(serde::Deserialize, serde::Serialize)]
pub struct TelemetryServiceConfig {
    enabled: bool,
}
impl Default for TelemetryServiceConfig {
    fn default() -> Self {
        Self { enabled: true }
    }
}
