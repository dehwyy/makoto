use jwt::{AlgorithmType,  Token,  SignWithKey, VerifyWithKey, Header, VerifyWithStore};
use std::collections::BTreeMap;
use std::time::{Duration, SystemTime, UNIX_EPOCH};
use hmac::{Hmac, Mac};

use makoto_logger::error;

const ACCESS_TOKEN_EXPIRATION_TIME_SECS: u64 = 60 * 60; // 1 hour

enum TokenKind {
  AccessToken(Duration),
  RefreshToken
}

pub struct JwtPayload {
  pub username: String,
  pub user_id: String
}

pub struct Jwt;

impl Jwt {
  pub fn new_access_token(jwt_payload: JwtPayload) -> Result<String, ()> {
    Self::sign(jwt_payload, TokenKind::AccessToken(Duration::from_secs(ACCESS_TOKEN_EXPIRATION_TIME_SECS)))
  }

  pub fn new_refresh_token(jwt_payload: JwtPayload) -> Result<String, ()> {
    Self::sign(jwt_payload, TokenKind::RefreshToken)
  }

  pub fn verify_access_token(token: String) -> Result<JwtPayload, ()> {
    let token: Token<Header, BTreeMap<String, String>, _>  = token.verify_with_key(&Self::get_jwt_signing_key()).map_err(|err| {
      error!("Error verifying token: {}", err);
    })?;

    let claims = token.claims();
    let exp_nanos = claims.get("exp").expect("cannot get key 'exp' in token's claims").parse::<u64>().expect("cannot parse to u64 'exp'");
    let exp = Duration::from_nanos(exp_nanos);

    // clarify whether token is expired
    if exp - Self::get_time_now() < Duration::from_secs(0) {
      return Err(());
    }

    Ok(
      JwtPayload {
        user_id: claims.get("user_id").expect("cannot get key 'user_id' in token's claims").to_string(),
        username: claims.get("username").expect("cannot get key 'username' in token's claims").to_string(),
      }
    )
  }


  fn sign(jwt_payload: JwtPayload, token_kind: TokenKind) -> Result<String, ()> {
    let header = Header {
      algorithm: AlgorithmType::Hs256,
      ..Default::default()
    };

    let mut claims = BTreeMap::from([
      ("username".to_string(), jwt_payload.username),
      ("user_id".to_string(), jwt_payload.user_id),
    ]);

    match token_kind {
      TokenKind::AccessToken(expiration_time) => {
        claims.insert("exp".to_string(), (Self::get_time_now() + expiration_time).as_nanos().to_string());
      }
      _ => {}
    }



    let token = Token::new(header, claims).sign_with_key(&Self::get_jwt_signing_key()).map_err(|err| {
      error!("Error signing token: {}", err);
    })?;


    Ok(token.as_str().to_string())
  }

  fn get_jwt_signing_key() -> Hmac<sha2::Sha256> {
      let jwt_secret = makoto_config::secrets::Secrets::new().jwt_secret.expect("cannot retrieve jwt secret from env!");

      Hmac::new_from_slice(jwt_secret.as_bytes()).expect("invalid length in generate hmac key! (according to docs)") as Hmac<sha2::Sha256>
  }

  fn get_time_now() -> Duration {
    SystemTime::now().duration_since(UNIX_EPOCH).expect("Time went back??")
  }
}
