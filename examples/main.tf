terraform {
  required_providers {
    hashicups = {
      version = "0.3.0"
      source  = "hashicorp.com/edu/hashicups"
    }
  }
}

provider "hashicups" {}

module "psl" {
  source = "./coffee"

  coffee_name = "Packer Spiced Latte"
}

output "psl" {
  value = module.psl.coffee
}
