use std::sync::Arc;

use tokio::sync::{Mutex, RwLock};

pub mod client;
pub mod devicemanager;
pub mod discovery;
pub mod node;
pub mod platform;
pub mod proto;
pub mod security;
pub mod server;
pub mod service;
pub mod storage;

pub type RWContainer<S> = Arc<RwLock<S>>;
pub type MContainer<S> = Arc<Mutex<S>>;

pub type Storage = dyn storage::Storage + Send + Sync + 'static;
pub type StorageC = MContainer<Storage>;

pub type Platform = dyn platform::Platform + Send + Sync + 'static;
pub type PlatformC = RWContainer<Platform>;

pub type Discovery = dyn discovery::DiscoveryManager + Send + Sync + 'static;
pub type DiscoveryC = RWContainer<Discovery>;

pub type DeviceManager = devicemanager::DeviceManager;
pub type DeviceManagerC = RWContainer<DeviceManager>;

pub type ServiceManager = service::ServiceManager;
pub type ServiceManagerC = RWContainer<ServiceManager>;

pub type Server = server::Server;
pub type ServerC = RWContainer<Server>;

pub type Service = dyn service::Service + Send + Sync + 'static;
pub type ServiceC = RWContainer<Service>;

pub type Node = node::Node;

pub type ErrorC = Box<dyn std::error::Error + Send + Sync>;

pub fn contain<T>(val: T) -> std::sync::Arc<tokio::sync::RwLock<T>> {
    std::sync::Arc::new(tokio::sync::RwLock::new(val))
}
pub fn mcontain<T>(val: T) -> std::sync::Arc<tokio::sync::Mutex<T>> {
    std::sync::Arc::new(tokio::sync::Mutex::new(val))
}

pub const APP_ID: &str = "cm.engon.alat";
