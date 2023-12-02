pub struct Fabric;

impl Fabric {
  pub fn new_state_table_node(name: &String, value: &String, timestamp: &String) -> String {
    format!(
        "<tr>
          <td>{0}</td>
          <td>{1}</td>
          <td>{2}</td>
        </tr>",
        name, value, timestamp
    )
  }

  pub fn new_events_table_node(event: &String, name: &String, value: &String, timestamp: &String) -> String {
    format!(
      "<tr>
        <td>{0}</td>
        <td>{1}</td>
        <td>{2}</td>
        <td>{3}</td>
      </tr>",
      event, name, value, timestamp
    )
  }

  pub fn new_headers_table_node(header: &String) -> String {
    format!(
      "<th scope='col'>{}</th>", header
    )
  }
}
