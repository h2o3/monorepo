exports_files([
    "WORKSPACE",
    "go.mod",
    "go.sum",
])

load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(name = "gazelle")

##############
# build settings
##############
load("@bazel_skylib//rules:common_settings.bzl", "string_flag")

string_flag(
    name = "env",
    build_setting_default = "local",
    values = [
        "local",
        "server",
    ],
)

config_setting(
    name = "local",
    flag_values = {
        "env": "local",
    },
    visibility = ["//visibility:public"],
)
