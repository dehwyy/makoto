pub struct CreatePayload<'a> {
  pub tag: &'a str, // want "<p>", tag -> "p"
  pub attrs: &'a str, // want "class='underline' style='padding: 0 5px;'", attrs => "class='underline' style='padding: 0 5px;'"
  pub inner: &'a str
}

impl<'a> Default for CreatePayload<'_> {
  fn default() -> Self {
    Self {
      tag: "div",
      attrs: "",
      inner: ""
    }
  }
}

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

  pub fn create_node(payload: CreatePayload) -> String {
    format!(
      "<{tag} {attrs}>{inner}</{tag}>", tag=payload.tag, attrs=payload.attrs, inner=payload.inner
    )
  }
}
