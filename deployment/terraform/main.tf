
terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 3.0"
    }
  }
}

provider "azurerm" {
  features {}
}

# Redo variable names here later
resource "azurerm_resource_group" "main" {
  name     = var.resource_group_name    # Comes from variables.tf
  location = var.location

  tags = {
    environment = var.environment
    project     = "hark-portfolio"
  }
}