---
name: DevOps & Infrastructure Mastery
description: Master Docker containerization, Kubernetes orchestration, CI/CD pipelines, Infrastructure as Code (Terraform), and cloud platforms. Learn deployment automation, monitoring, scaling, and build reliable, scalable infrastructure for production systems.
sasmp_version: "1.3.0"
bonded_agent: 01-programming-fundamentals
bond_type: PRIMARY_BOND
---

# 🚀 DevOps & Infrastructure Mastery

**Build, deploy, and operate production infrastructure at scale.**

Master containerization, orchestration, automation, and deployment to build reliable, scalable infrastructure.

## Quick Start: Docker & Kubernetes

### Docker Basics
```dockerfile
# Dockerfile for Node.js application
FROM node:20-alpine

WORKDIR /app

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm ci --only=production

# Copy application code
COPY . .

# Expose port
EXPOSE 3000

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD node healthcheck.js

# Run application
CMD ["node", "server.js"]
```

### Docker Compose for Local Development
```yaml
version: '3.8'

services:
  app:
    build: .
    ports:
      - "3000:3000"
    environment:
      - NODE_ENV=development
      - DATABASE_URL=postgres://user:password@db:5432/app
    depends_on:
      db:
        condition: service_healthy
    volumes:
      - .:/app
      - /app/node_modules
    networks:
      - app-network

  db:
    image: postgres:15-alpine
    environment:
      - POSTGRES_USER=user
      - POSTGRES_PASSWORD=password
      - POSTGRES_DB=app
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U user"]
      interval: 10s
      timeout: 5s
      retries: 5
    networks:
      - app-network

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    networks:
      - app-network

volumes:
  postgres_data:

networks:
  app-network:
    driver: bridge
```

### Kubernetes Deployment
```yaml
# deployment.yaml - Kubernetes deployment
apiVersion: apps/v1
kind: Deployment
metadata:
  name: api-deployment
  labels:
    app: api
spec:
  replicas: 3
  selector:
    matchLabels:
      app: api
  template:
    metadata:
      labels:
        app: api
    spec:
      containers:
      - name: api
        image: myregistry.azurecr.io/api:1.0.0
        ports:
        - containerPort: 3000
        env:
        - name: NODE_ENV
          value: "production"
        - name: DATABASE_URL
          valueFrom:
            secretKeyRef:
              name: db-secret
              key: connection-string
        resources:
          requests:
            memory: "256Mi"
            cpu: "250m"
          limits:
            memory: "512Mi"
            cpu: "500m"
        livenessProbe:
          httpGet:
            path: /health
            port: 3000
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /ready
            port: 3000
          initialDelaySeconds: 5
          periodSeconds: 5

---
# service.yaml - Kubernetes service
apiVersion: v1
kind: Service
metadata:
  name: api-service
spec:
  selector:
    app: api
  type: LoadBalancer
  ports:
  - protocol: TCP
    port: 80
    targetPort: 3000

---
# autoscaling.yaml - Horizontal Pod Autoscaling
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: api-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: api-deployment
  minReplicas: 2
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 70
  - type: Resource
    resource:
      name: memory
      target:
        type: Utilization
        averageUtilization: 80
```

### Infrastructure as Code with Terraform
```hcl
# variables.tf
variable "environment" {
  type        = string
  description = "Environment name"
  default     = "production"
}

variable "region" {
  type        = string
  description = "AWS region"
  default     = "us-east-1"
}

variable "app_name" {
  type        = string
  description = "Application name"
}

variable "instance_count" {
  type        = number
  description = "Number of EC2 instances"
  default     = 3
}

# main.tf
terraform {
  required_version = ">= 1.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }

  backend "s3" {
    bucket         = "terraform-state"
    key            = "prod/terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
    dynamodb_table = "terraform-lock"
  }
}

provider "aws" {
  region = var.region
}

# VPC
resource "aws_vpc" "main" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name        = "${var.app_name}-vpc"
    Environment = var.environment
  }
}

# Security Group
resource "aws_security_group" "app" {
  name        = "${var.app_name}-sg"
  description = "Security group for ${var.app_name}"
  vpc_id      = aws_vpc.main.id

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# EC2 Instances
resource "aws_instance" "app" {
  count           = var.instance_count
  ami             = data.aws_ami.ubuntu.id
  instance_type   = "t3.medium"
  security_groups = [aws_security_group.app.id]

  user_data = base64encode(file("${path.module}/scripts/init.sh"))

  tags = {
    Name        = "${var.app_name}-instance-${count.index + 1}"
    Environment = var.environment
  }
}

# outputs.tf
output "instance_ips" {
  value       = aws_instance.app[*].public_ip
  description = "Public IPs of instances"
}

output "security_group_id" {
  value       = aws_security_group.app.id
  description = "Security group ID"
}
```

## CI/CD Pipelines

### GitHub Actions Workflow
```yaml
# .github/workflows/deploy.yml
name: Deploy Application

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    runs-on: ubuntu-latest
    services:
      postgres:
        image: postgres:15
        env:
          POSTGRES_PASSWORD: postgres
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
        ports:
          - 5432:5432

    steps:
    - uses: actions/checkout@v3

    - name: Setup Node.js
      uses: actions/setup-node@v3
      with:
        node-version: '20'
        cache: 'npm'

    - name: Install dependencies
      run: npm ci

    - name: Run linter
      run: npm run lint

    - name: Run tests
      run: npm test
      env:
        DATABASE_URL: postgres://postgres:postgres@localhost:5432/test_db

    - name: Build application
      run: npm run build

  deploy:
    needs: test
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'

    steps:
    - uses: actions/checkout@v3

    - name: Login to Docker Registry
      uses: docker/login-action@v2
      with:
        registry: ${{ secrets.REGISTRY_URL }}
        username: ${{ secrets.REGISTRY_USERNAME }}
        password: ${{ secrets.REGISTRY_PASSWORD }}

    - name: Build and push Docker image
      uses: docker/build-push-action@v4
      with:
        context: .
        push: true
        tags: |
          ${{ secrets.REGISTRY_URL }}/app:${{ github.sha }}
          ${{ secrets.REGISTRY_URL }}/app:latest

    - name: Deploy to Kubernetes
      run: |
        curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
        chmod +x kubectl
        ./kubectl set image deployment/api-deployment \
          api=${{ secrets.REGISTRY_URL }}/app:${{ github.sha }} \
          -n production
```

## Monitoring & Logging

### Prometheus Configuration
```yaml
# prometheus.yml
global:
  scrape_interval: 15s
  evaluation_interval: 15s

alerting:
  alertmanagers:
    - static_configs:
        - targets:
            - alertmanager:9093

rule_files:
  - "alert_rules.yml"

scrape_configs:
  - job_name: 'prometheus'
    static_configs:
      - targets: ['localhost:9090']

  - job_name: 'app'
    static_configs:
      - targets: ['app:3000']

  - job_name: 'kubernetes'
    kubernetes_sd_configs:
      - role: pod
```

### ELK Stack for Logging
```yaml
# docker-compose.yml for ELK
version: '3.8'

services:
  elasticsearch:
    image: docker.elastic.co/elasticsearch/elasticsearch:8.0.0
    environment:
      - discovery.type=single-node
      - xpack.security.enabled=false
    ports:
      - "9200:9200"

  kibana:
    image: docker.elastic.co/kibana/kibana:8.0.0
    ports:
      - "5601:5601"
    depends_on:
      - elasticsearch

  logstash:
    image: docker.elastic.co/logstash/logstash:8.0.0
    ports:
      - "5000:5000"
    volumes:
      - ./logstash.conf:/usr/share/logstash/pipeline/logstash.conf
    depends_on:
      - elasticsearch
```

## Learning Path

### Phase 1: Foundation (Weeks 1-3)
- Docker basics and images
- Basic CI/CD concepts
- Cloud basics
- Simple monitoring
- **Projects**:
  - Containerize application
  - Create simple CI/CD pipeline
  - Deploy to cloud

### Phase 2: Intermediate (Weeks 4-9)
- Kubernetes fundamentals
- Advanced CI/CD
- Infrastructure as Code
- Multi-environment setup
- **Projects**:
  - Deploy to Kubernetes
  - Automate infrastructure
  - Multi-environment deployment

### Phase 3: Advanced (Weeks 10-18)
- Advanced Kubernetes
- Multi-cloud strategies
- Advanced monitoring
- Disaster recovery
- **Projects**:
  - Production Kubernetes setup
  - Automated deployments
  - Disaster recovery procedures

## 16+ Production Projects

### Docker Projects
1. **Containerize Monolith** - Single container
2. **Multi-container App** - App + DB + Cache
3. **Private Registry** - Container image registry
4. **Container Optimization** - Reduce image size

### Kubernetes Projects
5. **Simple Deployment** - Deploy single app
6. **Microservices** - Deploy multiple services
7. **Rolling Updates** - Zero-downtime deployments
8. **Auto-scaling** - HPA configuration

### CI/CD Projects
9. **GitHub Actions** - Automated testing and deployment
10. **GitLab CI** - Complex pipeline with multiple stages
11. **Multi-stage Pipeline** - Build, test, deploy
12. **Canary Deployments** - Gradual rollouts

### Infrastructure Projects
13. **Terraform AWS** - Infrastructure as code
14. **VPC & Networking** - Network setup
15. **Disaster Recovery** - Backup and restore
16. **Multi-region Setup** - Global infrastructure

## Best Practices Checklist

### Docker
- ✅ Use specific base image versions
- ✅ Minimize layers and image size
- ✅ Use health checks
- ✅ Don't run as root
- ✅ Use environment variables

### Kubernetes
- ✅ Set resource limits
- ✅ Use liveness/readiness probes
- ✅ Implement pod disruption budgets
- ✅ Use network policies
- ✅ Enable pod security policies

### CI/CD
- ✅ Automate all quality checks
- ✅ Run tests before deployment
- ✅ Use semantic versioning
- ✅ Automate deployments
- ✅ Monitor deployments

### Infrastructure
- ✅ Use IaC for all infrastructure
- ✅ Version control everything
- ✅ Automate provisioning
- ✅ Monitor infrastructure
- ✅ Test disaster recovery

## 🏆 Success Milestones

- [ ] Containerize application with Docker
- [ ] Create multi-stage CI/CD pipeline
- [ ] Deploy application to Kubernetes
- [ ] Implement horizontal auto-scaling
- [ ] Set up infrastructure with Terraform
- [ ] Implement monitoring and alerting
- [ ] Achieve zero-downtime deployments
- [ ] Set up disaster recovery procedures

---

**Ready to master DevOps? Start by containerizing your application!**
