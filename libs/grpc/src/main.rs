use std::{env, process::Command, io, fs, path::Path};
//
fn main() -> io::Result<()> {
    const TWIRP_COMPILE: &str = "TWIRP_GO_COMPILE";
    const GRPC_GO_COMPILE: &str = "GRPC_GO_COMPILE";
    const TS_COMPILE: &str = "TWIRP_TS_COMPILE";

    if env::var(TWIRP_COMPILE).is_ok() {
        let protos = vec!("auth", "general", "hashmap", "user");
        let out = "generated";
        let get_command = | out: String, name: &str | {
            format!("protoc --go_out={0} --experimental_allow_proto3_optional --twirp_out={0} --go_opt=paths=source_relative --twirp_opt=paths=source_relative --proto_path=protos {1}.proto",
            out, name)
        };

        compile(protos, out, get_command)?;
    }

    if env::var(TS_COMPILE).is_ok() {
        let protos = vec!("auth", "general", "hashmap", "user");
        let out = ".ts/generated";
        let get_command = |out: String, name: &str| {
            format!("npx protoc --ts_out={0} --experimental_allow_proto3_optional --ts_opt=generate_dependencies,eslint_disable,ts_nocheck,output_javascript --proto_path protos {1}.proto",
            out, name)
        };

        compile(protos, out, get_command)?;
    }

    if env::var(GRPC_GO_COMPILE).is_ok() {
        let protos = vec!("general", "user");
        let out = "generated";
        let get_command = | out: String, name: &str | {
            format!("protoc --go_out={0} --experimental_allow_proto3_optional --go_opt=paths=source_relative --go-grpc_out={0} --go-grpc_opt=paths=source_relative --proto_path=protos {1}.proto",
            out, name)
        };

        compile(protos, out, get_command)?;
    }

    Ok(())
}

fn compile(
    protos: Vec<&str>,
    out: &str,
    get_command: impl Fn(String, &str) -> String
    ) -> io::Result<()> {
    let out_dir = Path::new(out);

    protos.iter().for_each(|name|
        {
            let out_dir = create_dir(&out_dir, name).unwrap();
            run_command(&get_command(out_dir, name)).unwrap();
        }
    );

    Ok(())
}

fn create_dir(path: &Path, name: &str) -> io::Result<String> {
    let new_dir = path.join(name);

    fs::create_dir_all(&new_dir)?;
    Ok(new_dir.to_str().unwrap().to_string())
}

fn run_command(command: &String) -> io::Result<()> {
    Command::new(command.split(" ").next().unwrap())
    .args(&mut command.split(" ").collect::<Vec<&str>>()[1..])
    .spawn().unwrap()
    .wait()?;

    Ok(())
}
