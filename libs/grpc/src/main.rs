use std::{fs, process::Command, time, thread, io::Read};
fn main() {
    const PATH_PROTOS: &str = "./protos";
    const PATH_GENERATED: &str = "./generated";
    const PATH_DIST: &str = "./dist";

    // 0. prebuild : rm files
    let res = fs::remove_dir_all(PATH_GENERATED);
    match res {
        Ok(_) => println!("\n1.1.Removed old generated files"),
        Err(_) => println!("\n1.1.No /generated dir to remove"),
    };

    let res = fs::remove_dir_all(PATH_DIST);
    match res {
        Ok(_) => println!("1.2.Removed old dist files\n"),
        Err(_) => println!("1.2.No /dist dir to remove\n"),
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
            "protoc --ts_out=./generated/{0} --experimental_allow_proto3_optional --ts_opt=generate_dependencies,eslint_disable --proto_path ./protos {1}",
            filename, filename_with_ext
        );

        Command::new("protoc").args(go_args.split(" ").collect::<Vec<&str>>()).spawn().unwrap();
        Command::new("npx").args(ts_args.split(" ").collect::<Vec<&str>>()).spawn().unwrap();
    });

    // 2. post-build : transpile .ts -> .js and generate .d.ts
    thread::sleep(time::Duration::from_secs(5));


    let mut file = fs::File::options().write(true).read(true).append(true).open("./generated/auth/auth.ts").unwrap();

    let mut buf = String::new();
    file.read_to_string(&mut buf).unwrap();
    fs::remove_file("generated/auth/auth.ts").unwrap();
    fs::write("generated/auth/auth.ts", "//@ts-nocheck\n".to_string() + &buf).unwrap();


    Command::new("npx").arg("tsc").spawn().unwrap(); // that's all as tsc (via typescript config) would do everything by itself


    println!("2.Successfully generated grpc files!\n");

    thread::sleep(time::Duration::from_secs(7));
}
