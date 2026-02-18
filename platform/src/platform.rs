use super::storage;
use std::path::PathBuf;

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
impl nlem::platform::Platform<storage::JSONFileStorage, discovery::DiscoveryManager> for Platform {
    fn hostname(&self) -> Result<String, String> {
        hostname::get()
            .map_err(|e| e.to_string())?
            .into_string()
            .map_err(|e| format!("Could not convert osstring to string: {e:#?}"))
    }
    fn device_type(&self) -> nlem::storage::DeviceType {
        nlem::storage::DeviceType::Desktop
    }
    async fn discovery_manager(&self) -> Result<discovery::DiscoveryManager, String> {
        discovery::DiscoveryManager::init().await
    }
    async fn storage(&self) -> Result<storage::JSONFileStorage, String> {
        let mut cfg_path = self.config_dir().await?;
        cfg_path.push("data.json");
        Ok(storage::JSONFileStorage::new(cfg_path.as_path()))
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
            .map_err(|e| format!("Could not receive noification reply: {e}"))?
    }
}
