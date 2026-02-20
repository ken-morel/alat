use crate::proto;
#[derive(serde::Serialize, serde::Deserialize, Clone, Default)]
pub struct TelemetryInfo {
    pub uptime_secs: u64,
    pub os_name: String,
    pub os_version: String,
    pub kernel_version: String,
    pub hostname: String,
    pub cpu_usage_per_core: Vec<f32>,
    pub disk_names: Vec<String>,
    pub disk_available_spaces_gb: Vec<f32>,
    pub disk_total_spaces_gb: Vec<f32>,
    pub batteries_changing: Vec<bool>,
    pub batteries_percent: Vec<u32>,
    pub memory_used: u32,
    pub memory_total: u32,
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
