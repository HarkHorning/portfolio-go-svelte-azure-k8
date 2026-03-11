variable "resource_group_name" {
  description = "Name of the Azure resource group"
  type        = string
  default     = "hark-portfolio-rg"
}

variable "location" {
  description = "Azure region to deploy resources"
  type        = string
  default     = "eastus"    # Change to region closest to you
}

variable "environment" {
  description = "Environment name (dev, staging, prod)"
  type        = string
  default     = "dev"
}

# Sensitive variables - never have defaults for these!
variable "mysql_admin_password" {
  description = "MySQL administrator password"
  type        = string
  sensitive   = true    # Won't show in logs
}