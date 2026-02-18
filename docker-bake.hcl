group "default" {
    targets = ["api-service"]
}

# Special target: https://github.com/docker/metadata-action?tab=readme-ov-file#bake-definition
#
target "docker-metadata-action" {}

target "api-service" {
    inherits = ["docker-metadata-action"]
    context = "./api-service"
    dockerfile = "Dockerfile"
}
