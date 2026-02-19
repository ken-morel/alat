use std::time::Duration;

use super::*;
use crate::client;
use tokio::{
    sync::mpsc::{Receiver, Sender, channel},
    time::sleep,
};

pub(super) enum WorkerEvent {
    Wrapper(DeviceManagerEvent),
}

impl DeviceManager {
    async fn discovery_server_worker(
        coordinator: Sender<WorkerEvent>,
        this_device: Arc<RwLock<discovery::DiscoveredDevice>>,
        discovery: crate::DiscoveryC,
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

    async fn discoverer_worker(worker_events: Sender<WorkerEvent>, discovery: crate::DiscoveryC) {
        let (tx, mut rx) = channel(1);
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
                crate::discovery::DiscoveryEvent::Found(device) => {
                    worker_events
                        .send(WorkerEvent::Wrapper(DeviceManagerEvent::Found(Box::new(
                            device,
                        ))))
                        .await
                        .expect("Could not send Found message to worker_events");
                }
                crate::discovery::DiscoveryEvent::Lost(device_id) => {
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
        discovered_devices: Arc<RwLock<HashMap<security::DeviceID, discovery::DiscoveredDevice>>>,
        connected_devices: Arc<RwLock<HashMap<security::DeviceID, connected::ConnectedDevice>>>,
        paired_devices: Arc<RwLock<HashMap<security::DeviceID, storage::PairedDevice>>>,
    ) {
        global_events
            .send(DeviceManagerEvent::Started)
            .await
            .expect("Could not send device manager started event");
        let try_connect = async |device, address| {
            match client::Client::connect(address).await {
                Ok(client) => {
                    let connected_device = connected::ConnectedDevice { client, device };
                    connected_devices
                        .write()
                        .await
                        .entry(connected_device.device.info.id)
                        .or_insert_with(|| connected_device.clone()); // clone only if needed
                    global_events
                        .send(DeviceManagerEvent::Connected(Box::new(connected_device)))
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
        };
        while let Some(event) = worker_events.recv().await {
            match event {
                WorkerEvent::Wrapper(event) => match &event {
                    DeviceManagerEvent::Found(found_device) => {
                        discovered_devices
                            .write()
                            .await
                            .insert(found_device.info.id, *found_device.clone());
                        global_events
                            .send(event.clone())
                            .await
                            .expect("Could not send found device event"); // send the found event before trying to connect

                        // update the paired_devices if found and clone it to local paired_device
                        paired_devices
                            .write()
                            .await
                            .entry(found_device.info.id)
                            .and_modify(|paired| {
                                paired.info = found_device.info.clone();
                            });
                        // attempt to connect
                        if let Some(device) = paired_devices
                            .read()
                            .await
                            .get(&found_device.info.id)
                            .cloned()
                        {
                            try_connect(device, found_device.address).await;
                        }
                    }
                    echoable => {
                        match echoable {
                            DeviceManagerEvent::Lost(device_id) => {
                                connected_devices.write().await.remove(device_id);
                                discovered_devices.write().await.remove(device_id);
                            }
                            DeviceManagerEvent::Paired(paired) => {
                                paired_devices
                                    .write()
                                    .await
                                    .insert(paired.info.id, paired.clone());
                                if let Some(found) =
                                    discovered_devices.read().await.get(&paired.info.id)
                                {
                                    try_connect(paired.clone(), found.address).await;
                                }
                            }
                            DeviceManagerEvent::Unpaired(dev_id) => {
                                paired_devices.write().await.remove(dev_id);
                                connected_devices.write().await.remove(dev_id);
                            }
                            _ => {}
                        };
                        global_events
                            .send(event)
                            .await
                            .expect("COuld not transfer global event");
                    }
                },
            }
        }
        global_events
            .send(DeviceManagerEvent::Stopped)
            .await
            .expect("Could not send device manager started event");
    }
    pub async fn start_workers(&mut self, sender: Sender<DeviceManagerEvent>) {
        println!("Spawning workers");
        let (itx, irx) = channel::<WorkerEvent>(1);

        self.worker = Some(itx.clone());

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
