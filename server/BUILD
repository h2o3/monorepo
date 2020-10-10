load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")
load("@k8s_deploy//:defaults.bzl", "k8s_deploy")
load("@k8s_service//:defaults.bzl", "k8s_service")

go_binary(
    name = "server",
    embed = [":server_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "server_lib",
    srcs = ["main.go"],
    importpath = "h2o3.com/demo/server",
    visibility = ["//visibility:private"],
    deps = [
        "//protobuf",
        "@org_golang_google_grpc//:go_default_library",
    ],
)

go_image(
    name = "image",
    base = "@docker_base//image",
    embed = [":server_lib"],
    visibility = ["//visibility:public"],
)

k8s_deploy(
    name = "deployment",
    images = {
        "docker.io/zhzhxtrrk/fn:server": ":image",
    },
    template = ":deployment.yaml",
)

k8s_service(
    name = "service",
    template = ":service.yaml",
)
