use sea_orm::{DatabaseConnection, ActiveModelTrait, EntityTrait, QueryFilter, ColumnTrait};
use database::helpers::{nullable, not_null, AvailableLanguages};
use database::models::{user_infos, languages, users_languages};
use crate::result::T as Res;

pub struct FullUserModel {
  pub user: user_infos::Model,
  pub languages: Vec<languages::Model>
}

pub struct UpdateUserModel {
  pub light_bg: String,
  pub dark_bg: String,
  pub description: String,
  pub picture: String,
  pub languages: Vec<String>
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
      user_id: not_null(user_id),
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
    }
  }

  pub async fn update(&self, user_id: &String, user: UpdateUserModel) -> Res<bool> {
    // make new model that would be inserted in db
    let user_model = user_infos::ActiveModel {
      user_id: not_null(user_id.clone()),
      background_dark: nullable(user.dark_bg),
      background_light: nullable(user.light_bg),
      description: nullable(user.description),
      picture: nullable(user.picture),
      ..Default::default()
    };

    // insert
    user_model.save(&self.db).await?;

    // delete all user's languages relations
    users_languages::Entity::delete_many()
      .filter(users_languages::Column::UserInfoUserId.eq(user_id))
      .exec(&self.db).await?;

    // create new relations
    users_languages::Entity::insert_many(user.languages.iter().filter_map(| lang | {
      let language_id = AvailableLanguages::find(lang).unwrap_or(0);

      if language_id == 0 {
        return None;
      }

      let user_lang = users_languages::ActiveModel {
        language_id: not_null(language_id),
        user_info_user_id: not_null(user_id.to_string())
      };

      Some(user_lang)
    })).exec(&self.db).await?;

    Ok(true)
  }

}
