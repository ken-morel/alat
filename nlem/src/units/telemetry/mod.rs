mod config;
mod info;
mod server;

pub type TelemetryInfo = info::TelemetryInfo;

use crate::{proto, unit};

#[derive(Clone)]
pub struct TelemetryService {
    initialized: bool,
    node: crate::Node,
    config: Option<crate::RWContainer<config::TelemetryServiceConfig>>,
    storage: Option<crate::RWContainer<unit::config::UnitConfig>>,
    info: crate::RWContainer<Option<info::TelemetryInfo>>,
}

impl TelemetryService {
    pub fn new(node: crate::Node) -> Self {
        Self {
            node,
            initialized: false,
            config: None,
            storage: None,
            info: crate::contain(None),
        }
    }
}

#[tonic::async_trait]
impl unit::Unit for TelemetryService {
    fn is_init(&self) -> bool {
        self.initialized
    }
    fn id(&self) -> unit::UnitID {
        unit::UnitID::Service("telemetry")
    }
    async fn init(&mut self) -> unit::error::UnitResult<()> {
        let name = self.id();
        let mut storage = unit::config::UnitConfig::new(self.node.storage.clone(), self.id());
        self.config = Some(crate::contain(
            storage
                .init(config::TelemetryServiceConfig::default())
                .await
                .map_err(|e| unit::error::UnitError::StorageError(name, e))?,
        ));
        self.storage = Some(crate::contain(storage));
        self.initialized = true;
        Ok(())
    }
    async fn spawn_worker(&self, channel: unit::UnitSender) -> unit::SpawnWorkerResult {
        let send = async move |msg: unit::UnitEvent| {
            channel
                .send(msg)
                .await
                .expect("COuld not relay message to servicechannel");
        };
        if let Err(e) = self.ensure_init() {
            send(unit::UnitEvent::Error(self.id(), e)).await;
            return None;
        }
        let name = self.id();
        let info = self.info.clone();
        let config = self.config.clone().unwrap();
        let platform = self.node.clone().platform;

        Some(tokio::spawn(async move {
            send(unit::UnitEvent::Started(name.clone())).await;
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
            send(unit::UnitEvent::Stopped(name)).await;
        }))
    }
    async fn grpc(
        &self,
        server: tonic::transport::server::Router,
    ) -> Result<tonic::transport::server::Router, unit::error::UnitError> {
        let server = server.add_service(
            proto::telemetry_service_server::TelemetryServiceServer::new(self.clone()),
        );
        Ok(server)
    }
}
