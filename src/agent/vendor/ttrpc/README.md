# ttrpc-rust

_ttrpc-rust is a **non-core** subproject of containerd_

`ttrpc-rust` is the Rust version of [ttrpc](https://github.com/containerd/ttrpc). [ttrpc](https://github.com/containerd/ttrpc) is GRPC for low-memory environments.

The ttrpc compiler of `ttrpc-rust` `ttrpc_rust_plugin` is modified from gRPC compiler of [gRPC-rs](https://github.com/pingcap/grpc-rs) [grpcio-compiler](https://github.com/pingcap/grpc-rs/tree/master/compiler).

## Usage

### 1. Generate with `protoc` command
To generate the sources from proto files:

1. Install protoc from github.com/protocolbuffers/protobuf

2. Install protobuf-codegen
```
cargo install --force protobuf-codegen
```

3. Install ttrpc_rust_plugin from ttrpc-rust/compiler
```
cd ttrpc-rust/compiler
cargo install --force --path .
```

4. Generate the sources:

```
$ protoc --rust_out=. --ttrpc_out=. --plugin=protoc-gen-ttrpc=`which ttrpc_rust_plugin` example.proto
```


### 2. Generate programmatically

API to generate .rs files to be used e. g. from build.rs.

Example code:

```
fn main() {
    protoc_rust_ttrpc::Codegen::new()
        .out_dir("protocols")
        .inputs(&[
            "protocols/protos/agent.proto",
        ])
        .include("protocols/protos")
        .rust_protobuf() // also generate protobuf messages, not just services
        .run()
        .expect("Codegen failed.");
}
```

# async/.await
ttrpc-rust supports async/.await. By using async/.await you can reduce the overhead and resource consumption caused by threads.

## Usage
### 1. Generate codes in async version
Currently we only support generating async codes by using ttrpc-codegen

```
    ttrpc_codegen::Codegen::new()
        .out_dir("protocols/asynchronous")
        .inputs(&protos)
        .include("protocols/protos")
        .rust_protobuf()
        .customize(Customize {
            async_all: true, // It's the key option.
            ..Default::default()
        })
        .run()
        .expect("Gen async codes failed.");
```

Provide customize option
- `async_all`: generate async codes for both server and client
- `async_server`: generate async codes for server
- `async_client`: generate async codes for client

> See more in `example/build.rs`

### 2. Write your implemention in async/.await's way
Please follow the guidlines in `example/async-server.rs` and `example/async-client.rs`

# Run Examples
1. Go to the directory

    ```
    $ cd ttrpc-rust/example
    ```

2. Start the server

    ```
    $ cargo run --example server
    ```
    or

    ```
    $ cargo run --example async-server
    ```

3. Start a client

    ```
    $ cargo run --example client
    ```
    or
    ```
    $ cargo run --example async-client
    ```


# Notes: the version of protobuf
protobuf-codegen, ttrpc_rust_plugin and your code should use the same version protobuf.
You will get following fail if use the different version protobuf.
```
27 | const _PROTOBUF_VERSION_CHECK: () = ::protobuf::VERSION_2_8_0;
   |                                                 ^^^^^^^^^^^^^ help: a constant with a similar name exists: `VERSION_2_10_1`
```
The reason is that [files generated by protobuf-codegen are compatible only with the same version of runtime](https://github.com/stepancheg/rust-protobuf/commit/2ab4d50c27c4dd7803b64ce1a43e2c134532c7a6)

To fix this issue:
1. Rebuild protobuf-codegen with new protobuf:
```
cd grpc-rs
cargo clean
cargo update
cargo install --force protobuf-codegen
```
2. Rebuild ttrpc_rust_plugin with new protobuf:
```
cd ttrpc-rust/compiler
cargo clean
cargo update
cargo install --force --path .
```
3. Build your project.