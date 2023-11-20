use logger::info;
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

    user_model.save(&self.db).await?;

    users_languages::Entity::delete_many()
      .filter(users_languages::Column::UserInfoUserId.eq(user_id))
      .exec(&self.db).await?;



    let models: Vec<users_languages::ActiveModel> = user.languages.iter().filter_map(|language| {
      let language_id = AvailableLanguages::find(language).unwrap_or(0);

      if language_id == 0 {
        return None;
      }

      let user_lang = users_languages::ActiveModel {
        language_id: not_null(language_id),
        user_info_user_id: not_null(user_id.to_string())
      };

      Some(user_lang)
    }).collect();

    users_languages::Entity::insert_many(models).exec(&self.db).await?;

    Ok(true)
  }

}

#[cfg(test)]
mod tests {
  use sea_orm::DbErr;
  use super::*;
  use crate::Database;

  const USER_ID: &'static str = "67e55044-10b1-426f-9247-bb680e5fe0c8";
  const USER_PICTURE: &'static str = "https://google/image.jpg";

  async fn create_db() -> Result<DatabaseConnection, DbErr> {
    Database::new(String::from("postgres://postgres:postgres@localhost/postgres")).await
  }

  async fn create_repo() -> Res<UserInfoRepository> {
    let db = create_db().await?;
    Ok(UserInfoRepository::new(db))
  }

  async fn create_user() -> Res<()> {
    let repo = create_repo().await?;
    let response = repo.create(USER_ID.to_string(), USER_PICTURE.to_string()).await?;
    assert_eq!(response, true);
    Ok(())
  }

  async fn update_user() -> Res<()> {
    let repo = create_repo().await?;

    let user = UpdateUserModel {
      description: "I have description".to_string(),
      dark_bg: "".to_string(),
      light_bg:  "".to_string(),
      picture: USER_PICTURE.to_string(),
      languages: vec!("russian".to_string(), "japanese".to_string())
    };
    let response = repo.update(&USER_ID.to_string(), user).await?;

    assert_eq!(response, true);
    Ok(())
  }

  async fn get_user() -> Res<()> {
    let repo = create_repo().await?;
    let response = repo.get(USER_ID.to_string()).await?;

    match response {
      FindUserResult::Found(user) => {
        assert_eq!(user.user.picture.expect("user should have picture!"), USER_PICTURE.to_string());
        assert_eq!(user.user.description.unwrap().len() > 0, true);
        Ok(())
      },
      FindUserResult::NotFound => panic!("user should be found")
    }
  }

  #[tokio::test]
  async fn seq_test() -> Res<()> {
    create_user().await?;
    println!("Created user!");
    update_user().await?;
    println!("Updated user!");
    get_user().await?;
    println!("Get user!");

    Ok(())
  }
}
