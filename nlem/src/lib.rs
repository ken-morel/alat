use std::sync::Arc;

use tokio::sync::RwLock;

pub mod client;
pub mod devicemanager;
pub mod discovery;
pub mod node;
pub mod platform;
pub mod proto;
pub mod security;
pub mod server;
pub mod storage;

pub type Storage = dyn storage::Storage + Send + Sync + 'static;
pub type StorageC = Arc<RwLock<Storage>>;
pub type Platform = dyn platform::Platform + Send + Sync + 'static;
pub type PlatformC = Arc<RwLock<Platform>>;
pub type Discovery = dyn discovery::DiscoveryManager + Send + Sync + 'static;
pub type DiscoveryC = Arc<RwLock<Discovery>>;
pub type DeviceManager = devicemanager::DeviceManager;
pub type DeviceManagerC = Arc<RwLock<DeviceManager>>;
pub type Server = server::Server;
pub type ServerC = Arc<RwLock<Server>>;

pub const APP_ID: &str = "cm.engon.alat";
