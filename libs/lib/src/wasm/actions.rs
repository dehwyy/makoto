use wasm_bindgen::prelude::*;
use wasm_bindgen::closure::{Closure, IntoWasmClosure};
use web_sys::{window, MouseEvent};

pub struct EventListener;

impl EventListener {
  /// use 'move' keyword
  /// action is event_name (like 'click')
  pub fn add_mouse_event<T>(action: &str, node_selector: &str, callback: T)
  where T: FnMut(MouseEvent) + 'static
  {
      let node = window().unwrap().document().unwrap().query_selector(&node_selector).unwrap().expect("cannot get node by selector in add_mouse_event");

      let onclick = Closure::new(callback);

      node.add_event_listener_with_callback(action, onclick.as_ref().unchecked_ref()).unwrap();
      onclick.forget();
  }
}
