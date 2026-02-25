use super::storage;

use nlem::service;
use std::{path::PathBuf, sync::Arc};
use tokio::sync::{Mutex, RwLock};

use super::discovery;

pub struct Platform {}

impl Platform {
    pub fn init() -> Self {
        Self {}
    }
}

impl Platform {
    async fn config_dir(&self) -> Result<PathBuf, String> {
        match &mut dirs::config_dir() {
            Some(path) => {
                path.push(nlem::APP_ID);
                Ok(path.clone())
            }
            None => Err(String::from("Could not get application config dir")),
        }
    }
}

#[tonic::async_trait]
impl nlem::platform::Platform for Platform {
    async fn hostname(&self) -> Result<String, String> {
        hostname::get()
            .map_err(|e| e.to_string())?
            .into_string()
            .map_err(|e| format!("Could not convert osstring to string: {e:#?}"))
    }
    async fn device_type(&self) -> nlem::storage::DeviceType {
        nlem::storage::DeviceType::Desktop
    }
    async fn discovery_manager(&self) -> Result<nlem::DiscoveryC, String> {
        Ok(Arc::new(RwLock::new(
            discovery::DiscoveryManager::init().await?,
        )))
    }
    async fn storage(&self) -> Result<nlem::StorageC, String> {
        let mut cfg_path = self.config_dir().await?;
        cfg_path.push("data.json");
        Ok(Arc::new(Mutex::new(storage::JSONFileStorage::new(
            cfg_path.as_path(),
        ))))
    }
    async fn prompt_pair_request(
        &self,
        info: nlem::storage::DeviceInfo,
        _certificate: nlem::security::Certificate,
    ) -> Result<(), String> {
        let (tx, rx) = tokio::sync::oneshot::channel();
        tokio::task::spawn_blocking(move || {
            notify_rust::Notification::new()
                .summary("Pair request")
                .body(&format!("Device named {} wants to pair", info.name))
                .action("y", "Pair")
                .action("n", "Decline")
                .show()
                .unwrap()
                .wait_for_action(|action| {
                    if let Err(e) = match action {
                        "y" => tx.send(Ok(())),
                        "n" => tx.send(Err(String::from("User declined"))),
                        _ => tx.send(Err(String::from("No notification reply"))),
                    } {
                        println!("ERROR sending notification reply to oneshot channel: {e:?}");
                    }
                });
        });
        rx.await
            .map_err(|e| format!("Could not ama receive noification reply: {e}"))?
    }
    async fn log_info(&self, msg: String) {
        log::info!("{msg}");
    }
    async fn log_error(&self, msg: String) {
        log::error!("{msg}");
    }
    async fn log_warning(&self, msg: String) {
        log::warn!("{msg}");
    }
    async fn query_telemetry(&self) -> Result<service::telemetry::TelemetryInfo, String> {
        let mut info = service::telemetry::TelemetryInfo::default();
        crate::telemetry::collect_info(&mut info).await?;
        Ok(info)
    }

    async fn clipboard(&self) -> Result<Box<dyn service::clipboard::Clipboard>, String> {
        Box::new(crate::clipboard::Clipboard::init()?)
    }
}
