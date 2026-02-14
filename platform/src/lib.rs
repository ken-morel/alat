mod discovery;
mod platform;

cfg_if::cfg_if! {
    if #[cfg(target_os = "windows")] {
        todo!("Windows not supported");
    } else if #[cfg(target_os = "linux")] {
    }
}
