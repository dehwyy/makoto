use sea_orm::{DatabaseConnection, ActiveModelTrait, ColumnTrait, QueryFilter, EntityTrait};
use uuid::Uuid;
use chrono::{DateTime, FixedOffset};

use makoto_db::models::user_tokens::{self, Entity as UserTokens};
use makoto_db::utilities::*;

use crate::utils::jwt::{Jwt, JwtPayload};

pub struct Tokens {
  db: DatabaseConnection
}

impl Tokens {
  pub fn new(db: DatabaseConnection) -> Self {
    Self {
      db
    }
  }

  pub async fn create_new_token_pair(&self, user_id: Uuid, username: &str) -> Result<String, String> {

    let payload = JwtPayload {
      user_id: user_id.to_string(),
      username: username.to_string()
    };

    let new_access_token = Jwt::new_access_token(payload.clone())?;
    let new_refresh_token = Jwt::new_refresh_token(payload)?;

    let new_token_model = user_tokens::ActiveModel {
      access_token: nullable(vec!(new_access_token.0.clone())),
      refresh_token: not_null(new_refresh_token),
      user_id: not_null(user_id),
      expiry: not_null(new_access_token.1),
      ..Default::default()
    };
    new_token_model.insert(&self.db).await.map_err(|err| {
        err.to_string()
    })?;

    Ok(new_access_token.0)
  }
}
