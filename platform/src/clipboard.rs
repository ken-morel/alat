use nlem::units::clipboard::ClipboardContent;

cfg_if::cfg_if! {
    if #[cfg(target_os = "windows")] {
        todo!("Windows not supported");
    } else if #[cfg(target_os = "linux")] {
    }
}

#[derive(Clone)]
pub struct Clipboard {
    arboard: nlem::MContainer<arboard::Clipboard>,
}
impl std::fmt::Debug for Clipboard {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.write_str("<platform::clipboard{}>")
    }
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

#[tonic::async_trait]
impl nlem::units::clipboard::Clipboard for Clipboard {
    async fn get_content(&self) -> Result<ClipboardContent, String> {
        if let Ok(text) = self.arboard.lock().await.get_text() {
            Ok(ClipboardContent::Text(text))
        } else {
            Err(format!("Could not get clipboard content"))
        }
    }
    async fn set_content(&self, content: ClipboardContent) -> Result<(), String> {
        match content {
            ClipboardContent::Text(txt) => self.arboard.lock().await.set_text(txt),
            ClipboardContent::Image((width, height), data) => {
                self.arboard.lock().await.set_image(arboard::ImageData {
                    width: width as usize,
                    height: height as usize,
                    bytes: data.into(),
                })
            }
            ClipboardContent::FileList(files) => self.arboard.lock().await.set_text(
                files
                    .iter()
                    .map(|f| f.to_str().unwrap_or_default().to_string())
                    .collect::<Vec<_>>()
                    .join("\n"),
            ),
            ClipboardContent::Empty => self.arboard.lock().await.set_text(""),
        }
        .map_err(|e| format!("Could not set clipboard: {e}"))
    }
}
