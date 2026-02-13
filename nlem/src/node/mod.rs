use std::sync::{Arc, Mutex};

use super::{devicemanager, platform, server, storage};

pub struct Node<
    S: storage::Storage,
    P: platform::Platform,
    D: devicemanager::discovered::DiscoveryManager,
> {
    storage: Arc<Mutex<S>>,
    platform: Arc<Mutex<P>>,
    device_manager: Arc<Mutex<devicemanager::DeviceManager<S, P, D>>>,
    server: Arc<Mutex<server::Server>>,
}
