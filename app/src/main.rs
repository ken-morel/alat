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

    tokio::spawn(async move {
        println!("Starting node");
        let mut manager_event = node.start().await;
        while let Some(event) = manager_event.recv().await {
            println!("EVENT: {:?}", event);
        }
    });

    let window = ui::MainWindow::new()?;

    window.run()?;

    Ok(())
}
