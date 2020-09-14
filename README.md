# A simple Go application that make use of bazel

## Bazel in action

Bazel provides various features to build and deply Go applications.

The WORKSPACE file provides the Bazel rules to be used and the BUILD.bazel file provides the different tasks to be executed

To run the Go application directly on your OS:

```bash
bazel run //:go-hello
```
Once the application is running, visit [http://localhost:8080](http://localhost:8080)

To build and run the Go application within a container (making use of distroless):

```bash
bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //:image
docker run --rm -it -p8080:8080 bazel:image
```
Once the container is running, visit [http://localhost:8080](http://localhost:8080)

To push the container image to container registry:

```bash
bazel run //:push-image
```

To be able to push containers to registry, you will need to do the approtiate container registry configuration, follow the instructions from [here](https://cloud.google.com/container-registry/docs/advanced-authentication#gcloud-helper) and indicate the appropiate repository name to configure authentication [for Google Container Registry](https://cloud.google.com/container-registry/docs)

You will need to indicate directory where the config.json is located in the WORKSPACE file
