use super::proto;

#[derive(Debug, serde::Deserialize, serde::Serialize)]
pub struct Color(u8, u8, u8);
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
