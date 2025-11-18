---
name: terraform-scaling
description: Scale Terraform for large organizations - multi-cloud, governance, cost optimization, and enterprise-grade deployment patterns.
---

# Terraform Scaling at Enterprise

## Quick Start - Multi-Environment with Workspaces

```hcl
terraform {
  cloud {
    organization = "my-org"
    workspaces {
      name = "${var.environment}-${var.project}"
    }
  }
}

# Environment-specific configuration
variable "environment" {
  type = string
  validation {
    condition     = contains(["dev", "staging", "prod"], var.environment)
    error_message = "Invalid environment."
  }
}

# Use variable for configuration
resource "aws_instance" "app" {
  instance_type = var.environment == "prod" ? "t3.large" : "t3.micro"
  count         = var.environment == "prod" ? 3 : 1

  tags = {
    Environment = var.environment
  }
}
```

## Enterprise Patterns

### Multi-Cloud Strategy
- AWS, GCP, Azure providers
- Cross-cloud resources
- Provider abstraction
- Migration planning

### Governance & Compliance
- Sentinel policies
- RBAC implementation
- Audit logging
- Compliance standards
- Security groups

### Cost Optimization
- Resource right-sizing
- Spot instances
- Reserved capacity
- Lifecycle management
- Cost estimation
- FinOps integration

### Team Organization
- Workspace separation
- Module versioning
- Monorepo vs polyrepo
- CI/CD pipelines
- Code review process

### Large-Scale Operations
- Dependency management
- Parallel deployments
- Progressive rollouts
- Canary releases
- Blue-green deployments
- Disaster recovery
- High availability

## Scaling Challenges

- State file size
- Plan execution time
- Dependency resolution
- Resource limits
- API rate limits
- Lock contention
- Cost tracking
- Team coordination

## Solutions

- Module composition
- Workspaces
- Remote state
- Terraform Cloud
- Cost monitoring
- Automation
- Documentation
- Training
