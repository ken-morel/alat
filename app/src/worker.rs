use super::{ui, utils::*};
use std::sync::Arc;
use tokio::sync::RwLock;

pub async fn worker(node: Arc<RwLock<nlem::node::Node>>, window: slint::Weak<ui::MainWindow>) {
    let manager = node.read().await.device_manager.clone();
    let mut node_events = node
        .write()
        .await
        .start()
        .await
        .expect("Could not receive device manager event channel");

    while let Some(event) = node_events.recv().await {
        tokio::time::sleep(std::time::Duration::from_secs(2)).await;
        println!("EVENT: {:#?}", event);
        match event {
            nlem::node::NodeEvent::NodeStarted => {
                window
                    .upgrade_in_event_loop(move |window| {
                        let mut status = window.get_node_status();
                        status.okay = true;
                        status.running = true;

                        window.set_node_status(status);
                    })
                    .expect("Could not run callback in event loop");
            } // i need two patterns for lsp
            nlem::node::NodeEvent::NodeStopped => {
                window
                    .upgrade_in_event_loop(move |window| {
                        let mut status = window.get_node_status();
                        status.okay = true;
                        status.running = false;

                        window.set_node_status(status);
                    })
                    .expect("Could not run callback in event loop");
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
