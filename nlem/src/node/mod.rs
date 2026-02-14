use std::sync::Arc;
use tokio::sync::RwLock;

use super::{devicemanager, platform, server, storage};

pub struct Node<
    S: storage::Storage + 'static,
    P: platform::Platform<S, D> + 'static,
    D: devicemanager::discovered::DiscoveryManager + 'static,
> {
    pub storage: Arc<RwLock<S>>,
    pub platform: Arc<RwLock<P>>,
    pub device_manager: Arc<RwLock<devicemanager::DeviceManager<S, P, D>>>,
    pub server: Arc<RwLock<server::Server<S, P, D>>>,
}

impl<
    S: storage::Storage,
    P: platform::Platform<S, D>,
    D: devicemanager::discovered::DiscoveryManager,
> Node<S, P, D>
{
    pub async fn init(
        platform: Arc<RwLock<P>>,
    ) -> Result<Self, Box<dyn std::error::Error + Send + Sync>> {
        let storage = Arc::new(RwLock::new(platform.read().await.storage().await?));
        let discovery = Arc::new(RwLock::new(
            platform.write().await.discovery_manager().await?,
        ));
        let device_manager = Arc::new(RwLock::new(
            devicemanager::DeviceManager::init(
                storage.clone(),
                platform.clone(),
                discovery.clone(),
            )
            .await?,
        ));
        let server = Arc::new(RwLock::new(server::Server::new(device_manager.clone())));
        Ok(Self {
            storage,
            platform,
            device_manager,
            server,
        })
    }
    pub async fn start(
        &mut self,
    ) -> tokio::sync::mpsc::Receiver<devicemanager::DeviceManagerEvent> {
        tokio::spawn(
            self.server
                .write()
                .await
                .create_router()
                .serve(std::net::SocketAddr::new(
                    std::net::Ipv4Addr::LOCALHOST.into(),
                    server::ALAT_PORT,
                )),
        );
        let (tx, rx) = tokio::sync::mpsc::channel(0);
        // devicemanager::DeviceManagerEvent
        self.device_manager.write().await.start_workers(tx).await;
        rx
    }
}
