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
    id: Vec<i32>,
    relationship: DeviceRelationship,
}

unsafe impl Sync for Device {}

unsafe impl Send for Device {}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let platform = platform::Platform::init();
    let mut node = nlem::node::Node::init(Arc::new(RwLock::new(platform)))
        .await
        .expect("Could not create node");

    let window = ui::MainWindow::new()?;
    window.set_node_status(ui::NodeStatus {
        okay: true,
        running: false,
    });
    let handle = window.as_weak();

    tokio::spawn(async move {
        let mut manager_event = node.start().await;
        let node = Arc::new(RwLock::new(node));
        while let Some(event) = manager_event.recv().await {
            println!("EVENT: {:?}", event);
            match event {
                nlem::devicemanager::DeviceManagerEvent::Started => {
                    handle
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
            }

            let node = node.read().await;

            let mut devices = std::collections::HashMap::new();
            let manager = node.device_manager.read().await;

            for device in manager.paired_devices.read().await.values() {
                devices.insert(
                    device.info.id,
                    Device {
                        name: device.info.name.clone(),
                        color: device.info.color.clone(),
                        address: "".into(),
                        port: 0,
                        id: device.info.id.map(|v| v as i32).into(),
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
                        id: device.device.info.id.map(|v| v as i32).into(),
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
                    id: device.info.id.map(|v| v as i32).into(),
                    relationship: DeviceRelationship::Found,
                });
            }

            handle
                .upgrade_in_event_loop(move |window: ui::MainWindow| {
                    window.set_devices(slint::ModelRc::new(slint::VecModel::from(
                        devices
                            .into_values()
                            .map(|d| ui::Device {
                                address: d.address.into(),
                                color: slint_col(&d.color),
                                id: slint::ModelRc::new(slint::VecModel::from(d.id)),
                                name: d.name.into(),
                                port: d.port,
                                relationship: d.relationship.into(),
                            })
                            .collect::<Vec<_>>(),
                    )));
                })
                .expect("Could not run callback in event loop");
        }
    });

    window.run()?;

    Ok(())
}
