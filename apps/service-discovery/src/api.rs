mod data;

use std::{sync::Arc, collections::HashMap};
use tokio::net::TcpListener;
use axum::{routing::get,response::Json as JsonResponse, http::StatusCode, extract::State, Json};
use redis::{streams as RedisStreams, Commands, Client as RedisClient};

use data::data::{DashboardDataJson, ServiceAddress, Events};
use config::constants::redis as redis_const;
use tower_http::cors::{CorsLayer, Any};


const PORT: &str = "4223";

#[derive(Clone)]
struct AppState {
  redis_client: RedisClient
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>>  {
  let redis_client = redis::Client::open(redis_const::SERVER_DEFAULT)?;

  let shared_state = Arc::new(AppState {
    redis_client
  });

  let api_router = axum::Router::new()
        .route("/data", get(Endpoints::get_data)
                                           .delete(Endpoints::delete_data))
                                           .with_state(shared_state);

  let cors = CorsLayer::new()
    .allow_headers(Any)
    .allow_origin(Any)
    .allow_methods(Any)
    .allow_origin(Any);

  let app = axum::Router::new()
      .nest("/api", api_router)
      .layer(cors);


  let listener = TcpListener::bind(format!("127.0.0.1:{PORT}")).await.map_err(|err| {
    eprintln!("ERROR: could not bind on port {PORT}: {err}");
    err
  })?;

  axum::serve(listener, app).await.map_err(|err| {
    eprintln!("Error: cannot serve axum {err}");
    err
  })?;

  Ok(())
}

struct Endpoints;

impl Endpoints {
  /// deletes all dashboard data (state and events)
  async fn delete_data(State(state): State<Arc<AppState>>) -> Result<StatusCode, StatusCode> {
    let mut redis_conn = state.redis_client.get_connection().map_err(|err| {
      eprintln!("cannot connect to redis! {err}");
      StatusCode::INTERNAL_SERVER_ERROR
    })?;

    let _ = redis_conn.del::<_, ()>(redis_const::HASHMAP_KEY_SERVICES);
    let _ = redis_conn.del::<_, ()>(redis_const::STREAM_KEY_SERVICES_EVENTS);

    Ok(StatusCode::OK)
  }

  /// returns DashboardDataJson
  async fn get_data(State(state): State<Arc<AppState>>) -> Result<Json<DashboardDataJson>, StatusCode> {
    let mut redis_conn = state.redis_client.get_connection().map_err(|err| {
      eprintln!("cannot connect to redis! {err}");
      StatusCode::INTERNAL_SERVER_ERROR
    })?;


    let state = {
      // format: "key": "value,timestamp"
      let values: HashMap<String, String> = redis_conn.hgetall(redis_const::HASHMAP_KEY_SERVICES).map_err(|err| {
        eprintln!("cannot hget dashboard {err}");
        StatusCode::INTERNAL_SERVER_ERROR
      })?;

      let mut v: Vec<ServiceAddress> = vec!();
      for (key, value) in  values.iter() {
        let (addr, timestamp) = value.split_once(",").expect("cannot split_once value in redis_services");
        v.push(ServiceAddress { name: key.clone(), addr: addr.to_string(), timestamp: timestamp.to_string() });
      }

      v.sort_by(|a, b| a.name.cmp(&b.name));

      v
    };

    let events: Vec<Events> = {
      let events: RedisStreams::StreamReadReply= redis_conn.xread(&[redis_const::STREAM_KEY_SERVICES_EVENTS], &["0-0"]).map_err(|err| {
        eprintln!("cannot hget dashboard {err}");
        StatusCode::INTERNAL_SERVER_ERROR
      })?;



      let mut events_values: Vec<Events> = vec!();

      // if array is empty
      if events.keys.len() == 0 {
        return Ok(JsonResponse(DashboardDataJson {
          events: vec!(),
          state
        }));
      }


      for stream_record in events.keys.get(0).expect("cannot access events.keys at index 0").ids.iter() {
        let v = vec!("key", "address", "timestamp", "event").iter().map(|key| {
          let v =  stream_record.get::<String>(key);
          v.expect("cannot get value by key")
        }).collect::<Vec<String>>();

        // no validation as if was done before!
        events_values.push(Events {
          event: v[3].clone(),
          name: v[0].clone(),
          value: v[1].clone(),
          timestamp: v[2].clone()
        });
     }

      events_values
    };

    Ok(JsonResponse(DashboardDataJson {
      events,
      state
    }))

  }
}
