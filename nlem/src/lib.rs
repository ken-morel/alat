use std::sync::Arc;
pub mod error;

use tokio::sync::{Mutex, RwLock};

pub mod client;
pub mod devicemanager;
pub mod discovery;
pub mod node;
pub mod platform;
pub mod proto;
pub mod security;
pub mod server;
pub mod storage;
pub mod unit;
pub mod units;

use crate::devicemanager::DeviceManager;
use crate::server::Server;
use crate::unit::UnitManager;

pub type RWContainer<S> = Arc<RwLock<S>>;
pub type MContainer<S> = Arc<Mutex<S>>;

pub type Storage = dyn storage::Storage + Send + Sync + 'static;
pub type StorageC = MContainer<Storage>;

pub type Platform = dyn platform::Platform + Send + Sync + 'static;
pub type PlatformC = RWContainer<Platform>;

pub type Discovery = dyn discovery::DiscoveryManager + Send + Sync + 'static;
pub type DiscoveryC = RWContainer<Discovery>;

pub type DeviceManagerC = Arc<DeviceManager>;

pub type UnitManagerC = Arc<UnitManager>;

pub type ServerC = Arc<Server>;

pub type Unit = dyn unit::Unit + Send + Sync + 'static;
pub type UnitC = Arc<Unit>;

pub type Node = node::Node;

pub type ErrorC = Box<dyn std::error::Error + Send + Sync>;

pub fn contain<T>(val: T) -> std::sync::Arc<tokio::sync::RwLock<T>> {
    std::sync::Arc::new(tokio::sync::RwLock::new(val))
}
pub fn mcontain<T>(val: T) -> std::sync::Arc<tokio::sync::Mutex<T>> {
    std::sync::Arc::new(tokio::sync::Mutex::new(val))
}

pub const APP_ID: &str = "cm.engon.alat";
