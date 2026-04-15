use crate::proto;
use std::path::PathBuf;
#[derive(Debug, Clone, Default, PartialEq)]
pub enum ClipboardContent {
    Text(String),
    Image((u32, u32), Vec<u8>),
    FileList(Vec<PathBuf>),
    #[default]
    Empty,
}
impl From<ClipboardContent> for proto::ClipboardContent {
    fn from(c: ClipboardContent) -> Self {
        Self {
            data: Some(match c {
                ClipboardContent::Text(txt) => {
                    proto::clipboard_content::Data::Text(proto::ClipboardContentText { data: txt })
                }
                ClipboardContent::Image((width, height), data) => {
                    proto::clipboard_content::Data::Image(proto::ClipboardContentImage {
                        width,
                        height,
                        data,
                    })
                }
                ClipboardContent::FileList(files) => {
                    proto::clipboard_content::Data::Files(proto::ClipboardContentFiles {
                        files: files
                            .into_iter()
                            .map(|b| b.to_str().unwrap_or_default().to_string())
                            .collect(),
                    })
                }
                ClipboardContent::Empty => {
                    proto::clipboard_content::Data::Empty(proto::ClipboardContentEmpty {})
                }
            }),
        }
    }
}
