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
  pub fn create_node(payload: CreatePayload) -> String {
    format!(
      "<{tag} {attrs}>{inner}</{tag}>", tag=payload.tag, attrs=payload.attrs, inner=payload.inner
    )
  }
}
