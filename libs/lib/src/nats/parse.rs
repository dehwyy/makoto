use bytes::Bytes;

use crate::Result as MakotoResult;

pub struct MessageParser;

impl MessageParser {
  pub fn plain(b: &Bytes) -> MakotoResult<String> {
    Ok(String::from_utf8(b.to_vec())?)
  }
  pub fn key_value(b: &Bytes) -> MakotoResult<(String, String)> {
    let message_payload = String::from_utf8(b.to_vec())?;

    match message_payload.split_once(";") {
      Some(v) => Ok((v.0.to_string(), v.1.to_string())),
      None => Err(Box::new(std::io::Error::other("cannot parse message_payload")))
    }
  }
}
