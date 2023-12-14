use envconfig::Envconfig;
use super::init_from_file;

#[derive(Envconfig, Default, Debug)]
pub struct Secrets {
    #[envconfig(from = "JWT_SECRET")]
    pub jwt_secret: Option<String>,
}

impl Secrets {
    pub fn new() -> Self {
        init_from_file(".env")
    }
}
