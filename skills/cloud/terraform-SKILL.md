---
name: terraform-iac
description: Master Terraform infrastructure-as-code for declarative infrastructure management. Learn HCL, state management, modules, and deploying to AWS, GCP, Azure, and other providers.
---

# Terraform Infrastructure as Code

## Quick Start

```hcl
# main.tf - Define infrastructure
terraform {
  required_version = ">= 1.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

resource "aws_instance" "app_server" {
  ami           = data.aws_ami.ubuntu.id
  instance_type = var.instance_type

  tags = {
    Name = var.server_name
  }
}

output "instance_ip" {
  value = aws_instance.app_server.public_ip
}
```

## Core Competencies

### HCL Language
- Blocks, resources, and variables
- Expressions and interpolation
- Locals for reusability
- Functions and data sources
- Conditional logic and loops

### State Management
- Remote state (S3 backend)
- State locking with DynamoDB
- State encryption
- State backup and recovery
- Workspace management

### Modules & Reusability
- Module creation and structure
- Module registry (public and private)
- Composing modules
- Best practices for organization
- Versioning strategies

### Multi-Environment Deployments
- Workspace-based separation
- Directory-based separation
- Variable files (.tfvars)
- Environment-specific configurations

### Multi-Cloud Support
- AWS provider configuration
- Google Cloud provider
- Azure provider
- Multi-cloud architecture patterns

### Advanced Topics
- Custom providers
- Policy as code (Sentinel)
- Terraform Cloud/Enterprise
- GitOps workflows

## Workflow Essentials

- **terraform init** - Initialize workspace
- **terraform plan** - Show changes
- **terraform apply** - Deploy infrastructure
- **terraform destroy** - Remove resources
- **terraform import** - Manage existing resources
- **terraform refresh** - Update state

## Best Practices

- Use version control (Git)
- Implement code review process
- Automate testing (Terratest)
- Document infrastructure
- Use workspaces or directories for environments
- Implement cost controls

## Project Scenarios

- Multi-tier application on AWS
- Multi-region disaster recovery
- Kubernetes cluster provisioning
- Complete infrastructure stack
- Infrastructure testing

## Enterprise Patterns

- Module hierarchy
- Central registry
- CI/CD integration
- Team collaboration
- Cost management
