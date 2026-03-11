output "resource_group_name" {
  description = "The name of the resource group"
  value       = azurerm_resource_group.main.name
}

output "acr_login_server" {
  description = "The URL of the container registry"
  value       = azurerm_container_registry.main.login_server
}

output "aks_cluster_name" {
  description = "The name of the AKS cluster"
  value       = azurerm_kubernetes_cluster.main.name
}