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

  pub async fn create_new_access_token(&self, user_id: Uuid, username: &str) -> Result<String, String> {

    let payload = JwtPayload {
      user_id: user_id.to_string(),
      username: username.to_string()
    };

    let mut token_record: user_tokens::ActiveModel = self.get_token_model_by_user_id(user_id.clone()).await.map_err(|msg| msg.to_string())?.into();

    let new_access_token = Jwt::new_access_token(payload)?;

    // get all previous token (safe operations)
    let mut old_tokens = token_record.access_token.take().unwrap_or_default().unwrap_or_default();

    old_tokens.push(new_access_token.0.clone());

    token_record.access_token = nullable(old_tokens);
    token_record.expiry = not_null(new_access_token.1);

    token_record.update(&self.db).await.map_err(|err| err.to_string())?;

    Ok(new_access_token.0)
  }

  pub async fn get_user_id_by_access_token(&self, access_token: &str) -> Result<Uuid, String> {
    todo!()
  }

  async fn get_token_model_by_user_id(&self, user_id: Uuid) -> Result<user_tokens::Model, String> {

    let token_record = UserTokens::find().filter(
      user_tokens::Column::UserId.eq(user_id)
    ).one(&self.db).await.map_err(|msg| msg.to_string())?;

    match token_record {
      Some(token) => Ok(token),
      None => Err("token not found".to_string())
    }

  }
}
