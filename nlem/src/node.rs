use super::{client, storage};

pub struct Node<S: storage::Storage> {
    storage: S,
    client: client::Client,
}
