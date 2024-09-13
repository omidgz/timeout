terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 3.0.1"
    }
  }
}

provider "docker" {}

# Pull the official Golang Docker image
resource "docker_image" "golang_app" {
  name = "my-golang-timeout--app:latest"
  build {
    context    = "${path.module}"
    dockerfile = "${path.module}/Dockerfile"
  }
}

resource "docker_container" "golang_app" {
  image = docker_image.golang_app.image_id
  name  = "golang_timeout_project"
}
