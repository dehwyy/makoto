use sea_orm::{ActiveValue, Value, sea_query::Nullable};

pub fn nullable<T: Into<Value> + Nullable>(value: T) -> ActiveValue<Option<T>> {
  ActiveValue::Set(Some(value))
}
