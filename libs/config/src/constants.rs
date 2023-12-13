pub mod nats {
  pub const SERVER_DEFAULT: &str = "localhost:4222";
  pub const JS_SERVICE_DISCOVERY: &str = "ServiceDiscovery";
  pub const CONSUMER_DISCOVERY: &str = "DiscoveryConsumer";
}

pub mod redis {
  pub const SERVER_DEFAULT: &str = "redis://127.0.0.1:6379";
  pub const HASHMAP_KEY_SERVICES: &str = "services";
  pub const STREAM_KEY_SERVICES_EVENTS: &str = "services_events";
}
