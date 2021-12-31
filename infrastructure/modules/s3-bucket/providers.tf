terraform {
  required_version = "1.0.8"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "3.63.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.1.0"
    }
  }
}
