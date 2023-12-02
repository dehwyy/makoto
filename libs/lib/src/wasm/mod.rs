mod utils;

pub use utils::{Fetcher, FetchMethod};
use wasm_bindgen::JsValue;

pub fn convert_js_value<T>(v: Result<T, JsValue>) -> Result<T, std::io::Error> {
  match v {
    Ok(v) => Ok(v),
    Err(_) => Err(std::io::Error::other("js error"))
  }
}
