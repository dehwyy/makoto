use sea_orm::{DatabaseConnection, prelude::Uuid, ColumnTrait, QueryFilter, EntityTrait, ActiveModelTrait};
use makoto_db::models::user_credentials::{Entity as UserCredentials, self};
use makoto_db::utilities::*;

#[derive(Default)]
pub struct UserPayload {
  pub user_id: Uuid,
  pub username: String,
  pub email: String,
  pub password: String
}

pub struct Credentials {
  db: DatabaseConnection
}

impl Credentials {
  pub fn new(db: DatabaseConnection) -> Self {
    Self {
      db
    }
  }

  pub async fn create_user(&self, user_payload: UserPayload) -> Result<user_credentials::Model, String> {
    let user = user_credentials::ActiveModel {
      id: not_null(user_payload.user_id),
      username: not_null(user_payload.username),
      email: not_null(user_payload.email),
      password: nullable(user_payload.password),
      ..Default::default()
    };

    Ok(user.insert(&self.db).await.map_err(|err| err.to_string())?)
  }

  pub async fn is_username_available(&self, username: &str) -> Result<bool, String> {
    let user = UserCredentials::find()
          .filter(
            user_credentials::Column::Username.eq(username)
          ).one(&self.db)
          .await.map_err(|msg| msg.to_string())?;

    match user {
      Some(_) => Ok(false),
      None => Ok(true)
    }
  }

  pub async fn is_email_available(&self, email: &str) -> Result<bool, String> {
    let user = UserCredentials::find()
          .filter(
            user_credentials::Column::Email.eq(email)
          ).one(&self.db)
          .await.map_err(|msg| msg.to_string())?;

    match user {
      Some(_) => Ok(false),
      None => Ok(true)
    }
  }
}
