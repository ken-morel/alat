pub mod error;

pub mod pair;

use std::collections::HashMap;

use tokio::sync::RwLock;

type ServiceID = &'static str;

#[tonic::async_trait]
pub trait Service: Send + Sync {
    fn name(&self) -> ServiceID;
    async fn init(&mut self, node: crate::Node) -> Result<(), crate::ErrorC>;
    async fn worker(&mut self) -> Result<(), error::ServiceError>;
    async fn register_grpc_service_server(
        &self,
        server: tonic::transport::server::Router,
    ) -> Result<tonic::transport::server::Router, error::ServiceError>;
}

#[derive(Default)]
pub struct ServiceManager {
    services: RwLock<HashMap<ServiceID, crate::ServiceC>>,
}

impl ServiceManager {
    pub fn new() -> Self {
        Self::default()
    }
    pub async fn add_service(&mut self, s: crate::ServiceC) {
        self.services
            .write()
            .await
            .insert(s.clone().read().await.name(), s);
    }
    pub async fn get_service(&self, name: &str) -> Option<crate::ServiceC> {
        self.services.read().await.get(name).cloned()
    }
    pub async fn register_grpc_service_servers(
        &self,
        mut server: tonic::transport::server::Router,
    ) -> Result<tonic::transport::server::Router, error::ServiceError> {
        for service in self.services.read().await.values() {
            server = service
                .write()
                .await
                .register_grpc_service_server(server)
                .await?;
        }
        Ok(server)
    }
    pub async fn initialize_services(&mut self, node: crate::Node) -> Result<(), crate::ErrorC> {
        for service in self.services.read().await.values() {
            service.write().await.init(node.clone()).await?;
        }
        Ok(())
    }
}
