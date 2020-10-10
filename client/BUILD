load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")
load("@k8s_deploy//:defaults.bzl", "k8s_deploy")

go_binary(
    name = "client",
    embed = [":client_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "client_lib",
    srcs = ["main.go"],
    importpath = "h2o3.com/demo/client",
    visibility = ["//visibility:private"],
    deps = [
        "//protobuf",
        "@org_golang_google_grpc//:go_default_library",
    ],
)

go_image(
    name = "image",
    base = "@docker_base//image",
    embed = [":client_lib"],
    visibility = ["//visibility:public"],
)

k8s_deploy(
    name = "deployment",
    images = {
        "docker.io/zhzhxtrrk/fn:client": ":image",
    },
    template = ":deployment.yaml",
)
