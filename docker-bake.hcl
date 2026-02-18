group "default" {
    targets = ["api-service"]
}


target "api-service" {
    context = "./api-service"
    dockerfile = "Dockerfile"
}
