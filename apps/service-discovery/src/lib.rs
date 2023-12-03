mod data;
mod node;

use wasm_bindgen::prelude::*;
use wasm_bindgen_futures::spawn_local;
use web_sys::{RequestMode, MouseEvent, window};
use data::data::DashboardDataJson;
use node::fabric::{Fabric as NodeFabric, CreatePayload as NodeCreatePayload};
use makoto::wasm::{Fetcher, convert_js_value, FetchMethod};

#[wasm_bindgen]
extern "C" {
    #[wasm_bindgen(js_namespace = console)]
    fn log(s: &str);
}

#[derive(Clone, Copy, Debug)]
enum DashboardView {
    State,
    Events
}

#[derive(Default)]
struct Dashboard;

impl Dashboard {
    async fn new() -> makoto::Result<Dashboard> {
        // init dashboard
        let dashboard = Self {};
        Self::set_view(DashboardView::State).await;

        dashboard.setup_listeners();
        Ok(dashboard)

    }

    // setup listener (once)
    fn setup_listeners(&self) {
        let document = window().unwrap().document().unwrap();

        // SetState button
        {
            let set_state = document.query_selector("#set_state").unwrap().expect("cannot get #set_state");

            let onclick = Closure::<dyn FnMut(_)>::new(move |_: MouseEvent | {
                spawn_local(Self::set_view(DashboardView::State));

            });

            set_state.add_event_listener_with_callback("click", onclick.as_ref().unchecked_ref()).unwrap();
            onclick.forget();
        }

        // SetEvents button
        {
            let set_events = document.query_selector("#set_events").unwrap().expect("cannot get #set_events");

            let onclick = Closure::<dyn FnMut(_)>::new(|_: MouseEvent | {
                spawn_local(Self::set_view(DashboardView::Events));
            });

            set_events.add_event_listener_with_callback("click", onclick.as_ref().unchecked_ref()).unwrap();
            onclick.forget();
        }
    }

    /// fetches file content as json using window.fetch
    async fn get_file_content() -> Result<String, JsValue> {
        let fetcher = Fetcher::new(FetchMethod::GET, RequestMode::SameOrigin, "./data.json")?;

        // attach headers
        fetcher.add_header("Content-Type", "application/json")?;
        fetcher.add_header("Cache-Control", "no-store")?;

        Ok(fetcher.fetch_json_as_string(window().unwrap()).await?)
    }

    /// return new html by specified `view`
    async fn get_updated_html_by_view(view: DashboardView) -> makoto::Result<String> {
        let file_as_string = convert_js_value(Self::get_file_content().await)?;

        let data = DashboardDataJson::new(file_as_string)?;

        match view {
            DashboardView::State => {
                let mut nodes: Vec<String> = vec!();
                data.state.iter().for_each(|service_state| {
                    nodes.push(NodeFabric::new_state_table_node(&service_state.name, &service_state.addr, &service_state.timestamp))
                });

                Ok(nodes.join("\n"))
            },
            DashboardView::Events => {
                let mut nodes: Vec<String> = vec!();
                data.events.iter().for_each(|event| {
                    nodes.push(NodeFabric::new_events_table_node(&event.event, &event.name, &event.value, &event.timestamp))
                });

                Ok(nodes.join("\n"))
            }
        }
    }

    /// set view and updates html using document.querySelector('selector').innerHtml = $var
    async fn set_view(view: DashboardView) {
        let document = window().unwrap().document().unwrap();

        let table_node = document.query_selector("#table_data").unwrap().unwrap();
        let table_headers_node = document.query_selector("#table_headers").unwrap().unwrap();

        let html_to_render = Self::get_updated_html_by_view(view).await.unwrap();

        // add table headers
        let mut header_nodes:  Vec<String> = vec!();
        let headers = match view {
            DashboardView::State => vec!("ServiceName", "Value", "Timestamp"),
            DashboardView::Events => vec!("Events", "Name", "Value", "Timestamp")
        };
        headers.iter().for_each(| header | {
            header_nodes.push(NodeFabric::create_node(NodeCreatePayload {
                tag: "th",
                attrs: "scope='col'",
                inner: &header.to_string()
            }))
        });

        // set Html
        table_node.set_inner_html(&html_to_render);
        table_headers_node.set_inner_html(&header_nodes.join("\n"));
    }
}

#[wasm_bindgen]
pub async fn run() -> Result<(), JsValue> {
    //
    Dashboard::new().await.expect("cannot create Dashboard");
    Ok(())
}
