mod ui;
mod utils;
mod worker;

use std::sync::Arc;
use tokio::sync::RwLock;

use slint::ComponentHandle;

use worker::worker;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let platform = platform::Platform::init();
    let node = Arc::new(RwLock::new(
        nlem::node::Node::init(Arc::new(RwLock::new(platform)))
            .await
            .expect("Could not create node"),
    ));

    let window = ui::MainWindow::new()?;
    let weak_window = window.as_weak();

    let pair_node = node.clone();
    window.on_request_pair(move |device| {
        let weak_window = weak_window.clone();
        let node = pair_node.clone();
        let st: String = device.id.clone().into();
        let device_id = nlem::security::array_from_vec(
            (0..st.len())
                .step_by(2)
                .map(|i| u8::from_str_radix(&st[i..i + 2], 16).unwrap_or(0))
                .collect(),
        );
        weak_window
            .upgrade_in_event_loop(|window| {
                window.set_is_pairing(true);
                window.set_pairing_succeded(false);
                window.set_pairing_error("Unknown error".into());
            })
            .expect("Could not upgrade window");

        tokio::spawn(async move {
            let weak_window = weak_window.clone();
            let response = node.read().await.request_pair(&device_id).await;
            weak_window
                .upgrade_in_event_loop(move |window| {
                    window.set_is_pairing(false);
                    match response {
                        Ok(response) => match response {
                            Ok(device) => {
                                println!("Device paired: {device:#?}");
                                window.set_pairing_succeded(true);
                            }
                            Err(msg) => {
                                println!("Pairing failed, reason: {msg}");
                                window.set_pairing_succeded(false);
                                window.set_pairing_error(msg.into());
                            }
                        },
                        Err(e) => {
                            println!("Error pairing: {e}");
                            window.set_pairing_succeded(false);
                            window.set_pairing_error(e.into());
                        }
                    };
                })
                .expect("Could not upgrade window");
        });
    });
    window.set_devices(slint::ModelRc::new(slint::VecModel::from(Vec::new())));

    tokio::spawn(worker(node, window.as_weak()));

    window.run()?;
    Ok(())
}
