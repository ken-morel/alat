use super::super::{client, storage};

#[derive(Debug, Clone)]
pub struct ConnectedDevice {
    pub client: client::Client,
    pub device: storage::PairedDevice,
}
impl ConnectedDevice {
    pub fn new(client: client::Client, device: storage::PairedDevice) -> Self {
        Self { client, device }
    }
}
