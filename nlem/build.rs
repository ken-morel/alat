fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_prost_build::configure()
        .build_server(true)
        .build_client(true)
        .out_dir("src")
        .compile_protos(
            &[
                "src/proto/alat.proto",
                "src/proto/pair.proto",
                "src/proto/telemetry.proto",
            ], // main proto files
            &["src/proto"], // import dir
        )?;
    Ok(())
}
