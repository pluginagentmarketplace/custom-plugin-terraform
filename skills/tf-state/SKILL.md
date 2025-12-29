---
name: tf-state
description: Terraform state - backends and management
sasmp_version: "1.3.0"
bonded_agent: tf-state
bond_type: PRIMARY_BOND
---

# Terraform State Skill

## Remote Backends

```hcl
# S3 Backend
terraform {
  backend "s3" {
    bucket         = "my-terraform-state"
    key            = "prod/terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
    dynamodb_table = "terraform-locks"
  }
}

# Azure Backend
terraform {
  backend "azurerm" {
    resource_group_name  = "tfstate-rg"
    storage_account_name = "tfstatestorage"
    container_name       = "tfstate"
    key                  = "prod.terraform.tfstate"
  }
}
```

## State Commands

```bash
# Show state
terraform state list
terraform state show aws_instance.web

# Move resource
terraform state mv aws_instance.old aws_instance.new

# Remove from state
terraform state rm aws_instance.web

# Import existing resource
terraform import aws_instance.web i-1234567890

# Pull/Push state
terraform state pull > state.json
terraform state push state.json
```

## Quick Reference

| Backend | Locking |
|---------|---------|
| S3 | DynamoDB |
| Azure | Native |
| GCS | Native |
| Consul | Native |

## Related
- tf-modules - Module state
- tf-state agent
