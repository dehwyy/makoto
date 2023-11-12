use envconfig::Envconfig;
use dotenv;
use std::env::current_dir as cwd;

#[derive(Envconfig)]
pub struct Config {
    #[envconfig(from = "USER_URL")]
    pub user_url: String
}

impl Config {
    pub fn new() -> Self {
        // load env;
        let env_path = cwd().unwrap().join(".env");
        dotenv::from_path(env_path);

        Self::init_from_env().unwrap()
    }
}
