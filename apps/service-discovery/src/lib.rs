mod data;
mod node;
mod api_com;


use node::actions::Actions;
use wasm_bindgen::prelude::*;
use wasm_bindgen_futures::spawn_local;
use web_sys::{MouseEvent, window};
use makoto::wasm::EventListener;

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
        Actions::set_view(DashboardView::State).await.expect("cannot set_view");

        dashboard.setup_listeners();
        Ok(dashboard)

    }

    // setup listener (once)
    fn setup_listeners(&self) {

        // SetState button
        EventListener::add_mouse_event("click", "#set_state", move |_| {
            spawn_local(async {
                Actions::set_view(DashboardView::State).await.unwrap()
            });
        });

        // SetEvents button
        EventListener::add_mouse_event("click", "#set_events", move |_| {
                spawn_local(async {
                    Actions::set_view(DashboardView::Events).await.unwrap()
                });
        });

        // RemoveRecords button
         EventListener::add_mouse_event("click", "#btn_remove", move |_| {
                spawn_local(async {
                    Actions::remove_records().await.unwrap();
                    Actions::set_view(DashboardView::State).await.unwrap()
                });
        });

   }
}

#[wasm_bindgen]
pub async fn run() -> Result<(), JsValue> {
    //
    Dashboard::new().await.expect("cannot create Dashboard");
    Ok(())
}
