# gazelle:repo bazel_gazelle
workspace(name = "com_github_eggybytes_protobuf_two_ways")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Get Go rules
RULES_GO_VERSION = "v0.25.0"  # 2020-12-02

http_archive(
  name = "io_bazel_rules_go",
  sha256 = "6f111c57fd50baf5b8ee9d63024874dd2a014b069426156c55adbf6d3d22cb7b",
  urls = [
    "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/%s/rules_go-%s.tar.gz" % (RULES_GO_VERSION, RULES_GO_VERSION),
    "https://github.com/bazelbuild/rules_go/releases/download/%s/rules_go-%s.tar.gz" % (RULES_GO_VERSION, RULES_GO_VERSION),
  ],
)

# Get Gazelle
GAZELLE_VERSION = "v0.22.2"  # 2020-10-02

http_archive(
    name = "bazel_gazelle",
    sha256 = "b85f48fa105c4403326e9525ad2b2cc437babaa6e15a3fc0b1dbab0ab064bc7c",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/%s/bazel-gazelle-%s.tar.gz" % (GAZELLE_VERSION, GAZELLE_VERSION),
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/%s/bazel-gazelle-%s.tar.gz" % (GAZELLE_VERSION, GAZELLE_VERSION),
    ],
)

# Get Protobuf rules
RULES_PROTO_VERSION = "84ba6ec814eebbf5312b2cc029256097ae0042c3"  # 2020-11-19

http_archive(
  name = "rules_proto",
  sha256 = "3bce0e2fcf502619119c7cac03613fb52ce3034b2159dd3ae9d35f7339558aa3",
  strip_prefix = "rules_proto-%s" % RULES_PROTO_VERSION,
  url = "https://github.com/bazelbuild/rules_proto/archive/%s.tar.gz" % RULES_PROTO_VERSION,
)

# Load Go rules
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(
  go_version = "1.15.6",
)

# Load Gazelle
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

# Load Protobuf rules
load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()

rules_proto_toolchains()

# Go gRPC compiler deps

GO_GRPC_COMMIT = "2cf12a82ed4381ce0b0f7d487ee2eca772ee18e3"  # v1.34.0
go_repository(
  name = "org_golang_google_grpc",
  build_file_proto_mode = "disable",
  importpath = "google.golang.org/grpc",
  strip_prefix = "grpc-go-%s" % GO_GRPC_COMMIT,
  urls = ["https://github.com/grpc/grpc-go/archive/%s.tar.gz" % GO_GRPC_COMMIT],
  sha256 = "91f8f57b4f4cc3fb7822f9a8bb7a80ef42f948c64a039a5d9d44ae247b789edd",
)

go_repository(
  name = "org_golang_google_grpc_cmd_protoc_gen_go_grpc",
  importpath = "google.golang.org/grpc/cmd/protoc-gen-go-grpc",
  sum = "h1:lQ+dE99pFsb8osbJB3oRfE5eW4Hx6a/lZQr8Jh+eoT4=",
  version = "v1.0.0",
)

# required by @org_golang_google_grpc
go_repository(
  name = "org_golang_x_net",
  importpath = "golang.org/x/net",
  strip_prefix = "net-cdfb69ac37fc6fa907650654115ebebb3aae2087",
  urls = ["https://github.com/golang/net/archive/cdfb69ac37fc6fa907650654115ebebb3aae2087.tar.gz"],
  sha256 = "24c59dcd1739ea257a5157c6de3c10d73ace3f0520a58eddf517a63d81e20bd5",
)

# required by @org_golang_x_net
go_repository(
  name = "org_golang_x_text",
  importpath = "golang.org/x/text",
  strip_prefix = "text-06d492aade888ab8698aad35476286b7b555c961",
  urls = ["https://github.com/golang/text/archive/06d492aade888ab8698aad35476286b7b555c961.tar.gz"],
  sha256 = "612e5aacbe525c1700f9a3eb3e5b230fc6beb0cca5dcf108fe88a50c11e8281c",
)

# Our own Go dependencies

go_repository(
  name = "com_github_stretchr_testify",
  importpath = "github.com/stretchr/testify",
  strip_prefix = "testify-d23661d7605ef4d6951b987c000a46be0b9ad548",
  urls = ["https://github.com/grpc/grpc-go/archive/d23661d7605ef4d6951b987c000a46be0b9ad548.tar.gz"],
  sha256 = "993ec368d934692b4234b85e4eab6bad3802a48715b41ab35f78e08b792821aa",
)
