load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "hello_proto",
    srcs = ["hello.proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "hello_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "h2o3.com/demo/protobuf",
    proto = ":hello_proto",
    visibility = ["//visibility:public"],
)

go_library(
    name = "protobuf",
    srcs = ["protobuf.go"],
    embed = [":hello_go_proto"],
    importpath = "h2o3.com/demo/protobuf",
    visibility = ["//visibility:public"],
)
