terraform {
  required_providers {
    getter = {
      version = "0.0.2"
      source = "github.com/chilicat/getter"
    }
  }
}

provider "getter" {}

resource "getter_get" "myfile" {
  url = "https://raw.githubusercontent.com/chilicat/terraform-provider-pkcs12/master/main.go"
  dest = "./tmp/download2/"
}

