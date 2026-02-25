pub mod config;
pub mod error;

pub mod clipboard;
pub mod pair;
pub mod telemetry;

use std::collections::HashMap;
use tokio::sync::RwLock;

#[derive(Debug, Clone)]
pub enum ServiceEvent {
    Started(ServiceID),
    Stopped(ServiceID),

    Error(ServiceID, error::ServiceError),
}

#[derive(Debug, Clone)]
pub enum ServiceManagerEvent {
    ServiceEvent(ServiceEvent),

    Started,
    Stopped,
}

pub type ServiceID = &'static str;
pub type ServiceChannel = tokio::sync::mpsc::Sender<ServiceEvent>;
pub type SpawnWorkerResult = Option<tokio::task::JoinHandle<()>>;

#[tonic::async_trait]
pub trait Service: Send + Sync {
    fn name(&self) -> ServiceID;
    async fn init(&mut self, node: crate::Node) -> Result<(), error::ServiceError>;
    async fn spawn_worker(&self, channel: ServiceChannel) -> SpawnWorkerResult;
    async fn grpc(
        &self,
        server: tonic::transport::server::Router,
    ) -> Result<tonic::transport::server::Router, error::ServiceError>;
    fn is_init(&self) -> bool;
    fn ensure_init(&self) -> Result<(), error::ServiceError> {
        if !self.is_init() {
            Err(error::ServiceError::NotInitialized(self.name()))
        } else {
            Ok(())
        }
    }
}

#[derive(Default)]
pub struct ServiceManager {
    pub services: RwLock<HashMap<ServiceID, crate::ServiceC>>,
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
            server = service.write().await.grpc(server).await?;
        }
        Ok(server)
    }
    pub async fn init(&mut self, node: crate::Node) -> Result<(), crate::ErrorC> {
        for service in self.services.read().await.values() {
            service.write().await.init(node.clone()).await?;
        }
        Ok(())
    }
    pub async fn start(
        &mut self,
    ) -> Result<tokio::sync::mpsc::Receiver<ServiceManagerEvent>, crate::ErrorC> {
        let (stx, mut srx) = tokio::sync::mpsc::channel(1);
        let (tx, rx) = tokio::sync::mpsc::channel(1);
        let send = async move |e: ServiceManagerEvent| {
            tx.send(e)
                .await
                .expect("Could not relay message to main service manage channel");
        };
        let services = self.services.read().await;
        for service in services.values() {
            let sender = stx.clone();
            let service = service.clone();
            tokio::spawn(async move {
                service.write().await.spawn_worker(sender).await;
            });
        }

        tokio::spawn(async move {
            send(ServiceManagerEvent::Started).await;
            while let Some(event) = srx.recv().await {
                match &event {
                    ServiceEvent::Started(_)
                    | ServiceEvent::Stopped(_)
                    | ServiceEvent::Error(_, _) => {
                        send(ServiceManagerEvent::ServiceEvent(event)).await;
                    }
                };
            }
            send(ServiceManagerEvent::Stopped).await;
        });

        Ok(rx)
    }
}
