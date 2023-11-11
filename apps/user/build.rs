use std::env;

fn main() -> Result<(), Box<dyn std::error::Error>> {

  let mut cwd = env::current_dir().unwrap();
  cwd.pop();
  cwd.pop();

  tonic_build::configure()
      .build_client(false)
      .compile(&vec!("general", "user")
                      .iter().map(|item| format!("{}/libs/grpc/protos/{}.proto", cwd.to_string_lossy(), item)).collect::<Vec<_>>()
      , &[format!("{}/libs/grpc/protos", cwd.to_string_lossy())]).unwrap();
  Ok(())
}
