use wasm_bindgen::prelude::*;
use wasm_bindgen_futures::{wasm_bindgen::{JsValue, JsCast}, JsFuture, js_sys::JSON};
use web_sys::{RequestInit, Request, Window, Response};

#[wasm_bindgen]
pub enum FetchMethod {
  GET,
  POST,
  DELETE
}

#[wasm_bindgen]
pub struct Fetcher {
  request: Request
}

#[wasm_bindgen]
impl Fetcher {
  #[wasm_bindgen(constructor)]
  pub fn new(method: FetchMethod, mode: web_sys::RequestMode, url: &str) -> Result<Fetcher, JsValue> {
    let method = match method {
      FetchMethod::GET => "GET",
      FetchMethod::POST => "POST",
      FetchMethod::DELETE => "DELETE"
    };

    let mut options = RequestInit::new();
    options.method(method);
    options.mode(mode);

    let request = Request::new_with_str_and_init(url, &options)?;


    Ok(Fetcher {
      request,
    })
  }

  pub fn add_header(&self, key: &str, value: &str) -> Result<(), JsValue> {
    self.request.headers().set(key, value)?;

    Ok(())
  }

  pub async fn fetch_json_as_string(&self, window: Window) -> Result<String, JsValue> {
    let response = JsFuture::from(window.fetch_with_request(&self.request)).await?;
    let response: Response = response.dyn_into()?;
    let response_json = JsFuture::from(response.json()?).await?;

    Ok(JSON::stringify(&response_json)?.as_string().expect("cannot convert into string!"))
  }
}
