pub mod clipboard;
pub mod filesystem;
pub mod pair;
pub mod telemetry;

pub async fn register_units(node: &crate::Node) -> Result<(), crate::ErrorC> {
    use std::sync::Arc;
    for unit in vec![
        Arc::new(pair::PairService::new(node.clone())) as crate::UnitC,
        Arc::new(clipboard::ClipboardService::new(node.clone())) as crate::UnitC,
        Arc::new(telemetry::TelemetryService::new(node.clone())) as crate::UnitC,
    ] {
        node.unit_register(unit).await?;
    }
    Ok(())
}
