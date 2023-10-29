use envconfig::Envconfig;
use dotenv;
use std::env::current_dir as cwd;

#[derive(Envconfig)]
pub struct Config {
    #[envconfig(from = "DISCORD_CLIENT_ID")]
    pub discord_client_id: String,

    #[envconfig(from = "DISCORD_CLIENT_SECRET")]
    pub discord_client_secret: String,
}

impl Config {
    pub fn new() -> Self {
        // load env;
        let env_path = cwd().unwrap().join(".env");
        dotenv::from_path(env_path).ok();

        Self::init_from_env().unwrap()
    }
}
