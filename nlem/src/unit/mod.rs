pub mod config;
pub mod error;

use std::fmt::Display;

use crate::{devicemanager::DeviceManager, devicemanager::connected, proto, security};

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
pub trait Unit: Send + Sync + std::fmt::Debug {
    fn id(&self) -> UnitID;
    async fn init(&self) -> Result<(), error::UnitError>;
    async fn spawn_worker(&self, channel: UnitSender) -> SpawnWorkerResult;
    async fn grpc(
        &self,
        server: tonic::transport::server::Router,
    ) -> Result<tonic::transport::server::Router, error::UnitError>;
    async fn authenticate(
        &self,
        man: &DeviceManager,
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
    async fn is_init(&self) -> bool;
    async fn ensure_init(&self) -> Result<(), error::UnitError> {
        if !self.is_init().await {
            Err(error::UnitError::NotInitialized(self.id()))
        } else {
            Ok(())
        }
    }
}

#[derive(Default, Debug)]
pub struct UnitManager {
    pub units: dashmap::DashMap<UnitID, crate::UnitC>,
    pub initialized: tokio::sync::RwLock<bool>,
}
impl UnitManager {
    pub fn new() -> Self {
        Self::default()
    }
    pub async fn add_unit(&self, s: crate::UnitC) -> Result<(), crate::ErrorC> {
        let init = self.initialized.read().await;
        if *init {
            s.init().await?;
        }
        let mut inserted = false;

        let id = s.id();
        self.units.entry(s.id()).or_insert_with(|| {
            inserted = true;
            s
        });
        if !inserted {
            return Err(Box::new(error::UnitError::UnitAlreadyRegistered(id)));
        }
        Ok(())
    }
    pub async fn get_unit(&self, id: &UnitID) -> Option<crate::UnitC> {
        match self.units.get(&id) {
            Some(u) => Some(u.clone()),
            None => None,
        }
    }
    pub async fn register_grpc_unit_servers(
        &self,
        mut server: tonic::transport::server::Router,
    ) -> Result<tonic::transport::server::Router, error::UnitError> {
        let units: Vec<_> = self.units.iter().map(|u| u.value().clone()).collect();
        for unit in units {
            server = unit.grpc(server).await?;
        }
        Ok(server)
    }
    pub async fn init(&self) -> Result<(), crate::ErrorC> {
        let mut init = self.initialized.write().await;
        if !*init {
            let units: Vec<_> = self.units.iter().map(|u| u.value().clone()).collect();
            for unit in units {
                unit.init().await?;
            }
        }
        *init = true;
        Ok(())
    }
    pub async fn start(
        &self,
    ) -> Result<tokio::sync::mpsc::Receiver<UnitManagerEvent>, crate::ErrorC> {
        let (stx, mut srx) = tokio::sync::mpsc::channel(1);
        let (tx, rx) = tokio::sync::mpsc::channel(1);
        let send = async move |e: UnitManagerEvent| {
            tx.send(e)
                .await
                .expect("Could not relay message to main unit manage channel");
        };
        let units: Vec<_> = self.units.iter().map(|u| u.value().clone()).collect();
        for unit in units {
            let sender = stx.clone();
            tokio::spawn(async move {
                unit.spawn_worker(sender).await;
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
