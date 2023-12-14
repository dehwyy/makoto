use sea_orm::DatabaseConnection;

pub struct Credentials {
  db: DatabaseConnection
}

impl Credentials {
  async fn new(db: DatabaseConnection) -> Self {
    Self {
      db
    }
  }
}
