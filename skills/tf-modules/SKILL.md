---
name: tf-modules
description: Terraform modules - creation and usage
sasmp_version: "1.3.0"
bonded_agent: tf-modules
bond_type: PRIMARY_BOND
---

# Terraform Modules Skill

## Module Structure

```
modules/
└── vpc/
    ├── main.tf
    ├── variables.tf
    ├── outputs.tf
    └── README.md
```

## Using Modules

```hcl
# Local module
module "vpc" {
  source = "./modules/vpc"
  
  cidr_block  = "10.0.0.0/16"
  environment = var.environment
}

# Registry module
module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "5.0.0"
  
  name = "my-vpc"
  cidr = "10.0.0.0/16"
}

# Git module
module "vpc" {
  source = "git::https://github.com/org/modules.git//vpc?ref=v1.0.0"
}

# Access outputs
resource "aws_instance" "web" {
  subnet_id = module.vpc.public_subnet_ids[0]
}
```

## Module Development

```hcl
# variables.tf
variable "environment" {
  type        = string
  description = "Environment name"
}

# outputs.tf
output "vpc_id" {
  value       = aws_vpc.main.id
  description = "VPC ID"
}
```

## Quick Reference

| Source | Example |
|--------|---------|
| Local | `./modules/vpc` |
| Registry | `hashicorp/consul/aws` |
| GitHub | `github.com/org/repo` |
| Git | `git::https://...` |

## Related
- tf-variables - Module inputs
- tf-outputs - Module outputs
