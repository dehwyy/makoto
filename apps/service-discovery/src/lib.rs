mod data;

use wasm_bindgen::prelude::*;
use web_sys::{Document,RequestMode};
use data::data::DashboardDataJson;

#[wasm_bindgen]
extern "C" {
    #[wasm_bindgen(js_namespace = console)]
    fn log(s: &str);
}

pub struct Dashboard;

impl Dashboard {
    pub fn update(document: Document, file_s: String) -> makoto::Result<()> {
        const DOCUMENT_TABLE_ID: &str = "#table_data";

        let node = document.query_selector(DOCUMENT_TABLE_ID).unwrap().expect("cannot get node with id `table_data`");
        let mut new_nodes = vec!(node.inner_html());

        let data = DashboardDataJson::new(file_s)?;

        data.state.iter().for_each(|service_state| {
            new_nodes.push(Self::create_html_node(service_state.name.clone(), service_state.addr.clone(), service_state.timestamp.clone()))
        });

        let new_html  = new_nodes.join("\n");

        node.set_inner_html(&new_html);

        Ok(())
    }

    fn clear() {
        todo!()
    }

    fn create_html_node(name: String, value: String, timestamp: String) -> String {
        format!("<tr>
          <td>{0}</td>
          <td>{1}</td>
          <td>{2}</td>
        </tr>", name, value, timestamp)
    }
}

#[wasm_bindgen]
pub async fn run() -> Result<(), JsValue> {
    let window = web_sys::window().expect("no global `window` exists");
    let document = window.document().expect("should have a document on window");

    let fetcher = makoto::wasm::Fetcher::new(makoto::wasm::FetchMethod::GET, RequestMode::SameOrigin, "./data.json")?;
    fetcher.add_header("Content-Type", "application/json")?;
    let file_content = fetcher.fetch_json_as_string(window).await?;

    //
    Dashboard::update(document, file_content).expect("cannot init load!");

    Ok(())
}
