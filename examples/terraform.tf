terraform {
  required_version = ">= 0.13.0"
  required_providers {
    hashicups = {
      versions = ["0.3"]
      source = "hashicorp.com/edu/hashicups"
    }
  }
}