# Portfolio

A full-stack portfolio application deployed on Azure Kubernetes Service.

## Tech Stack

**Frontend:** SvelteKit 2, TypeScript, Vite
**Backend:** Go 1.23, Gin
**Database:** MySQL 8.0
**Infrastructure:** Terraform, Kubernetes, Azure (AKS, ACR, Blob Storage)
**CI/CD:** GitHub Actions

## Project Structure

```
├── frontend/          # SvelteKit app
├── backend/           # Go API
└── deployment/
    ├── docker/        # Docker Compose for local dev
    ├── kubernetes/    # K8s manifests
    └── terraform/     # Infrastructure as code
```

## Local Development

**Prerequisites:** Docker, Docker Compose

```bash
cd deployment/docker
docker compose up --build
```

- Frontend: http://localhost:3000
- Backend: http://localhost:8080

## Deployment

### Infrastructure

```bash
cd deployment/terraform
terraform init
terraform apply
```

### Kubernetes

```bash
az aks get-credentials --resource-group Portfolio --name hark-portfolio-aks
kubectl apply -f deployment/kubernetes/
```

### CI/CD

Pushing to `main` triggers automatic deployment via GitHub Actions.

## Environment Variables

**Backend:**
| Variable | Description |
|----------|-------------|
| DB_HOST | MySQL hostname |
| DB_PORT | MySQL port |
| DB_USER | Database user |
| DB_PASSWORD | Database password |
| DB_NAME | Database name |
