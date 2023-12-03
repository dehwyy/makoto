use serde::{Deserialize, Serialize};
use std::{io::prelude::*, fs::File, collections::BTreeMap};

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

pub struct DashboardData;

impl DashboardData {
  /// Add to Events and insert into ServiceAddress if key doesn't exist, else -> update key's value
  pub fn insert(key: String, value: String, timestamp: String) -> makoto::Result<()> {
    // read file as string and parse it to struct
    let file_content = DashboardDataJson::new(DashboardData::read_file_as_string()?)?;

    // init emtpy binary tree map (sorted hashmap)
    let mut file_state: BTreeMap<String, ServiceAddress> = BTreeMap::new();

    // store existing value to hashmap
    for val in file_content.state.iter() {
      let _ = file_state.insert(val.name.clone(),
        ServiceAddress {
          name: val.name.clone(),
          timestamp: val.timestamp.clone(),
          addr: val.addr.clone()
        }
      );
    }

    // add new ServiceAddress to state
    let _ = file_state.insert(key.clone(),
      ServiceAddress {
        name: key.clone(),
        addr: value.clone(),
        timestamp: timestamp.clone(),
      }
    );

    // append to events list
    let mut file_events: Vec<Events> = file_content.events;
    file_events.push(Events {
        event: (if value == "" {"unregister"} else {"register"}).to_string(),
        name: key,
        value,
        timestamp
      }
    );

    // create new struct
    let new_file_content = DashboardDataJson {
      state: file_state.values().map(|val| val.clone()).collect(),
      events: file_events
    };

    // struct -> json
    let new_content = serde_json::to_string_pretty(&new_file_content)?;

    // open clear file (either created or truncated)
    let mut file = File::options().write(true).truncate(true).create(true).open("data.json")?;
    file.write(new_content.as_bytes())?;

    Ok(())
  }

  pub fn read_file_as_string() -> makoto::Result<String> {
    let mut f = File::options().read(true).write(true).create(true).open("data.json")?;
    let mut s = String::new();

    f.read_to_string(&mut s)?;

    Ok(s)
  }
}
