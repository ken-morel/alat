use std::time::Duration;

use super::*;
use tokio::{
    sync::mpsc::{Receiver, Sender, channel},
    time::sleep,
};

enum WorkerEvent {
    Wrapper(DeviceManagerEvent),
}

impl<
    S: storage::Storage + 'static,
    P: platform::Platform + 'static,
    D: discovered::DiscoveryManager + 'static,
> DeviceManager<S, P, D>
{
    async fn discovery_server_worker(
        coordinator: Sender<WorkerEvent>,
        this_device: Arc<RwLock<discovered::DiscoveredDevice>>,
        discovery: Arc<RwLock<D>>,
    ) {
        let mut current_info = this_device.read().await.clone();
        coordinator
            .send(WorkerEvent::Wrapper(
                DeviceManagerEvent::DiscoveryServerStarted(current_info.clone()),
            ))
            .await
            .expect("Could not send server starting message");
        discovery
            .write()
            .await
            .advertise(current_info.clone())
            .await
            .expect("Discovery worker failed");
        loop {
            sleep(Duration::from_secs(2)).await;
            let new_info = this_device.read().await.clone();
            if new_info != current_info {
                current_info = new_info;
                discovery
                    .write()
                    .await
                    .cease_advertising()
                    .await
                    .expect("Could not stop advertiser");
                sleep(Duration::from_secs(1)).await;
                discovery
                    .write()
                    .await
                    .advertise(current_info.clone())
                    .await
                    .expect("Could not restart advertiser");
            }
        }
    }

    async fn discoverer_worker(worker_events: Sender<WorkerEvent>) {}

    async fn coordinator_worker(
        worker_events: Receiver<WorkerEvent>,
        global_events: Sender<DeviceManagerEvent>,
    ) {
    }
    pub async fn start_workers(&mut self, sender: Sender<DeviceManagerEvent>) {
        let (itx, irx) = channel::<WorkerEvent>(0);

        tokio::spawn(Self::discovery_server_worker(
            itx.clone(),
            self.this_device.clone(),
            self.discovery.clone(),
        ));
        tokio::spawn(Self::discoverer_worker(itx.clone()));
        tokio::spawn(Self::coordinator_worker(irx, sender));
    }
}
