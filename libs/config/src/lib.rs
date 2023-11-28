use envconfig::Envconfig;
use dotenv;
use std::env::current_dir as cwd;

pub trait ConfigRead {
    fn new() -> Self;
}

#[derive(Envconfig, Debug, Default)]
pub struct Config {
    #[envconfig(from = "USER_URL")]
    pub user_url: String,

    #[envconfig(from = "DATABASE_URL")]
    pub database_url: String,

    #[envconfig(from = "REDIS_BASE_URL", default = "redis://127.0.0.1:6379")]
    pub redis_base_url: String
}

impl ConfigRead for Config {
    fn new() -> Self {
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
    pub use super::*;

    #[derive(Envconfig, Debug, Default)]
    pub struct Nats {
        // ? Servers
        #[envconfig(from = "NATS_SERVER_DEFAULT", default = "localhost:4222")]
        pub server_default: String,

        // ? jetstreams
        #[envconfig(from = "NATS_JETSTREAM_SERVICE_DISCOVERY", default = "ServiceDiscovery")]
        pub js_service_discovery: String,

        // ? consumers
        #[envconfig(from = "NATS_CONSUMER_DISCOVERY", default = "DiscoveryConsumer")]
        pub consumer_discovery: String,
    }

    impl ConfigRead for Nats {
        fn new() -> Self {
            let env_path = cwd().unwrap().join(".env");
            let _ = dotenv::from_path(env_path);

            match Self::init_from_env() {
                Ok(val) => val,
                Err(_) => Nats {
                    ..Default::default()
                }
            }
        }
    }
}
