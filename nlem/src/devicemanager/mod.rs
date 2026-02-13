mod connected;

use crate::storage::StorageError;

use super::{platform, security, storage};
use std::sync::{Arc, Mutex};

#[derive(Debug)]
pub struct DeviceManager<S: storage::Storage, P: platform::Platform> {
    pub storage: Arc<Mutex<S>>,
    pub platform: Arc<Mutex<P>>,

    pub paired_devices: Vec<storage::PairedDevice>,
    pub device_info: storage::DeviceInfo,
    pub device_certificate: security::Certificate,

    pub connected_devices: Vec<connected::ConnectedDevice>,
}

impl<S: storage::Storage, P: platform::Platform> DeviceManager<S, P> {
    pub fn new(store: Arc<Mutex<S>>, platform: Arc<Mutex<P>>) -> Self {
        Self {
            device_info: DeviceManager::<S, P>::default_device_info(&platform),
            platform,
            storage: store,
            paired_devices: Vec::new(),
            device_certificate: security::generate_certificate(),
            connected_devices: Vec::new(),
        }
    }
    pub fn load(&mut self) -> Result<(), StorageError> {
        let store = self.storage.lock().expect("Could not lock node storage");
        self.paired_devices = store.load_paired()?;
        self.device_info = store.load_info()?;
        self.device_certificate = store.load_certificate()?;
        Ok(())
    }
    pub fn save(&self) -> Result<(), StorageError> {
        let store = self.storage.lock().expect("Could not lock storage");
        store.save_paired(self.paired_devices.clone())?;
        store.save_info(self.device_info.clone())?;
        store.save_certificate(self.device_certificate.clone())?;
        Ok(())
    }
    pub fn default_device_info(p: &Arc<Mutex<P>>) -> storage::DeviceInfo {
        let lck = p.lock().expect("Could not lock platform");
        storage::DeviceInfo {
            id: security::generate_id(),
            color: storage::Color::random(),
            name: lck.hostname().expect("Could not get hostname"),
            device_type: lck.device_type(),
        }
    }
}
