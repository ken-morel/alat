use super::{ui, utils::*};
use std::sync::Arc;

pub async fn worker(
    node: nlem::node::Node,
    window: slint::Weak<ui::MainWindow>,
) {
    let mut node_events = node
        .start()
        .await
        .expect("Could not start node and get event manager channel");

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

        {
            let manager = node.device_manager.clone();

            for device in manager.paired_devices.iter() {
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
            for device in manager.connected_devices.iter() {
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
            for device in manager.discovered_devices.iter() {
                devices.entry(device.info.id).or_insert(Device {
                    name: device.info.name.clone(),
                    color: device.info.color.clone(),
                    address: device.address.ip().to_string(),
                    port: device.address.port().into(),
                    id: device.info.id,
                    relationship: DeviceRelationship::Found,
                });
            }
            drop(manager);
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
