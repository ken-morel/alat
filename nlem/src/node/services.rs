use std::sync::Arc;
use tokio::sync::RwLock;

use crate::service;

pub async fn register_services(
    services: &mut service::ServiceManager,
) -> Result<(), crate::ErrorC> {
    services
        .add_service(Arc::new(RwLock::new(service::pair::PairService::new())))
        .await;
    Ok(())
}
