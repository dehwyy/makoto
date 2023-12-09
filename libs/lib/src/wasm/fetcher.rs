use wasm_bindgen_futures::{wasm_bindgen::{JsValue, JsCast}, JsFuture, js_sys::JSON};
use web_sys::{RequestInit, Request, Response, window};

pub enum FetchMethod {
  GET,
  POST,
  DELETE
}

pub struct Fetcher {
  request: Request
}

impl Fetcher {
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

  /// Doesn't return response
  pub async fn fetch(&self) -> Result<(), JsValue> {
    let w = window().unwrap();
    JsFuture::from(w.fetch_with_request(&self.request)).await?;
    Ok(())
  }

  pub async fn fetch_json_as_string(&self) -> Result<String, JsValue> {
    let w = window().unwrap();
    let response = JsFuture::from(w.fetch_with_request(&self.request)).await?;
    let response: Response = response.dyn_into()?;
    let response_json = JsFuture::from(response.json()?).await?;

    Ok(JSON::stringify(&response_json)?.as_string().expect("cannot convert into string!"))
  }
}
