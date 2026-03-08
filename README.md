# Hark Art Portfolio

A scalable art portfolio and sales platform built with modern technologies and cloud-native deployment practices.

## Tech Stack

| Layer | Technology | Purpose |
|-------|------------|---------|
| Frontend | Svelte + TypeScript | Lightweight, compiled UI framework with type safety |
| Backend | Go | Fast, simple API server |
| Database | MySQL | Relational data (users, orders, art metadata) |
| Image Storage | Azure Blob Storage | Art images and media files |
| Containerization | Docker | Consistent environments |
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
│   ├── deployments/
│   │   ├── docker/        # Dockerfiles and docker-compose
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
| Local Dev | Docker Compose | Daily development | Free |
| Interview Demo | AKS + Azure services | Demonstrate DevOps skills | ~$5-15/day when running |

---

## Detailed Deployment Plan

### Phase 1: Local Development Setup

**Goal**: Get the app running locally with Docker Compose.

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
   docker-compose up -d          # Start all services
   docker-compose logs -f        # Watch logs
   docker-compose down           # Stop all services
   ```

**Cost**: Free

---

### Phase 2: Azure Infrastructure with Terraform

**Goal**: Define all Azure resources as code for reproducible deployments.

#### Azure Resources Needed:

| Resource | Purpose | Estimated Cost |
|----------|---------|----------------|
| Resource Group | Container for all resources | Free |
| Azure Container Registry (ACR) | Store Docker images | ~$5/month (Basic tier) |
| Azure Blob Storage | Art images | ~$0.02/GB/month + transactions |
| Azure Database for MySQL (Flexible) | Managed database | ~$12/month (Burstable B1ms) OR $0 if self-hosted in K8s |
| Azure Kubernetes Service (AKS) | Container orchestration | Free control plane, pay for nodes |
| AKS Node Pool | Worker VMs | ~$30-60/month per node (B2s spot instances) |

#### Terraform Structure:

```
terraform/
├── main.tf              # Provider config, resource group
├── variables.tf         # Input variables
├── outputs.tf           # Output values (URLs, connection strings)
├── acr.tf               # Container registry
├── storage.tf           # Blob storage account
├── aks.tf               # Kubernetes cluster
├── mysql.tf             # Database (optional, can run in K8s)
└── terraform.tfvars     # Your specific values (gitignored)
```

#### Key Terraform Commands:

```bash
# First time setup
terraform init

# Preview what will be created
terraform plan

# Create infrastructure (STARTS COSTING MONEY)
terraform apply

# Destroy everything (STOPS COSTS)
terraform destroy
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
  build-images:     # Build Docker images (don't push on PR)
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
      - Build and push Docker images
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
cd deployments/docker
docker-compose up -d        # Start
docker-compose down         # Stop
docker-compose logs -f api  # View logs
```

### Azure CLI
```bash
az login                                    # Authenticate
az aks get-credentials -g hark-portfolio-rg -n hark-aks  # Get kubeconfig
az acr login -n harkportfolioacr            # Docker login to ACR
```

### Terraform
```bash
cd deployments/terraform
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
cd deployments/terraform
az aks nodepool scale -g hark-portfolio-rg -c hark-aks -n agentpool --node-count 1
kubectl get pods -w  # Wait for pods to be Ready

# After interview:
az aks nodepool scale -g hark-portfolio-rg -c hark-aks -n agentpool --node-count 0
```

---

## Next Steps

1. [ ] Initialize Go module and basic API structure
2. [ ] Initialize Svelte + TypeScript project
3. [ ] Create Dockerfiles for both services
4. [ ] Create docker-compose.yml for local dev
5. [ ] Set up Terraform configuration
6. [ ] Create Kubernetes manifests
7. [ ] Set up GitHub Actions workflows
8. [ ] Test full deployment to AKS

---

## Resources

- [Svelte Documentation](https://svelte.dev/docs)
- [Go Documentation](https://go.dev/doc/)
- [Docker Compose Reference](https://docs.docker.com/compose/)
- [Terraform Azure Provider](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs)
- [Kubernetes Documentation](https://kubernetes.io/docs/home/)
- [AKS Documentation](https://docs.microsoft.com/en-us/azure/aks/)
- [Azure Blob Storage Go SDK](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/storage/azblob)
