use std::sync::{Arc, Mutex};

use super::{client, devicemanager, platform, server, storage};

pub struct Node<S: storage::Storage, P: platform::Platform> {
    storage: Arc<Mutex<S>>,
    platform: Arc<Mutex<P>>,
    device_manager: Arc<Mutex<devicemanager::DeviceManager<S, P>>>,
    server: Arc<Mutex<server::Server>>,
}
