use web_sys::window;

use crate::{DashboardView, data::data::DashboardDataJson, api_com::api_fetch::ApiFetch, log as console_log};
use super::fabric::{Fabric as NodeFabric, CreatePayload as NodeCreatePayload};


pub struct Actions;

impl Actions {
  pub async fn set_view(view: DashboardView) -> Result<(), ()> {
    let document = window().unwrap().document().unwrap();

    let table_node = document.query_selector("#table_data").unwrap().unwrap();
    let table_headers_node = document.query_selector("#table_headers").unwrap().unwrap();

    let html_to_render = async {
        let file_content = ApiFetch::get_data().await.map_err(|_| {
            eprintln!("cannot fetch `get_data`");
            console_log("cannot fetch get_data");
        })?;
        let serialized_data = DashboardDataJson::new(file_content).map_err(|_| {
            eprintln!("cannot serialize get_data response");
            console_log("cannot serialize get_data response");
        })?;

        match view {
            DashboardView::State => {
                let mut nodes: Vec<String> = vec!();
                serialized_data.state.iter().for_each(|service_state| {
                    let node = NodeFabric::create_node(NodeCreatePayload {
                        tag: "tr",
                        attrs: "",
                        inner: &vec!(&service_state.name, &service_state.addr, &service_state.timestamp).iter()
                            .map(|i| {
                                NodeFabric::create_node(NodeCreatePayload {
                                    tag: "td",
                                    attrs: "",
                                    inner: i
                                })
                            }).collect::<Vec<String>>().join("\n")
                    });

                    nodes.push(node)
                });

                Ok(nodes.join("\n"))
            },
            DashboardView::Events => {
                let mut nodes: Vec<String> = vec!();
                serialized_data.events.iter().rev().for_each(|event| {
                    let node = NodeFabric::create_node(NodeCreatePayload {
                        tag: "tr",
                        attrs: "",
                        inner: &vec!(&event.event, &event.name, &event.value, &event.timestamp).iter()
                            .map(|i| {
                                NodeFabric::create_node(NodeCreatePayload {
                                    tag: "td",
                                    attrs: "",
                                    inner: i
                                })
                            }).collect::<Vec<String>>().join("\n")
                    });
                    nodes.push(node)
                });

                Ok(nodes.join("\n"))
            }
        }
    }.await?;

        // add table headers
    let mut header_nodes:  Vec<String> = vec!();
    let headers = match view {
        DashboardView::State => vec!("ServiceName", "Value", "Timestamp"),
        DashboardView::Events => vec!("Events", "Name", "Value", "Timestamp")
    };
    headers.iter().for_each(| header | {
      header_nodes.push(NodeFabric::create_node(NodeCreatePayload  {
          tag: "th",
          attrs: "scope='col'",
          inner: &header.to_string()
        }))
      });

    // set Html
    table_node.set_inner_html(&html_to_render);
    table_headers_node.set_inner_html(&header_nodes.join("\n"));

    Ok(())
  }

  pub async fn remove_records() -> Result<(), ()> {
    ApiFetch::clear_data().await.map_err(|err| {
        eprintln!("cannot clear data {:?}", err);
        console_log("cannot clear data");
    })?;

    Ok(())
  }
}
