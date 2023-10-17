use std::{fs, process::Command, thread, time};
fn main() {
    const PATH_PROTOS: &str = "./protos";
    const PATH_GENERATED: &str = "./generated";

    // 0. prebuild : rm files
    let res = fs::remove_dir_all(PATH_GENERATED);
    match res {
        Ok(_) => println!("\n1.Removed old generated files"),
        Err(_) => println!("\n1.No /generated dir to remove"),
    };


    // 1. build
    fs::read_dir(PATH_PROTOS).unwrap().for_each(|e| {
        let filename_with_ext = &e.unwrap().file_name().to_str().unwrap().to_string();
        let filename = &filename_with_ext.split(".").next().unwrap().to_string();

        // create appropirate dirs
        fs::create_dir_all(PATH_GENERATED.to_owned() + "/" + filename).unwrap();

        let go_args = format!(
            "--go_out=./generated/{0} --experimental_allow_proto3_optional --twirp_out=./generated/{0} --go_opt=paths=source_relative --twirp_opt=paths=source_relative --proto_path=./protos {1}",
            filename, filename_with_ext
        );

        let ts_args = format!(
            "--ts_out=./generated/{0} --experimental_allow_proto3_optional --ts_opt=generate_dependencies,eslint_disable --proto_path ./protos {1}",
            filename, filename_with_ext
        );

        vec![go_args, ts_args].iter().for_each(|arg| {
            Command::new("protoc").args(arg.split(" ").collect::<Vec<&str>>()).spawn().unwrap();
        });
    });

    // 2. post-build : transpile ts -> js
    thread::sleep(time::Duration::from_millis(150));
    Command::new("bun").args(["run", "ts"]).spawn().unwrap(); // that's all cuz tsc (via typescript config) would do everything by itself

    println!("2.Successfully generated grpc files!\n");

    // круто костыль круто люблю tsc
    thread::sleep(time::Duration::from_millis(1500)); // waiting for tsc
    println!("\n3.Generated .d.ts");
}
