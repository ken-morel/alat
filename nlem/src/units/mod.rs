pub mod clipboard;
pub mod pair;
pub mod telemetry;

pub async fn register_units(node: &crate::Node) {
    node.unit_register(std::sync::Arc::new(pair::PairService::new(node.clone())))
        .await;
}
