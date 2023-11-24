use sea_orm::{Database as SeoOrmDatabase, ConnectOptions, DatabaseConnection, DbErr};
use std::time::Duration;

pub mod models;
pub mod helpers;
pub struct Database;

impl Database {
  /// example: "protocol://username:password@host/database"
  pub async fn new(database_url: String) -> Result<DatabaseConnection, DbErr> {

    let mut opts = ConnectOptions::new(database_url);
    opts
      .max_connections(10)
      .connect_timeout(Duration::from_secs(10))
      .idle_timeout(Duration::from_secs(60));

    let db = SeoOrmDatabase::connect(opts).await?;

    Ok(db)
  }
}
