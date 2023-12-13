use makoto::wasm::{Fetcher, FetchMethod};
use wasm_bindgen::JsValue;
use web_sys::RequestMode;

const API_URL: &str = "http://localhost:4223/api";

struct ApiUrl;

impl ApiUrl {
  /// f.e. `data/boba`
  fn get(endpoint: &str) -> String {
    format!("{0}/{1}", API_URL, endpoint)
  }
}

pub struct ApiFetch;

impl ApiFetch {
  pub async fn get_data() -> Result<String, JsValue> {
      let fetcher = Fetcher::new(FetchMethod::GET, RequestMode::Cors, &ApiUrl::get("data"))?;

      // attach headers
      fetcher.add_header("Content-Type", "application/json")?;
      fetcher.add_header("Cache-Control", "no-store")?;

      Ok(fetcher.fetch_json_as_string().await?)
  }

  pub async fn clear_data() -> Result<(), JsValue> {
    let fetcher = Fetcher::new(FetchMethod::DELETE, RequestMode::Cors, &ApiUrl::get("data"))?;

    fetcher.fetch().await?;

    Ok(())
  }
}
