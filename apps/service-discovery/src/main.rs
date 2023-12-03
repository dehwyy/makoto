mod data;

use std::time::{Duration, SystemTime};

use async_nats::jetstream::{stream::{Config, RetentionPolicy}, consumer, Message as NatsJetStreamMessage};
use futures::StreamExt;
use config::constants::{nats as nats_const, redis as redis_const};
use logger::info;
extern crate redis;
use redis::{Connection as RedisConnection, Commands};
use data::data::DashboardData;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error + Send + Sync>> {
  logger::Logger::init().unwrap();

  // make a connection to NATS
  let client = async_nats::connect(nats_const::SERVER_DEFAULT).await?;

  // get JetStreram context using connection
  let js = async_nats::jetstream::new(client);

  // create or get JetStream
  let stream: async_nats::jetstream::stream::Stream = js.get_or_create_stream(Config {
    name: nats_const::JS_SERVICE_DISCOVERY.to_string(),
    description: Some("service discovery jetstream".to_string()),
    subjects: vec!("discovery.*".to_string()),
    max_age: Duration::from_secs(60),
    max_bytes: 1024,
    retention: RetentionPolicy::WorkQueue,
    ..Default::default()
  }).await?;

  // Create or get JetStreamConsumer
  let consumer = stream.get_or_create_consumer("discovery-consumer",consumer::pull::Config {
    durable_name: Some(nats_const::CONSUMER_DISCOVERY.to_string()),
    description: Some("service discovery consumer".to_string()),
    ack_policy: consumer::AckPolicy::Explicit, // default
    ..Default::default()
  }).await?;

  // messages as Stream
  let mut messages = consumer.messages().await?;

  let redis_client = redis::Client::open(redis_const::SERVER_DEFAULT)?;
  let mut redis_connection = redis_client.get_connection()?;


  // Listening to the stream and do something on
  while let Some(Ok(message)) = messages.next().await {
    println!("message receiver: {:?}", message);
    message.ack().await?;


    // discovery.reg || discovery.unreg
    let message_subject = message.subject.split(".").nth(1).unwrap_or_default();

    println!("message subject {}", message_subject);

    match message_subject {
      "reg" => {
        match handle_register(message, &mut redis_connection) {
          Ok(_) => info!("successfully registered service"),
          Err(err) => info!("register service: {}", err)
        }
      },
      "unreg" => {
        match handle_unregister(message, &mut redis_connection) {
          Ok(_) => info!("successfully unregistered service"),
          Err(err) => info!("unregister service: {}", err)
        }
      },
      _ => info!("wrong message subject {subject}", subject=message_subject)
    };

  }

  // js.delete_stream(nats_vars.js_service_discovery).await?;

  Ok(())
}

fn handle_register(message: NatsJetStreamMessage, redis_connection: &mut RedisConnection) -> makoto::Result<()> {

  let (name, address) = makoto::nats::MessageParser::key_value(&message.payload)?;

  redis_connection.hset::<_, _, _, ()>(redis_const::HASHMAP_KEY_SERVICES, &name, &address)?;
  DashboardData::insert(name.clone(), address.clone(), get_time_now())?;

  Ok(())
}

fn handle_unregister(message: NatsJetStreamMessage, redis_connection: &mut RedisConnection ) -> makoto::Result<()> {

  let name = makoto::nats::MessageParser::plain(&message.payload)?;

  redis_connection.hset::<_, _, _, ()>(redis_const::HASHMAP_KEY_SERVICES, &name, "")?;
  DashboardData::insert(name.clone(), "".to_string(), get_time_now())?;

  Ok(())
}

fn get_time_now() -> String {
  chrono::Utc::now().format("%Y-%m-%d %H:%M:%S").to_string()
}
