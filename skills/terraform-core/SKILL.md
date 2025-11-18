---
name: terraform-core-fundamentals
description: Master Terraform basics - HCL syntax, resources, providers, state, and the complete terraform workflow. Essential foundation for all Terraform work.
---

# Terraform Core Fundamentals

## Quick Start

```hcl
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

# Define a variable
variable "aws_region" {
  type    = string
  default = "us-east-1"
}

# Create a resource
resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"

  tags = {
    Name = "main-vpc"
  }
}

# Output the resource
output "vpc_id" {
  value = aws_vpc.main.id
}
```

## Core Concepts

### HCL Structure
- Blocks: terraform, provider, resource, variable, output
- Arguments and attributes
- Comments (#)
- File organization

### Providers & Resources
- Provider configuration
- Resource types
- Data sources
- State tracking

### Variables & Outputs
- Input variables with validation
- Local values
- Output values for sharing
- Sensitive data handling

### Terraform Workflow
1. **Init**: Initialize workspace
2. **Validate**: Check configuration
3. **Plan**: Preview changes
4. **Apply**: Deploy resources
5. **Destroy**: Remove resources

## Essential Commands

```bash
terraform init        # Initialize
terraform plan       # Preview changes
terraform apply      # Deploy
terraform destroy    # Remove
terraform import     # Import existing
terraform state      # Manage state
terraform fmt        # Format code
terraform validate   # Validate syntax
```

## Best Practices

- Always use version constraints
- DRY with modules
- Organize with workspaces
- Document with comments
- Use variable validation
- Manage secrets securely
