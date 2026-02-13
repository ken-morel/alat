pub const DEVICE_ID_LENGTH: usize = 1 << 10; // 1kB
pub const DEVICE_CERTIFICATE_LENGTH: usize = 1 << 20; // 1MB

pub type DeviceID = Vec<u8>;
pub type Certificate = Vec<u8>;

fn generate_vecu8(size: usize) -> Vec<u8> {
    let mut vec = vec![0u8; size];
    rand::fill(&mut vec);
    vec
}

pub fn generate_id() -> DeviceID {
    generate_vecu8(DEVICE_ID_LENGTH)
}
pub fn generate_certificate() -> Certificate {
    generate_vecu8(DEVICE_CERTIFICATE_LENGTH)
}
