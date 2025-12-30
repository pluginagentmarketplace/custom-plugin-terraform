---
name: terraform-modules
description: Create reusable, composable Terraform modules with proper versioning and registry integration
sasmp_version: "1.3.0"
version: "2.0.0"
bonded_agent: 02-terraform-modules
bond_type: PRIMARY_BOND
---

# Terraform Modules Skill

Design and build production-ready, reusable Terraform modules following industry best practices.

## Module Structure

```
my-module/
├── main.tf           # Primary resources
├── variables.tf      # Input declarations
├── outputs.tf        # Output declarations
├── versions.tf       # Provider requirements
├── locals.tf         # Computed values
├── README.md         # Documentation
├── examples/
│   ├── basic/
│   └── complete/
└── tests/
    └── module_test.go
```

## Quick Start Template

### versions.tf
```hcl
terraform {
  required_version = ">= 1.5.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.0, < 6.0"
    }
  }
}
```

### variables.tf
```hcl
variable "name" {
  type        = string
  description = "Resource name prefix"

  validation {
    condition     = length(var.name) <= 32
    error_message = "Name must be 32 characters or less."
  }
}

variable "environment" {
  type        = string
  description = "Deployment environment"

  validation {
    condition     = contains(["dev", "staging", "prod"], var.environment)
    error_message = "Must be dev, staging, or prod."
  }
}

variable "tags" {
  type        = map(string)
  description = "Resource tags"
  default     = {}
}
```

### main.tf
```hcl
locals {
  common_tags = merge(var.tags, {
    Module      = "my-module"
    Environment = var.environment
  })
}

resource "aws_resource" "main" {
  name = var.name
  tags = local.common_tags
}
```

### outputs.tf
```hcl
output "id" {
  description = "Resource ID"
  value       = aws_resource.main.id
}

output "arn" {
  description = "Resource ARN"
  value       = aws_resource.main.arn
}
```

## Input Patterns

### Required vs Optional
```hcl
# Required - no default
variable "vpc_id" {
  type        = string
  description = "VPC ID (required)"
}

# Optional with default
variable "instance_type" {
  type        = string
  description = "EC2 instance type"
  default     = "t3.micro"
}

# Optional nullable
variable "kms_key_arn" {
  type        = string
  description = "KMS key ARN (null = use default)"
  default     = null
}
```

### Complex Objects
```hcl
variable "scaling_config" {
  type = object({
    min_size         = number
    max_size         = number
    desired_capacity = optional(number)
    metrics          = optional(list(string), ["CPUUtilization"])
  })

  default = {
    min_size = 1
    max_size = 3
  }

  validation {
    condition     = var.scaling_config.min_size <= var.scaling_config.max_size
    error_message = "min_size must be <= max_size."
  }
}
```

## Output Patterns

### Conditional Outputs
```hcl
output "nat_gateway_ips" {
  description = "NAT Gateway IPs (empty if not created)"
  value       = var.enable_nat ? aws_eip.nat[*].public_ip : []
}
```

### Structured Outputs
```hcl
output "cluster_config" {
  description = "Cluster configuration for kubectl"
  value = {
    endpoint       = aws_eks_cluster.main.endpoint
    ca_certificate = aws_eks_cluster.main.certificate_authority[0].data
    name           = aws_eks_cluster.main.name
  }
}
```

### Sensitive Outputs
```hcl
output "credentials" {
  description = "Database credentials"
  value = {
    username = aws_db_instance.main.username
    password = random_password.db.result
  }
  sensitive = true
}
```

## Module Composition

### Root Module Calling Child Modules
```hcl
module "vpc" {
  source = "./modules/vpc"

  name        = var.project_name
  cidr_block  = var.vpc_cidr
  environment = var.environment
}

module "eks" {
  source = "./modules/eks"

  cluster_name = var.project_name
  vpc_id       = module.vpc.vpc_id
  subnet_ids   = module.vpc.private_subnet_ids

  depends_on = [module.vpc]
}
```

### Using Registry Modules
```hcl
module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "5.1.2"  # Pin exact version

  name = var.project_name
  cidr = var.vpc_cidr

  azs             = data.aws_availability_zones.available.names
  private_subnets = var.private_subnet_cidrs
  public_subnets  = var.public_subnet_cidrs
}
```

### for_each with Modules
```hcl
module "services" {
  source   = "./modules/ecs-service"
  for_each = var.services

  name           = each.key
  container_port = each.value.port
  cpu            = each.value.cpu
  memory         = each.value.memory

  cluster_id = aws_ecs_cluster.main.id
}
```

## Versioning

### Semantic Versioning
```hcl
# Major: Breaking changes
# Minor: New features, backward compatible
# Patch: Bug fixes

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "~> 5.0"    # >= 5.0.0, < 6.0.0
}

module "internal" {
  source = "git::https://github.com/org/modules.git//vpc?ref=v2.3.0"
}
```

### Version Constraints
```hcl
version = "5.1.2"      # Exact version
version = ">= 5.0"     # Minimum version
version = "~> 5.0"     # Minor version range
version = ">= 5.0, < 6.0"  # Explicit range
```

## Troubleshooting

| Error | Cause | Solution |
|-------|-------|----------|
| `Module not found` | Wrong path/URL | Verify source path |
| `Unsupported argument` | Version mismatch | Check changelog |
| `Cycle detected` | Circular references | Restructure dependencies |
| `Count/for_each conflict` | Both used together | Use only one |

### Debug Commands
```bash
# Initialize modules
terraform init -upgrade

# Show module tree
terraform providers

# Validate module inputs
terraform validate
```

## Usage

```python
Skill("terraform-modules")
```

## Related

- **Agent**: 02-terraform-modules (PRIMARY_BOND)
