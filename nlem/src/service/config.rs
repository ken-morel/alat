use crate::storage;

#[derive(Clone)]
pub struct ServiceConfig {
    storage: crate::StorageC,
    service_id: super::ServiceID,
    settings_key: String,
}

unsafe impl Send for ServiceConfig {}

unsafe impl Sync for ServiceConfig {}

impl ServiceConfig {
    pub fn new(storage: crate::StorageC, service_id: super::ServiceID) -> Self {
        let mut settings_key = String::from("services/");
        settings_key.push_str(service_id);
        Self {
            storage,
            service_id,
            settings_key,
        }
    }
    pub async fn load<T: serde::de::DeserializeOwned>(&self) -> storage::StorageResult<Option<T>> {
        if let Some(json_value) = self
            .storage
            .lock()
            .await
            .load_settings(self.service_id)
            .await?
        {
            self.deserialize(json_value)
        } else {
            Ok(None)
        }
    }
    pub fn serialize<T: serde::Serialize>(
        &self,
        data: T,
    ) -> storage::StorageResult<serde_json::Value> {
        Ok(serde_json::to_value(data)
            .map_err(|e| storage::StorageError::Serialize(e.to_string()))?)
    }
    pub fn deserialize<T: for<'a> serde::Deserialize<'a>>(
        &self,
        data: serde_json::Value,
    ) -> storage::StorageResult<T> {
        Ok(serde_json::from_value(data)
            .map_err(|e| storage::StorageError::Deserialize(e.to_string()))?)
    }
    pub async fn save<T: serde::Serialize>(&self, data: T) -> storage::StorageResult<()> {
        Ok(self
            .storage
            .lock()
            .await
            .save_settings(&self.settings_key, &self.serialize(data)?)
            .await?)
    }
    pub async fn init<T: serde::Serialize + serde::de::DeserializeOwned + Clone>(
        &mut self,
        data: T,
    ) -> storage::StorageResult<T> {
        if let Ok(Some(data)) = self.load().await {
            Ok(data)
        } else {
            self.save(data.clone()).await;
            Ok(data)
        }
    }
}
