# A simple Go application that make use of bazel

## Bazel in action

Bazel provides various features to build and deply Go applications.

WORKSPACE proovides the Bazel rules to be used and BUILD.bazel provides the different tasks to be executed

To run the Go application directly on your OS:

```bash
bazel run //:go-hello
```
Once the application is running, visit [http://localhost:8080](http://localhost:8080)

To build and run the Go application within a container (making use of distroless):

```bash
bazel run //:image
```
Once the container is running, visit [http://localhost:8080](http://localhost:8080)

To push the container image to container registry:

```bash
bazel run //:push-image
```

To be able to push containers to registry, you will need to do the approtiate container registry configuration, follow instructions from [here](https://cloud.google.com/container-registry/docs/advanced-authentication#gcloud-helper) to configure authentication [for Google Container Registry](https://cloud.google.com/container-registry/docs)

You will need to indicate directory where the config.json is located in WORKSPACE
