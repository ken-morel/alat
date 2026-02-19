pub const CERTIFICATE_SIZE: usize = 8; // has no use for now
pub type Certificate = Vec<u8>;
pub type DeviceID = [u8; 16];
pub type PairToken = [u8; 32];

pub fn array_from_vec<const N: usize>(mut v: Vec<u8>) -> [u8; N] {
    v.resize(N, 0u8);
    v.try_into().unwrap()
}

pub fn generate_id() -> DeviceID {
    let mut devid: DeviceID = [0u8; _];
    rand::fill(&mut devid);
    devid
}
pub fn generate_certificate() -> Certificate {
    let mut cert = vec![0u8; CERTIFICATE_SIZE];
    rand::fill(&mut cert);
    cert
}

pub fn generate_pair_token() -> PairToken {
    let mut token: PairToken = [0u8; _];
    rand::fill(&mut token);
    token
}
