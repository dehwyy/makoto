use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Default)]
pub struct DashboardDataJson {
    pub state: Vec<ServiceAddress>,
    pub events: Vec<Events>
}

impl DashboardDataJson {
  pub fn new(s: String) -> makoto::Result<Self> {

    if s == "" {
      return Ok(DashboardDataJson::default());
    }

    Ok(serde_json::from_str::<DashboardDataJson>(&s)?)
  }
}

#[derive(Serialize, Deserialize, Clone)]
pub struct ServiceAddress {
  pub name: String,
  pub addr: String,
  pub timestamp: String
}

#[derive(Serialize, Deserialize)]
pub struct Events {
  pub event: String,
  pub name: String,
  pub value: String,
  pub timestamp: String
}
