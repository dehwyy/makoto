use sea_orm::{DatabaseConnection, ActiveValue, ActiveModelTrait, EntityTrait, QueryFilter, ColumnTrait};
use database::helpers::{nullable};
use database::models::{user_infos, languages};
use crate::result::T as Res;

struct FullUserModel {
  user: user_infos::Model,
  languages: Vec<languages::Model>
}

pub enum FindUserResult {
  Found(FullUserModel),
  NotFound
}

#[derive(Debug)]
pub struct UserInfoRepository {
  db: DatabaseConnection
}

impl UserInfoRepository {
  pub fn new(db: DatabaseConnection) -> Self {
    Self {
      db
    }
  }

  pub async fn create(&self, user_id: String, picture: String) -> Res<bool> {
    let user = user_infos::ActiveModel {
      user_id: nullable(user_id),
      picture: nullable(picture),
      ..Default::default()
    };

    user.insert(&self.db).await?;

    Ok(true)
  }

  pub async fn get(&self, user_id: String) -> Res<FindUserResult> {
    let user = user_infos::Entity::find()
      .find_with_related(languages::Entity)
      .filter(user_infos::Column::UserId.eq(user_id))
      .all(&self.db).await?;

    let user = user.get(0);

    match user {
      Some(u) => return Ok(
        {
          let full_user=  FullUserModel {
            user: u.0.clone(),
            languages: u.1.clone()
          };

          FindUserResult::Found(full_user)
        }),
      None => return Ok(FindUserResult::NotFound)
    };

  }
}
