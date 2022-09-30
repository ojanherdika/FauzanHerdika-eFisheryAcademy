job "ojanherdika-echo" {
  datacenters = ["dc1"]
  type = "service"

  group "web" {
    count = 1

    network {
      port "http" {
        to = 1323
      }
    }

    task "ojanherdika-echo" {
      driver = "docker"

      config {
        image = "ojanherdika/go-echo:v1"
        ports = ["http"]
      }

      resources {
        cpu    = 100
        memory = 128
      }
    }

    service {
      name = "ojanherdika-echo"
      port = "http"
      tags = [
        "traefik.enable=true",
        "traefik.http.routers.ojanherdika-echo-demo.rule=Host(\"ojanherdika.cupang.efishery.ai\")",
      ]
      check {
        port        = "http"
        type        = "tcp"
        interval    = "15s"
        timeout     = "14s"
      }
    }

  }
}