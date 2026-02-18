mod ui;
use std::sync::Arc;
use tokio::sync::RwLock;

use slint::ComponentHandle;

pub fn slint_col(col: &nlem::storage::Color) -> slint::Color {
    slint::Color::from_rgb_u8(col.0, col.1, col.2)
}

#[derive(Debug, Clone)]
enum DeviceRelationship {
    Paired,
    Connected,
    Found,
}
impl From<DeviceRelationship> for ui::DeviceRelationship {
    fn from(value: DeviceRelationship) -> Self {
        match value {
            DeviceRelationship::Paired => ui::DeviceRelationship::Paired,
            DeviceRelationship::Connected => ui::DeviceRelationship::Connected,
            DeviceRelationship::Found => ui::DeviceRelationship::Found,
        }
    }
}
unsafe impl Sync for DeviceRelationship {}
unsafe impl Send for DeviceRelationship {}

#[derive(Debug, Clone)]
struct Device {
    name: String,
    color: nlem::storage::Color,
    address: String,
    port: i32,
    id: nlem::security::DeviceID,
    relationship: DeviceRelationship,
}

unsafe impl Sync for Device {}

unsafe impl Send for Device {}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let platform = platform::Platform::init();
    let node = Arc::new(RwLock::new(
        nlem::node::Node::init(Arc::new(RwLock::new(platform)))
            .await
            .expect("Could not create node"),
    ));

    let window = ui::MainWindow::new()?;

    let pair_node = node.clone();
    window.on_request_pair(move |device| {
        let node = pair_node.clone();
        let st: String = device.id.clone().into();
        let device_id = nlem::security::array_from_vec(
            (0..st.len())
                .step_by(2)
                .map(|i| u8::from_str_radix(&st[i..i + 2], 16).unwrap_or(0))
                .collect(),
        );
        tokio::spawn(async move {
            match node.read().await.request_pair(&device_id).await {
                Ok(response) => match response {
                    Ok(device) => {
                        println!("Device paired: {device:#?}");
                    }
                    Err(msg) => {
                        println!("Pairing failed, reason: {msg}");
                    }
                },
                Err(e) => {
                    println!("Error pairing: {e}");
                }
            };
        });
    });

    tokio::spawn(worker(node, window.as_weak()));

    window.run()?;
    Ok(())
}

async fn worker<
    S: nlem::storage::Storage + 'static,
    P: nlem::platform::Platform<S, D> + 'static,
    D: nlem::devicemanager::discovered::DiscoveryManager + 'static,
>(
    node: Arc<RwLock<nlem::node::Node<S, P, D>>>,
    window: slint::Weak<ui::MainWindow>,
) {
    let manager = node.read().await.device_manager.clone();
    let mut manager_event = node.write().await.start().await;

    while let Some(event) = manager_event.recv().await {
        tokio::time::sleep(std::time::Duration::from_secs(1)).await;
        println!("EVENT: {:?}", event);
        match event {
            nlem::devicemanager::DeviceManagerEvent::Started => {
                window
                    .upgrade_in_event_loop(move |window| {
                        let mut status = window.get_node_status();
                        status.okay = true;
                        status.running = true;

                        window.set_node_status(status);
                    })
                    .expect("Could not run callback in event loop");
            }
            nlem::devicemanager::DeviceManagerEvent::InfoLog(log) => println!("[info] {log}"),
            nlem::devicemanager::DeviceManagerEvent::WarningLog(log) => {
                println!("[warn] {log}")
            }

            _ => {}
        };

        let mut devices = std::collections::HashMap::new();
        let manager = manager.read().await;

        for device in manager.paired_devices.read().await.values() {
            devices.insert(
                device.info.id,
                Device {
                    name: device.info.name.clone(),
                    color: device.info.color.clone(),
                    address: "".into(),
                    port: 0,
                    id: device.info.id,
                    relationship: DeviceRelationship::Paired,
                },
            );
        }
        for device in manager.connected_devices.read().await.values() {
            devices.insert(
                device.device.info.id,
                Device {
                    name: device.device.info.name.clone(),
                    color: device.device.info.color.clone(),
                    address: device.client.server_addr.ip().to_string(),
                    port: device.client.server_addr.port().into(),
                    id: device.device.info.id,
                    relationship: DeviceRelationship::Connected,
                },
            );
        }
        for device in manager.discovered_devices.read().await.values() {
            devices.entry(device.info.id).or_insert(Device {
                name: device.info.name.clone(),
                color: device.info.color.clone(),
                address: device.address.ip().to_string(),
                port: device.address.port().into(),
                id: device.info.id,
                relationship: DeviceRelationship::Found,
            });
        }

        window
            .upgrade_in_event_loop(move |window: ui::MainWindow| {
                window.set_devices(slint::ModelRc::new(slint::VecModel::from(
                    devices
                        .into_values()
                        .map(|d| ui::Device {
                            address: d.address.into(),
                            color: slint_col(&d.color),
                            id: d
                                .id
                                .iter()
                                .map(|v| format!("{v:02X}"))
                                .collect::<String>()
                                .into(),
                            name: d.name.into(),
                            port: d.port,
                            relationship: d.relationship.into(),
                        })
                        .collect::<Vec<_>>(),
                )));
            })
            .expect("Could not run callback in event loop");
    }
}
