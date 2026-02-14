use std::time::Duration;

use super::*;
use crate::client;
use tokio::{
    sync::mpsc::{Receiver, Sender, channel},
    time::sleep,
};

enum WorkerEvent {
    Wrapper(DeviceManagerEvent),
}

impl<
    S: storage::Storage + 'static,
    P: platform::Platform<D> + 'static,
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
                DeviceManagerEvent::DiscoveryServerStarted(Box::new(current_info.clone())),
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

    async fn discoverer_worker(worker_events: Sender<WorkerEvent>, discovery: Arc<RwLock<D>>) {
        let (tx, mut rx) = channel(0);
        if let Err(e) = discovery.write().await.scan(tx).await {
            worker_events
                .send(WorkerEvent::Wrapper(DeviceManagerEvent::DiscoveryError(e)))
                .await
                .expect("Could not sent error to cworker channel");
            return;
        } else {
            worker_events
                .send(WorkerEvent::Wrapper(DeviceManagerEvent::DiscoveryStarted))
                .await
                .expect("Could not send start to worker");
        }
        while let Some(event) = rx.recv().await {
            match event {
                discovered::DiscoveryEvent::Found(device) => {
                    worker_events
                        .send(WorkerEvent::Wrapper(DeviceManagerEvent::Found(Box::new(
                            device,
                        ))))
                        .await
                        .expect("Could not send Found message to worker_events");
                }
                discovered::DiscoveryEvent::Lost(device_id) => {
                    worker_events
                        .send(WorkerEvent::Wrapper(DeviceManagerEvent::Lost(device_id)))
                        .await
                        .expect("Could not send Lost message to worker_events");
                }
            }
        }
        discovery
            .write()
            .await
            .cease_scan()
            .await
            .expect("Could not cease scan");
        worker_events
            .send(WorkerEvent::Wrapper(DeviceManagerEvent::DiscoveryStopped))
            .await
            .expect("Could not send stopped message to worker_events");
    }

    async fn coordinator_worker(
        mut worker_events: Receiver<WorkerEvent>,
        global_events: Sender<DeviceManagerEvent>,
        discovered_devices: Arc<RwLock<HashMap<security::DeviceID, discovered::DiscoveredDevice>>>,
        connected_devices: Arc<RwLock<HashMap<security::DeviceID, connected::ConnectedDevice>>>,
        paired_devices: Arc<RwLock<HashMap<security::DeviceID, storage::PairedDevice>>>,
    ) {
        'event_loop: while let Some(event) = worker_events.recv().await {
            match event {
                WorkerEvent::Wrapper(event) => match &event {
                    DeviceManagerEvent::Found(found_device) => {
                        // let's do this, we pass the data to send in events, and we clone the data
                        // to the designated vec or hashmap.
                        let mut already_found = true;
                        discovered_devices
                            .write()
                            .await
                            .entry(found_device.info.id.clone())
                            .or_insert_with(|| {
                                already_found = false;
                                *found_device.clone()
                            });
                        if already_found {
                            continue 'event_loop; // just to be clear
                        }
                        global_events
                            .send(event.clone())
                            .await
                            .expect("Could not send found device event"); // send the found event before trying to connect

                        let mut paired_device = Option::None;

                        // update the paired_devices if found and clone it to local paired_device
                        paired_devices
                            .write()
                            .await
                            .entry(found_device.info.id.clone())
                            .and_modify(|paired| {
                                paired.info = found_device.info.clone();
                                paired_device = Some(paired.clone());
                            });
                        // attempt to connect
                        if let Some(device) = paired_device {
                            match client::Client::connect(found_device.address).await {
                                Ok(client) => {
                                    let connected_device =
                                        connected::ConnectedDevice { client, device };

                                    connected_devices
                                        .write()
                                        .await
                                        .entry(connected_device.device.info.id.clone())
                                        .or_insert_with(|| connected_device.clone()); // clone only if needed
                                    global_events
                                        .send(DeviceManagerEvent::Connected(Box::new(
                                            connected_device,
                                        )))
                                        .await
                                        .expect("Could not send connected event to global_events");
                                }
                                Err(err) => {
                                    global_events
                                        .send(DeviceManagerEvent::ConnectionError(err.to_string()))
                                        .await
                                        .expect("Could not send connection error to global events");
                                }
                            }
                        }
                    }
                    DeviceManagerEvent::Lost(device_id) => {
                        connected_devices.write().await.remove(device_id);
                        discovered_devices.write().await.remove(device_id);
                        global_events
                            .send(event)
                            .await
                            .expect("Could not send lost event to global_events");
                    }
                    _ => global_events
                        .send(event)
                        .await
                        .expect("Could not send global_event"),
                },
            }
        }
    }
    pub async fn start_workers(&mut self, sender: Sender<DeviceManagerEvent>) {
        let (itx, irx) = channel::<WorkerEvent>(0);

        tokio::spawn(Self::discovery_server_worker(
            itx.clone(),
            self.this_device.clone(),
            self.discovery.clone(),
        ));
        tokio::spawn(Self::discoverer_worker(itx.clone(), self.discovery.clone()));
        tokio::spawn(Self::coordinator_worker(
            irx,
            sender,
            self.discovered_devices.clone(),
            self.connected_devices.clone(),
            self.paired_devices.clone(),
        ));
    }
}
