use std::{fs, process::Command};
fn main() {
    const PATH_PROTOS: &str = "./protos";
    const PATH_GENERATED: &str = "./generated";

    // 0. prebuild : rm files
    let res = fs::remove_dir_all(PATH_GENERATED);
    match res {
        Ok(_) => println!("Removed old generated files"),
        Err(_) => println!("No /generated dir to remove"),
    };


    // 1. build
    fs::read_dir(PATH_PROTOS).unwrap().for_each(|e| {
        let filename_with_ext = &e.unwrap().file_name().to_str().unwrap().to_string();
        let filename = &filename_with_ext.split(".").next().unwrap().to_string();

        // create appropirate dirs
        fs::create_dir_all(PATH_GENERATED.to_owned() + "/" + filename).unwrap();

        let go_args = format!(
            "--go_out=./generated/{0} --twirp_out=./generated/{0} --go_opt=paths=source_relative --twirp_opt=paths=source_relative --proto_path=./protos {1}",
            filename, filename_with_ext
        );

        let ts_args = format!(
            "--ts_out=./generated/{0} --ts_opt=generate_dependencies,eslint_disable --proto_path ./protos {1}",
            filename, filename_with_ext
        );

        vec![go_args, ts_args].iter().for_each(|arg| {
            Command::new("protoc").args(arg.split(" ").collect::<Vec<&str>>()).spawn().unwrap();
        });
    });

    // 2. post-build : transpile ts -> js
    Command::new("npx").arg("tsc").spawn().unwrap(); // that's all cuz tsc (via typescript config) would do everything by itself

    println!("");
    println!("Successfully generated grpc files!");
}
