# protobuf-two-ways

This repository demonstrates two ways of using Go's [new protobuf API](https://blog.golang.org/protobuf-apiv2) to extend your protobufs and make them more powerful: 
1. Code generation at compile-time, with [`protogen`](https://pkg.go.dev/google.golang.org/protobuf/compiler/protogen) — code in [go/plugin/](./go/plugin/)
2. Reflection at runtime, with [`protoreflect`](https://pkg.go.dev/google.golang.org/protobuf/reflect/protoreflect) — code in [go/reflect/](./go/reflect/)

This also demonstrates building protobufs with Bazel, which makes invoking `protoc` (the protobuf compiler) a smooth part of the development workflow. See [protos/](./protos/) and `BUILD.bazel` files throughout this repo for details.

# For more details

Read more details about this in [the corresponding blog post](https://eggybits.com/posts/proto-two-ways/) on eggybits.com.

# To run and build as-written

## Install [Bazelisk](https://github.com/bazelbuild/bazelisk)

Set up `bazelisk`:
```shell
# if on a Mac
brew tap bazelbuild/tap
brew install bazelisk
```

Verify that `bazel` points to `bazelisk`:
```shell
$ which bazel
/usr/local/bin/bazel
melinda at xmbp in ~/code/eggybytes/protobuf-two-ways on ml-add-readme*
$ ls -l /usr/local/bin/bazel
lrwxr-xr-x  1 melinda  admin  34 Dec  2 11:33 /usr/local/bin/bazel -> ../Cellar/bazelisk/1.7.3/bin/bazel
```

## Build everything

```shell
bazel build //...
```

## Inspect generated files

In `bazel-out/.../protos/todo/...`, you'll find:  
- **todo.pb.go**: standard protoc-gen-go output
- **todo_grpc.pb.go**: standard protoc-gen-grpc output
- **todo.pb.custom.go**: our custom code output
