mod discovery;
mod platform;
mod storage;

pub type Platform = platform::Platform;

cfg_if::cfg_if! {
    if #[cfg(target_os = "windows")] {
        todo!("Windows not supported");
    } else if #[cfg(target_os = "linux")] {
    }
}
