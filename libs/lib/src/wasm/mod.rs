mod fetcher;
mod actions;

pub use fetcher::{Fetcher, FetchMethod};
pub use actions::EventListener;
use wasm_bindgen::JsValue;

pub fn convert_js_value<T>(v: Result<T, JsValue>) -> Result<T, std::io::Error> {
  match v {
    Ok(v) => Ok(v),
    Err(_) => Err(std::io::Error::new(std::io::ErrorKind::Other, "js error"))
  }
}
