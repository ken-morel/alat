use nlem::service::clipboard::ClipboardContent;

cfg_if::cfg_if! {
    if #[cfg(target_os = "windows")] {
        todo!("Windows not supported");
    } else if #[cfg(target_os = "linux")] {
    }
}

pub struct Clipboard {
    arboard: nlem::MContainer<arboard::Clipboard>,
}

impl Clipboard {
    pub fn init() -> Result<Self, String> {
        Ok(Self {
            arboard: nlem::mcontain(
                arboard::Clipboard::new()
                    .map_err(|e| format!("Platform could not get clipboard: {e}"))?,
            ),
        })
    }
}

impl nlem::service::clipboard::Clipboard for Clipboard {
    async fn watch(
        &self,
    ) -> Result<tokio::sync::mpsc::Receiver<nlem::service::clipboard::ClipboardContent>, String>
    {
        let clipboard = self.clipboard.clone();
        let (tx, rx) = tokio::sync::mpsc::channel(1);
        let mut previous = self.get_content().await.unwrap_or_default();
        tokio::spawn(async move {
            loop {
                tokio::time::sleep(std::time::Duration::from_secs(5)).await;
                if let Ok(current) = self.get_content() {
                    if current != previous {
                        previous = current.clone();
                        tx.send(current)
                            .await
                            .expect("Could not send clipboard content to receiving channel");
                    }
                }
            }
        });
        Ok(rx)
    }
    async fn get_content(&self) -> Result<ClipboardContent, String> {
        if let Ok(text) = self.arboard.lock().await.get_text() {
            ClipboardContent::Text(text)
        }
    }
    async fn set_content(&self, content: ClipboardContent) -> Result<(), String> {
        match content {
            ClipboardContent::Text(txt) => self.arboard.lock().await.set_text(text),
            ClipboardContent::Image((width, height), data) => {
                self.arboard.lock().await.set_image(arboard::ImageData {
                    width: width,
                    height: height,
                    bytes: data,
                })
            }
            ClipboardContent::FileList(files) => self
                .arboard
                .lock()
                .await
                .set_text(files.iter().map(|f| f.to_string_lossy()).collect().join()),
        }
        .map_err(|e| format!("Could not set clipboard: {e}"))
    }
}
