use envconfig::Envconfig as EnvConfigTrait;
use dotenv;
use std::env::current_dir as get_cwd;

pub mod constants;
pub mod hosts;
pub mod db;
pub mod secrets;

fn init_from_file<T>(filename: &str) -> T
where T: Default + EnvConfigTrait
{

    let env_path = get_cwd().unwrap().join(filename);
    let _ = dotenv::from_path(env_path);

    T::init_from_env().unwrap_or_default()
}
