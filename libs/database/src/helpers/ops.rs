use sea_orm::{ActiveValue, Value, sea_query::Nullable};

pub fn nullable<T: Into<Value> + Nullable>(value: T) -> ActiveValue<Option<T>> {
  ActiveValue::Set(Some(value))
}

pub fn not_null<T: Into<Value>>(value: T) -> ActiveValue<T> {
  ActiveValue::Set(value)
}
