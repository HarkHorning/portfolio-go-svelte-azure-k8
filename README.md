# Hark Art Portfolio

A scalable art portfolio and sales platform built with modern technologies and cloud-native deployment practices.

## Tech Stack

| Layer | Technology | Purpose |
|-------|------------|---------|
| Frontend | Svelte + TypeScript | Lightweight, compiled UI framework with type safety |
| Backend | Go | Fast, simple API server |
| Database | MySQL | Relational data (users, orders, art metadata) |
| Image Storage | Azure Blob Storage | Art images and media files |
| Containerization | Podman | Consistent environments (Docker-compatible, rootless) |
| Orchestration | Kubernetes (AKS) | Container orchestration, interview demo |
| Infrastructure | Terraform | Infrastructure as Code |
| CI/CD | GitHub Actions | Automated builds and deployments |

---

## Project Structure

```
hark/
├── portfolio/
│   ├── frontend/          # Svelte + TypeScript application
│   ├── backend/           # Go API server
│   ├── deployment/
│   │   ├── docker/        # Dockerfiles and podman-compose
│   │   ├── kubernetes/    # K8s manifests
│   │   └── terraform/     # Azure infrastructure
│   └── .github/
│       └── workflows/     # CI/CD pipelines
```

---

## Deployment Architecture

### Environment Strategy

| Environment | Infrastructure | Purpose | Cost |
|-------------|---------------|---------|------|
| Local Dev | Podman Compose | Daily development | Free |
| Interview Demo | AKS + Azure services | Demonstrate DevOps skills | ~$5-15/day when running |

---

## Detailed Deployment Plan

### Phase 1: Local Development Setup

**Goal**: Get the app running locally with Podman Compose.

#### Steps:

1. **Create Dockerfiles**
   - `frontend/Dockerfile` - Multi-stage build: Node for building, Nginx for serving
   - `backend/Dockerfile` - Multi-stage build: Go builder, minimal Alpine for runtime

2. **Create docker-compose.yml**
   ```yaml
   services:
     frontend:    # Svelte/TypeScript app on port 3000
     backend:     # Go API on port 8080
     mysql:       # MySQL 8.0 on port 3306
     blob-mock:   # Azurite (local Azure Blob emulator) on port 10000
   ```

3. **Local development workflow**
   ```bash
   podman-compose up -d          # Start all services
   podman-compose logs -f        # Watch logs
   podman-compose down           # Stop all services
   ```

**Cost**: Free

---

### Phase 2: Azure Infrastructure with Terraform

**Goal**: Define all Azure resources as code for reproducible deployments.

#### What is Terraform?

Terraform is an Infrastructure as Code (IaC) tool. Instead of clicking through the Azure Portal to create resources, you write configuration files that describe what you want. Benefits:

- **Reproducible**: Run the same config to recreate identical infrastructure
- **Version controlled**: Track infrastructure changes in Git
- **Documentable**: The config files ARE the documentation
- **Destroyable**: Tear down everything with one command (saves money!)

#### Prerequisites

Before starting, install these tools:

1. **Terraform CLI**
   ```bash
   # Windows (using winget)
   winget install HashiCorp.Terraform

   # Verify installation
   terraform --version
   ```

2. **Azure CLI**
   ```bash
   # Windows (using winget)
   winget install Microsoft.AzureCLI

   # Verify installation
   az --version
   ```

3. **Azure Account**
   - Create a free account at https://azure.microsoft.com/free/
   - Free tier includes $200 credit for 30 days

#### Step 1: Authenticate with Azure

```bash
# Login to Azure (opens browser)
az login

# Verify you're logged in and see your subscription
az account show

# If you have multiple subscriptions, set the one to use
az account set --subscription "Your Subscription Name"
```

#### Step 2: Understand the Terraform Files

Each `.tf` file has a specific purpose. Terraform reads ALL `.tf` files in a directory and combines them.

```
deployment/terraform/
├── main.tf              # Provider config, resource group
├── variables.tf         # Input variables (like function parameters)
├── outputs.tf           # Output values (URLs, passwords to display after creation)
├── acr.tf               # Azure Container Registry
├── storage.tf           # Blob storage for art images
├── aks.tf               # Kubernetes cluster
├── mysql.tf             # Database (optional)
└── terraform.tfvars     # YOUR values (gitignored, never commit!)
```

**File explanations:**

| File | Purpose | Example Content |
|------|---------|-----------------|
| `main.tf` | Configures Azure provider, creates resource group | `provider "azurerm" { ... }` |
| `variables.tf` | Declares variables with types and defaults | `variable "location" { default = "eastus" }` |
| `terraform.tfvars` | YOUR actual values for variables | `location = "westus2"` |
| `outputs.tf` | Values to display after `terraform apply` | `output "acr_url" { value = azurerm_container_registry.main.login_server }` |
| `acr.tf` | Container registry resource definition | `resource "azurerm_container_registry" "main" { ... }` |

#### Step 3: Create the Terraform Files

Create the directory structure:
```bash
mkdir -p deployment/terraform
cd deployment/terraform
```

**main.tf** - Provider and resource group:
```hcl
# Tell Terraform we're using Azure
terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 3.0"    # Use version 3.x
    }
  }
}

# Configure the Azure provider
provider "azurerm" {
  features {}    # Required, even if empty
}

# Resource Group - a container for all your Azure resources
# Like a folder that holds everything together
resource "azurerm_resource_group" "main" {
  name     = var.resource_group_name    # Comes from variables.tf
  location = var.location               # Azure region (eastus, westus2, etc.)

  tags = {
    environment = var.environment
    project     = "hark-portfolio"
  }
}
```

**variables.tf** - Input variables:
```hcl
# Variables are like function parameters
# They make your config reusable and keep secrets out of code

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
```

**terraform.tfvars** - Your actual values (ADD TO .gitignore!):
```hcl
# This file contains YOUR specific values
# NEVER commit this file to Git!

resource_group_name  = "hark-portfolio-rg"
location             = "eastus"
environment          = "dev"
mysql_admin_password = "YourSecurePassword123!"
```

**outputs.tf** - Values to display after creation:
```hcl
# Outputs display useful information after terraform apply

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

# Don't output sensitive values! Use Azure CLI to retrieve them.
```

#### Step 4: Initialize and Apply

```bash
cd deployment/terraform

# 1. Initialize Terraform (downloads Azure provider, creates state file)
#    Run this once, or after adding new providers
terraform init

# 2. Format your files (optional but recommended)
terraform fmt

# 3. Validate syntax
terraform validate

# 4. Preview what will be created (NO COST - just a dry run)
#    Review this carefully before applying!
terraform plan

# 5. Create the infrastructure (THIS STARTS COSTING MONEY!)
#    Type "yes" when prompted
terraform apply

# 6. When done/to save money, destroy everything
terraform destroy
```

#### Key Terraform Commands Explained

| Command | What It Does | When to Use |
|---------|--------------|-------------|
| `terraform init` | Downloads providers, sets up state | First time, or after adding providers |
| `terraform fmt` | Auto-formats `.tf` files | Before committing code |
| `terraform validate` | Checks syntax errors | Before plan/apply |
| `terraform plan` | Shows what WILL happen (dry run) | Always before apply |
| `terraform apply` | Creates/updates real resources | When ready to deploy |
| `terraform destroy` | Deletes ALL resources | To stop costs |
| `terraform state list` | Shows managed resources | Debugging |
| `terraform output` | Shows output values | Get URLs, names, etc. |

#### Understanding Terraform State

Terraform keeps track of what it created in a **state file** (`terraform.tfstate`). This file:

- Maps your config to real Azure resources
- Is created locally by default
- Should NEVER be committed to Git (contains secrets)
- Can be stored remotely (Azure Storage) for team collaboration

Add to `.gitignore`:
```
# Terraform
*.tfstate
*.tfstate.*
.terraform/
terraform.tfvars    # Contains your secrets!
```

#### Azure Resources Needed

| Resource | Purpose | Estimated Cost |
|----------|---------|----------------|
| Resource Group | Container for all resources | Free |
| Azure Container Registry (ACR) | Store container images | ~$5/month (Basic tier) |
| Azure Blob Storage | Art images | ~$0.02/GB/month + transactions |
| Azure Database for MySQL (Flexible) | Managed database | ~$12/month (Burstable B1ms) OR $0 if self-hosted in K8s |
| Azure Kubernetes Service (AKS) | Container orchestration | Free control plane, pay for nodes |
| AKS Node Pool | Worker VMs | ~$30-60/month per node (B2s spot instances) |

#### Common Issues & Solutions

| Problem | Solution |
|---------|----------|
| "Provider not found" | Run `terraform init` |
| "Not authenticated" | Run `az login` |
| "Resource already exists" | Resource was created outside Terraform. Import it or delete manually |
| "Quota exceeded" | Request quota increase in Azure Portal or use smaller VM size |
| State file locked | Another terraform process is running. Wait or delete `.terraform.lock.hcl` |

#### Verifying Your Infrastructure

After `terraform apply` succeeds:

```bash
# Check resource group was created
az group show --name hark-portfolio-rg

# Check AKS cluster
az aks list --output table

# Get credentials for kubectl
az aks get-credentials --resource-group hark-portfolio-rg --name hark-aks

# Verify kubectl works
kubectl get nodes
```

---

### Phase 3: Kubernetes Manifests

**Goal**: Define how your app runs in Kubernetes.

#### Manifest Structure:

```
kubernetes/
├── namespace.yaml           # Isolate resources
├── secrets.yaml             # DB passwords, API keys (use sealed-secrets or external-secrets in production)
├── configmap.yaml           # Non-sensitive config
│
├── frontend/
│   ├── deployment.yaml      # Pod definition, replicas
│   ├── service.yaml         # Internal networking
│   └── ingress.yaml         # External access
│
├── backend/
│   ├── deployment.yaml
│   ├── service.yaml
│   └── ingress.yaml
│
└── mysql/                   # Optional: self-hosted MySQL
    ├── statefulset.yaml     # Persistent workload
    ├── service.yaml
    └── pvc.yaml             # Persistent volume claim
```

#### Key Kubernetes Concepts You'll Use:

| Concept | What It Does |
|---------|--------------|
| Deployment | Manages pod replicas, rolling updates |
| Service | Stable internal DNS/IP for pods |
| Ingress | Routes external traffic (HTTPS) to services |
| ConfigMap | Environment variables, config files |
| Secret | Sensitive data (passwords, keys) |
| PersistentVolumeClaim | Disk storage for MySQL |
| HorizontalPodAutoscaler | Auto-scale based on CPU/memory |

---

### Phase 4: CI/CD with GitHub Actions

**Goal**: Automate building, testing, and deploying.

#### Workflow Files:

```
.github/workflows/
├── ci.yaml              # Runs on every PR: lint, test, build
├── deploy-dev.yaml      # Manual trigger: deploy to AKS
└── destroy.yaml         # Manual trigger: tear down AKS
```

#### CI Pipeline (ci.yaml):

```yaml
on: [push, pull_request]
jobs:
  test-frontend:    # npm test
  test-backend:     # go test ./...
  build-images:     # Build container images (don't push on PR)
```

#### Deploy Pipeline (deploy-dev.yaml):

```yaml
on: workflow_dispatch    # Manual trigger only

jobs:
  deploy:
    steps:
      - Checkout code
      - Login to Azure (using service principal secret)
      - Login to ACR
      - Build and push container images
      - Deploy to AKS with kubectl apply
```

---

## Cost Management Strategy

### Daily/Interview Use Pattern

The biggest costs are **AKS node VMs** and **Azure Database for MySQL**.

#### Option A: Scale Node Pool to Zero (Recommended)

```bash
# Before interview: Scale up (takes 3-5 minutes)
az aks nodepool scale \
  --resource-group hark-portfolio-rg \
  --cluster-name hark-aks \
  --name agentpool \
  --node-count 1

# After interview: Scale down (immediate)
az aks nodepool scale \
  --resource-group hark-portfolio-rg \
  --cluster-name hark-aks \
  --name agentpool \
  --node-count 0
```

**Cost when scaled to 0**: ~$0.10/day (just AKS control plane, storage)
**Cost when running**: ~$1-3/day (B2s spot instance)

#### Option B: Full Terraform Destroy/Apply

```bash
# Destroy everything
terraform destroy -auto-approve

# Recreate for interview
terraform apply -auto-approve
```

**Pros**: True $0 when destroyed
**Cons**: Takes 10-15 minutes to recreate, may lose data if not backed up

#### Option C: Use Spot Instances (Do This Regardless)

In `terraform/aks.tf`:
```hcl
resource "azurerm_kubernetes_cluster_node_pool" "spot" {
  name                = "spot"
  kubernetes_cluster_id = azurerm_kubernetes_cluster.main.id
  vm_size             = "Standard_B2s"
  node_count          = 1
  priority            = "Spot"
  eviction_policy     = "Delete"
  spot_max_price      = 0.01  # Max $0.01/hour (~$7/month if always on)
}
```

**Savings**: 60-90% compared to regular VMs

### Cost Summary

| Scenario | Monthly Cost |
|----------|--------------|
| Only local development | $0 |
| AKS scaled to 0, ACR + Blob active | ~$5-7/month |
| AKS running 8 hours for interviews (4x/month) | ~$8-12/month |
| AKS running 24/7 (spot instance) | ~$20-40/month |

### Services That Can Be Turned Off

| Service | How to Stop | Restart Time |
|---------|-------------|--------------|
| AKS Node Pool | Scale to 0 nodes | 3-5 min |
| Azure MySQL Flexible | Stop server | 2-3 min |
| Entire Infrastructure | `terraform destroy` | 10-15 min |

### Services That Should Stay On (Cheap)

| Service | Why | Cost |
|---------|-----|------|
| ACR (Basic) | Stores your images | ~$5/month |
| Blob Storage | Your art images | ~$0.02/GB |
| AKS Control Plane | Free tier | $0 |

---

## Quick Reference Commands

### Local Development
```bash
cd deployment/docker
podman-compose up -d        # Start
podman-compose down         # Stop
podman-compose logs -f api  # View logs
```

### Azure CLI
```bash
az login                                    # Authenticate
az aks get-credentials -g hark-portfolio-rg -n hark-aks  # Get kubeconfig
az acr login -n harkportfolioacr            # Docker login to ACR
```

### Terraform
```bash
cd deployment/terraform
terraform init          # First time
terraform plan          # Preview changes
terraform apply         # Create/update resources
terraform destroy       # Delete everything
```

### Kubernetes
```bash
kubectl get pods                    # List running pods
kubectl get svc                     # List services
kubectl logs -f deployment/backend  # View logs
kubectl apply -f kubernetes/        # Apply all manifests
kubectl delete -f kubernetes/       # Remove all manifests
```

### Interview Day Checklist
```bash
# 30 minutes before interview:
cd deployment/terraform
az aks nodepool scale -g hark-portfolio-rg -c hark-aks -n agentpool --node-count 1
kubectl get pods -w  # Wait for pods to be Ready

# After interview:
az aks nodepool scale -g hark-portfolio-rg -c hark-aks -n agentpool --node-count 0
```

---

## Next Steps

1. [x] Initialize Go module and basic API structure
2. [x] Initialize Svelte + TypeScript project
3. [x] Create Dockerfiles for both services
4. [x] Create docker-compose.yml for local dev (Podman-compatible)
5. [ ] Set up Terraform configuration
6. [ ] Create Kubernetes manifests
7. [ ] Set up GitHub Actions workflows
8. [ ] Test full deployment to AKS

---

## Resources

- [Svelte Documentation](https://svelte.dev/docs)
- [Go Documentation](https://go.dev/doc/)
- [Podman Documentation](https://docs.podman.io/)
- [Podman Compose](https://github.com/containers/podman-compose)
- [Terraform Azure Provider](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs)
- [Kubernetes Documentation](https://kubernetes.io/docs/home/)
- [AKS Documentation](https://docs.microsoft.com/en-us/azure/aks/)
- [Azure Blob Storage Go SDK](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/storage/azblob)
