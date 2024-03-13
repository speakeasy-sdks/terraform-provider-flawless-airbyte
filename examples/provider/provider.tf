terraform {
  required_providers {
    airbyte = {
      source  = "flawless/airbyte"
      version = "0.13.2"
    }
  }
}

provider "airbyte" {
  # Configuration options
}