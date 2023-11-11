terraform {
  required_providers {
    airbyte = {
      source  = "flawless/airbyte"
      version = "0.9.1"
    }
  }
}

provider "airbyte" {
  # Configuration options
}