
use phf::phf_map;

static LANGUAGES: phf::Map<&'static str, i64> = phf_map! {
    "russian" => 1,
    "english" => 2,
    "japanese" => 3,
};

pub struct AvailableLanguages;


impl AvailableLanguages {
  pub fn find(s: &String) -> Option<i64> {
    let s = s.to_lowercase();

    match LANGUAGES.get(s.as_str()) {
      Some(id) => Some(id.clone()),
      None => None
    }
 }
}
