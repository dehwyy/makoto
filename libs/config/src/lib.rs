use envconfig::Envconfig;
use dotenv;
use std::env::current_dir as cwd;

#[derive(Envconfig, Debug, Default)]
pub struct Config {
    #[envconfig(from = "USER_URL")]
    pub user_url: String,

    #[envconfig(from = "DATABASE_URL")]
    pub database_url: String,
}

impl Config {
    pub fn new() -> Self {
        // load env;
        let env_path = cwd().unwrap().join(".env");
        let _ = dotenv::from_path(env_path);

        match Self::init_from_env() {
            Ok(val) => val,
            Err(_) => Config{
                ..Default::default()
            }
        }
    }
}

pub mod constants {
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
}
