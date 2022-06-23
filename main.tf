terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 2.13.0"
    }
  }
}

provider "docker" {}

resource "docker_image" "opsforqa" {
  name         = "opsforqa:latest"
  keep_locally = true
}

resource "docker_container" "opsforqa" {
  image = docker_image.opsforqa.latest
  name  = "terraform_learn"
  command = ["go", "run", "./"]
  ports {
    internal = 8080
    external = 1234
  }
}
