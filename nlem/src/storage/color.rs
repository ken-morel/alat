use rand::random_range;

use super::proto;

#[derive(Debug, Clone, PartialEq, Eq, serde::Deserialize, serde::Serialize)]
pub struct Color(pub u8, pub u8, pub u8);
impl From<proto::Color> for Color {
    fn from(col: proto::Color) -> Self {
        Self(
            col.r.try_into().unwrap_or(0),
            col.g.try_into().unwrap_or(0),
            col.b.try_into().unwrap_or(0),
        )
    }
}
impl From<Color> for proto::Color {
    fn from(col: Color) -> Self {
        Self {
            r: col.0.into(),
            g: col.1.into(),
            b: col.2.into(),
        }
    }
}
impl Color {
    pub fn random() -> Self {
        Self(
            random_range(0..=2) * 127,
            random_range(0..=2) * 127,
            random_range(0..=2) * 127,
        )
    }
}
