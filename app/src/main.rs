mod ui;
use std::sync::Arc;
use tokio::sync::RwLock;

use slint::ComponentHandle;

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
        while let Some(event) = manager_event.recv().await {
            tokio::time::sleep(tokio::time::Duration::from_secs(1)).await;
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
        }
    });

    window.run()?;

    Ok(())
}
