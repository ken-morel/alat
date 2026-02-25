mod config;
mod info;
mod server;

pub type TelemetryInfo = info::TelemetryInfo;

use crate::proto;

#[derive(Default, Clone)]
pub struct TelemetryService {
    initialized: bool,
    node: Option<crate::Node>,
    config: Option<crate::RWContainer<config::TelemetryServiceConfig>>,
    storage: Option<crate::RWContainer<super::config::ServiceConfig>>,
    info: Option<crate::RWContainer<Option<info::TelemetryInfo>>>,
}

impl TelemetryService {
    pub fn new() -> Self {
        Self::default()
    }
}

#[tonic::async_trait]
impl super::Service for TelemetryService {
    fn is_init(&self) -> bool {
        self.initialized
    }
    fn name(&self) -> super::ServiceID {
        "telemetry"
    }
    async fn init(&mut self, node: crate::Node) -> super::error::ServiceResult<()> {
        let name = self.name();
        let mut storage = super::config::ServiceConfig::new(node.storage.clone(), self.name());
        self.config = Some(crate::contain(
            storage
                .init(config::TelemetryServiceConfig::default())
                .await
                .map_err(|e| super::error::ServiceError::StorageError(name, e))?,
        ));
        self.storage = Some(crate::contain(storage));
        self.node = Some(node);
        self.info = Some(crate::contain(None));
        self.initialized = true;
        Ok(())
    }
    async fn spawn_worker(&self, channel: super::ServiceChannel) -> super::SpawnWorkerResult {
        let send = async move |msg: super::ServiceEvent| {
            channel
                .send(msg)
                .await
                .expect("COuld not relay message to servicechannel");
        };
        if let Err(e) = self.ensure_init() {
            send(super::ServiceEvent::Error(self.name(), e)).await;
            return None;
        }
        let name = self.name();
        let info = self.info.clone().unwrap();
        let config = self.config.clone().unwrap();
        let platform = self.node.clone().unwrap().platform;

        Some(tokio::spawn(async move {
            send(super::ServiceEvent::Started(name)).await;
            loop {
                match platform.read().await.query_telemetry().await {
                    Ok(data) => {
                        *info.write().await = Some(data);
                    }
                    Err(e) => {
                        println!("[service/telemetry] ERROR {e}");
                    }
                }

                tokio::time::sleep(std::time::Duration::from_secs(
                    config.read().await.poll_interval_secs.into(),
                ))
                .await;
                if false {
                    break;
                }
            }
            send(super::ServiceEvent::Stopped(name)).await;
        }))
    }
    async fn grpc(
        &self,
        server: tonic::transport::server::Router,
    ) -> Result<tonic::transport::server::Router, super::error::ServiceError> {
        let server = server.add_service(
            proto::telemetry_service_server::TelemetryServiceServer::new(self.clone()),
        );
        Ok(server)
    }
}
