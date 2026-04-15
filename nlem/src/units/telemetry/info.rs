use crate::proto;
#[derive(serde::Serialize, serde::Deserialize, Clone, Default)]
pub struct TelemetryInfo {
    pub uptime_secs: u64,
    pub os_name: String,
    pub os_version: String,
    pub kernel_version: String,
    pub hostname: String,
    pub cpu_usages: Vec<f32>,
    pub disk_names: Vec<String>,
    pub disk_used_spaces_gb: Vec<f32>,
    pub disk_total_spaces_gb: Vec<f32>,
    pub batteries_charging: Vec<bool>,
    pub batteries_percent: Vec<u32>,
    pub memory_used_mb: u32,
    pub memory_total_mb: u32,
    pub swap_used_mb: u32,
    pub swap_total_mb: u32,
}

impl From<TelemetryInfo> for proto::TelemetryStatus {
    fn from(value: TelemetryInfo) -> Self {
        Self { ..value.into() }
    }
}

impl From<proto::TelemetryStatus> for TelemetryInfo {
    fn from(value: proto::TelemetryStatus) -> Self {
        Self { ..value.into() }
    }
}
