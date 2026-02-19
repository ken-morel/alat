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
pub mod service;
pub mod storage;

pub type StorageC = Arc<RwLock<dyn storage::Storage + Send + Sync + 'static>>;

pub type Platform = dyn platform::Platform + Send + Sync + 'static;
pub type PlatformC = Arc<RwLock<Platform>>;

pub type Discovery = dyn discovery::DiscoveryManager + Send + Sync + 'static;
pub type DiscoveryC = Arc<RwLock<Discovery>>;

pub type DeviceManager = devicemanager::DeviceManager;
pub type DeviceManagerC = Arc<RwLock<DeviceManager>>;

pub type ServiceManager = service::ServiceManager;
pub type ServiceManagerC = Arc<RwLock<ServiceManager>>;

pub type Server = server::Server;
pub type ServerC = Arc<RwLock<Server>>;

pub type Service = dyn service::Service + Send + Sync + 'static;
pub type ServiceC = Arc<RwLock<Service>>;

pub type Node = node::Node;

pub type ErrorC = Box<dyn std::error::Error + Send + Sync>;

pub const APP_ID: &str = "cm.engon.alat";
