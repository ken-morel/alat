use std::sync::Arc;
use tokio::sync::RwLock;

use super::{client, devicemanager, security, server, storage};

pub struct Node {
    pub storage: crate::StorageC,
    pub platform: crate::PlatformC,
    pub device_manager: crate::DeviceManagerC,
    pub server: crate::ServerC,
}

impl Node {
    pub async fn init(
        platform: crate::PlatformC,
    ) -> Result<Self, Box<dyn std::error::Error + Send + Sync>> {
        let storage = platform.read().await.storage().await?;
        storage
            .write()
            .await
            .init(storage::StorageData {
                certificate: security::generate_certificate(),
                paired_devices: Vec::new(),
                info: Self::default_device_info(&*platform.read().await).await,
            })
            .await;
        let discovery = platform.write().await.discovery_manager().await?;
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
    pub async fn default_device_info(p: &crate::Platform) -> storage::DeviceInfo {
        storage::DeviceInfo {
            id: security::generate_id(),
            color: storage::Color::random(),
            name: p.hostname().await.expect("Could not get hostname"),
            device_type: p.device_type().await,
        }
    }
    pub async fn start(
        &mut self,
    ) -> tokio::sync::mpsc::Receiver<devicemanager::DeviceManagerEvent> {
        let router = self.server.write().await.create_router();
        tokio::spawn(async move {
            let addr = std::net::SocketAddr::new(
                std::net::Ipv4Addr::UNSPECIFIED.into(),
                server::ALAT_PORT,
            );
            println!("[node/server] Starting server at {addr}");
            let r = router.serve(addr).await;
            println!("[node/server] Server at {addr} stopped");
            if let Err(e) = r {
                println!("[node/server::error] {e}");
            }
        });

        let (tx, rx) = tokio::sync::mpsc::channel(1);
        self.device_manager.write().await.start_workers(tx).await;
        rx
    }

    pub async fn request_pair(
        &self,
        device_id: &security::DeviceID,
    ) -> Result<Result<storage::PairedDevice, String>, String> {
        let manager = self.device_manager.read().await;
        let device = manager
            .discovered_devices
            .read()
            .await
            .get(device_id)
            .ok_or(String::from("Device not found"))?
            .clone();
        let this_info = manager.this_device.read().await.clone().info;
        let this_certificate = manager.device_certificate.read().await.clone();
        drop(manager);

        let mut cl = client::Client::connect(device.address).await.map_err(|e| {
            format!(
                "Client could not connect to device at {0}: {e}",
                device.address
            )
        })?;

        match cl.request_pair(this_info, this_certificate).await {
            Ok(response) => match response {
                Ok((token, certificate, info)) => {
                    let paired_device = storage::PairedDevice {
                        token,
                        certificate,
                        info,
                    };
                    self.device_manager
                        .read()
                        .await
                        .add_paired_device(paired_device.clone())
                        .await;
                    Ok(Ok(paired_device))
                }
                Err(message) => Ok(Err(message)),
            },
            Err(err) => Err(format!("Could not send pair request: {err}")),
        }
    }
}
