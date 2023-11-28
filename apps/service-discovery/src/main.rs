use std::{time::Duration, io};

use async_nats::jetstream::{stream::{Config, RetentionPolicy}, consumer::{self, DeliverPolicy}, message, Message as NatsJetStreamMessage};
use futures::StreamExt;
use bytes::Bytes;
use config::{constants::Nats as NatsConfig, ConfigRead, Config as EnvConfig};
use logger::info;
extern crate redis;
use redis::{Connection as RedisConnection, Commands};

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error + Send + Sync>> {
  let nats_vars = NatsConfig::new();
  let config = EnvConfig::new();
  logger::Logger::init();

  // make a connection to NATS
  let client = async_nats::connect(nats_vars.server_default).await?;

  // get JetStreram context using connection
  let js = async_nats::jetstream::new(client);

  // create or get JetStream
  let stream: async_nats::jetstream::stream::Stream = js.get_or_create_stream(Config {
    name: nats_vars.js_service_discovery,
    description: Some("service discovery jetstream".to_string()),
    subjects: vec!("discovery.*".to_string()),
    max_age: Duration::from_secs(60),
    max_bytes: 1024,
    retention: RetentionPolicy::WorkQueue,
    ..Default::default()
  }).await?;

  // Create or get JetStreamConsumer
  let consumer = stream.get_or_create_consumer("discovery-consumer",consumer::pull::Config {
    durable_name: Some(nats_vars.consumer_discovery),
    description: Some("service discovery consumer".to_string()),
    ack_policy: consumer::AckPolicy::Explicit, // default
    ..Default::default()
  }).await?;

  // messages as Stream
  let mut messages = consumer.messages().await?;

  let redis_client = redis::Client::open(config.redis_base_url)?;
  let mut redis_connection = redis_client.get_connection()?;


  // Listening to the stream and do something on
  while let Some(Ok(message)) = messages.next().await {
    println!("message receiver: {:?}", message);
    message.ack().await?;


    // discovery.register || discovery.unregister
    let message_subject = message.subject.split(".").nth(1).unwrap_or_default();

    println!("message subject {}", message_subject);

    match message_subject {
      "reg" => handle_register(message, &mut redis_connection),
      "unreg" => handle_unregister(message, &mut redis_connection),
      _ => info!("wrong message subject {subject}", subject=message_subject)
    };

  }

  // js.delete_stream(nats_vars.js_service_discovery).await?;

  Ok(())
}

fn handle_register(message: NatsJetStreamMessage, redis_connection: &mut RedisConnection) {

  let (name, address) = match parse_nats_keyvalue_message(&message.payload) {
    Ok(v) => v,
    Err(e) => {
      info!("error occuried while parsing nats message {}", e.to_string());
      return;
    }
  };

  println!("name and address {} {}", name, address);

  match redis_connection.set::<_, _, ()>(&name, &address) {
    Ok(_) => {
      info!("Successfully set: {name} {address}", name=name, address=address);
    },
    Err(e) => {
      info!("Error occuried: {e}");
    }
  };
}

fn handle_unregister(message: NatsJetStreamMessage, redis_connection: &mut RedisConnection ) {

  let name = match String::from_utf8(message.payload.to_vec()) {
    Ok(v) => v,
    Err(e) => {
      info!("cannot parse message payload {}", e.to_string());
      return;
    }
  };

  match redis_connection.set::<_, _, ()>(&name, "") {
    Ok(_) => {
      info!("set empty value for {}", name);
    },
    Err(e) => {
      info!("Error occuried: {e}");
    }
  };
}

fn parse_nats_keyvalue_message(message_payload: &Bytes) -> Result<(String, String), Box<dyn std::error::Error>> {
  let message_payload = String::from_utf8(message_payload.to_vec())?;

  match message_payload.split_once(";") {
    Some(v) => Ok((v.0.to_string(), v.1.to_string())),
    None => Err(Box::new(io::Error::other("cannot parse message_payload")))
  }
}
