pub mod config;
pub mod error;

use std::{collections::HashMap, fmt::Display};
use tokio::sync::RwLock;

use crate::{devicemanager::connected, proto, security};

#[derive(Debug, Clone)]
pub enum UnitEvent {
    Started(UnitID),
    Stopped(UnitID),

    Error(UnitID, error::UnitError),
}

#[derive(Debug, Clone)]
pub enum UnitManagerEvent {
    UnitEvent(UnitEvent),

    Started,
    Stopped,
}

#[derive(PartialEq, Debug, Eq, Clone, Hash)]
pub enum UnitID {
    Service(&'static str),
    Controller(&'static str),
}
impl Display for UnitID {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.write_str(&match self {
            Self::Service(n) => format!("service:{n}"),
            Self::Controller(n) => format!("controller:{n}"),
        })
    }
}

pub type UnitSender = tokio::sync::mpsc::Sender<UnitEvent>;
pub type SpawnWorkerResult = Option<tokio::task::JoinHandle<()>>;

#[tonic::async_trait]
pub trait Unit: Send + Sync {
    fn id(&self) -> UnitID;
    async fn init(&mut self) -> Result<(), error::UnitError>;
    async fn spawn_worker(&self, channel: UnitSender) -> SpawnWorkerResult;
    async fn grpc(
        &self,
        server: tonic::transport::server::Router,
    ) -> Result<tonic::transport::server::Router, error::UnitError>;
    async fn authenticate(
        &self,
        man: &crate::DeviceManager,
        req: &proto::ServiceCall,
    ) -> error::UnitResult<connected::ConnectedDevice> {
        if let Some(auth) = &req.auth {
            if let Some(dev) = man
                .get_connected_device_by_token(&security::array_from_vec(auth.token.clone()))
                .await
            {
                return Ok(dev);
            }
        }
        Err(error::UnitError::Unauthenticated())
    }
    fn is_init(&self) -> bool;
    fn ensure_init(&self) -> Result<(), error::UnitError> {
        if !self.is_init() {
            Err(error::UnitError::NotInitialized(self.id()))
        } else {
            Ok(())
        }
    }
}

#[derive(Default)]
pub struct UnitManager {
    pub units: RwLock<HashMap<UnitID, crate::UnitC>>,
}

impl UnitManager {
    pub fn new() -> Self {
        Self::default()
    }
    pub async fn add_unit(&mut self, s: crate::UnitC) {
        self.units
            .write()
            .await
            .insert(s.clone().read().await.id(), s);
    }
    pub async fn get_unit(&self, id: &UnitID) -> Option<crate::UnitC> {
        self.units.read().await.get(&id).cloned()
    }
    pub async fn register_grpc_unit_servers(
        &self,
        mut server: tonic::transport::server::Router,
    ) -> Result<tonic::transport::server::Router, error::UnitError> {
        for unit in self.units.read().await.values() {
            server = unit.write().await.grpc(server).await?;
        }
        Ok(server)
    }
    pub async fn init(&mut self) -> Result<(), crate::ErrorC> {
        for unit in self.units.read().await.values() {
            unit.write().await.init().await?;
        }
        Ok(())
    }
    pub async fn start(
        &mut self,
    ) -> Result<tokio::sync::mpsc::Receiver<UnitManagerEvent>, crate::ErrorC> {
        let (stx, mut srx) = tokio::sync::mpsc::channel(1);
        let (tx, rx) = tokio::sync::mpsc::channel(1);
        let send = async move |e: UnitManagerEvent| {
            tx.send(e)
                .await
                .expect("Could not relay message to main unit manage channel");
        };
        let units = self.units.read().await;
        for unit in units.values() {
            let sender = stx.clone();
            let unit = unit.clone();
            tokio::spawn(async move {
                unit.write().await.spawn_worker(sender).await;
            });
        }

        tokio::spawn(async move {
            send(UnitManagerEvent::Started).await;
            while let Some(event) = srx.recv().await {
                match &event {
                    UnitEvent::Started(_) | UnitEvent::Stopped(_) | UnitEvent::Error(_, _) => {
                        send(UnitManagerEvent::UnitEvent(event)).await;
                    }
                };
            }
            send(UnitManagerEvent::Stopped).await;
        });

        Ok(rx)
    }
}
