# General Rules
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")


# Rules Docker
http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "4521794f0fba2e20f3bf15846ab5e01d5332e587e9ce81629c7f96c793bb7036",
    strip_prefix = "rules_docker-0.14.4",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.14.4/rules_docker-v0.14.4.tar.gz"],
)

load("@io_bazel_rules_docker//repositories:repositories.bzl", container_repositories = "repositories")
container_repositories()

load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")
container_deps()

load("@io_bazel_rules_docker//repositories:pip_repositories.bzl", "pip_deps")
pip_deps()

load("@io_bazel_rules_docker//go:image.bzl", _go_image_repos = "repositories")
_go_image_repos()

load("@io_bazel_rules_docker//toolchains/docker:toolchain.bzl", docker_toolchain_configure="toolchain_configure")
docker_toolchain_configure(
  name = "docker_config",
  # Replace this with an absolute path to a directory which has a custom docker
  # client config.json. Note relative paths are not supported.
  # Docker allows you to specify custom authentication credentials
  # in the client configuration JSON file.
  # See https://docs.docker.com/engine/reference/commandline/cli/#configuration-files
  # for more details.
  client_config="/tmp/docker",
)

# Rules Go
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "7f1aa43d986df189f7cf30e81dd2dc9d8ed7c74e356341a17267f6d7e5748382",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.24.1/rules_go-v0.24.1.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.24.1/rules_go-v0.24.1.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
go_rules_dependencies()
go_register_toolchains()


# Gazelle
http_archive(
    name = "bazel_gazelle",
    sha256 = "d4113967ab451dd4d2d767c3ca5f927fec4b30f3b2c6f8135a2033b9c05a5687",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.0/bazel-gazelle-v0.22.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.0/bazel-gazelle-v0.22.0.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
gazelle_dependencies()
