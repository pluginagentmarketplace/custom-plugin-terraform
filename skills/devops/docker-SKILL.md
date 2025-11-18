---
name: docker-containers
description: Master Docker containerization, image optimization, Docker Compose, and container orchestration. Build, deploy, and manage containerized applications at scale.
---

# Docker Containers

## Quick Start

```dockerfile
# Multi-stage Dockerfile for optimization
FROM node:18 AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci --only=production

FROM node:18-alpine
WORKDIR /app
COPY --from=builder /app/node_modules ./node_modules
COPY . .
EXPOSE 3000
CMD ["node", "index.js"]
```

## Core Competencies

### Docker Basics
- Images and containers
- Dockerfile creation and optimization
- Layers and caching strategies
- Multi-stage builds for efficiency

### Image Management
- Docker registries (Docker Hub, ECR, Harbor)
- Image tagging and versioning
- Image scanning for vulnerabilities
- Private registry management
- Image optimization techniques

### Docker Compose
- Multi-container applications
- Networking between containers
- Volume mounting for persistence
- Environment variable management
- Override files for different environments

### Container Networking
- Bridge networks
- Host and overlay networks
- Service discovery
- Port mapping

### Container Security
- User permissions inside containers
- Secret management
- Resource limits (CPU, memory)
- Network policies
- Image signing

### Best Practices
- Small base images (Alpine)
- Layer reduction
- Health checks
- Logging strategies
- Container immutability

## Real-World Workflows

- Building production images
- Local development with Compose
- Container registry operations
- Debugging running containers
- Performance optimization

## Essential Tools

- **Docker CLI** - Container management
- **Docker Compose** - Multi-container orchestration
- **Docker Desktop** - Local development
- **Trivy** - Container scanning
- **Dive** - Image optimization analysis
