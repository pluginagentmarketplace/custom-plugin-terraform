---
name: terraform-state
description: Master Terraform state management including remote backends, locking, workspaces, and state operations
sasmp_version: "1.3.0"
version: "2.0.0"
bonded_agent: 03-terraform-state
bond_type: PRIMARY_BOND
---

# Terraform State Skill

Comprehensive guide to state management, remote backends, and state manipulation operations.

## Backend Configuration

### AWS S3 Backend
```hcl
terraform {
  backend "s3" {
    bucket         = "company-terraform-state"
    key            = "prod/terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
    dynamodb_table = "terraform-locks"

    # Cross-account access
    role_arn = "arn:aws:iam::ACCOUNT:role/TerraformAccess"
  }
}
```

### Azure Storage Backend
```hcl
terraform {
  backend "azurerm" {
    resource_group_name  = "terraform-state-rg"
    storage_account_name = "tfstatecompany"
    container_name       = "tfstate"
    key                  = "prod.terraform.tfstate"
    use_azuread_auth     = true
  }
}
```

### GCP Cloud Storage Backend
```hcl
terraform {
  backend "gcs" {
    bucket = "company-terraform-state"
    prefix = "terraform/state"
  }
}
```

### Terraform Cloud
```hcl
terraform {
  cloud {
    organization = "my-org"

    workspaces {
      tags = ["app:my-app", "env:prod"]
    }
  }
}
```

## State Locking

### DynamoDB Lock Table
```hcl
resource "aws_dynamodb_table" "terraform_locks" {
  name         = "terraform-locks"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }

  tags = {
    Purpose = "Terraform state locking"
  }
}
```

## Workspace Strategies

### Directory-Based (Recommended)
```
infrastructure/
├── environments/
│   ├── dev/
│   │   ├── main.tf
│   │   ├── backend.tf
│   │   └── terraform.tfvars
│   ├── staging/
│   └── prod/
└── modules/
```

### CLI Workspaces
```bash
# Create and switch
terraform workspace new prod
terraform workspace select prod
terraform workspace list

# Use in config
locals {
  env_config = {
    dev  = { instance_type = "t3.micro" }
    prod = { instance_type = "t3.large" }
  }
  config = local.env_config[terraform.workspace]
}
```

## State Operations

### Import Resources
```bash
# Basic import
terraform import aws_instance.web i-1234567890

# Import into module
terraform import module.vpc.aws_vpc.main vpc-abc123

# Import with for_each
terraform import 'aws_instance.web["app-1"]' i-1234567890

# Generate config (Terraform 1.5+)
terraform plan -generate-config-out=generated.tf
```

### Move Resources
```bash
# Rename resource
terraform state mv aws_instance.old aws_instance.new

# Move to module
terraform state mv aws_instance.web module.compute.aws_instance.web

# Move between states
terraform state mv -state-out=other.tfstate aws_instance.web aws_instance.web
```

### Remove from State
```bash
# Remove without destroying
terraform state rm aws_instance.legacy

# Remove entire module
terraform state rm module.deprecated
```

### State Commands
```bash
# List resources
terraform state list

# Show resource details
terraform state show aws_instance.web

# Pull remote state
terraform state pull > state.json

# Push state (DANGEROUS)
terraform state push state.json
```

## Drift Detection

```bash
# Refresh and show drift
terraform apply -refresh-only

# Detailed plan with refresh
terraform plan -refresh=true

# Check for changes only
terraform plan -detailed-exitcode
# Exit codes: 0=no changes, 1=error, 2=changes
```

## Troubleshooting

### State Lock Issues
```bash
# Force unlock (only if confirmed stale!)
terraform force-unlock LOCK_ID

# Check DynamoDB for lock
aws dynamodb get-item \
  --table-name terraform-locks \
  --key '{"LockID":{"S":"path/to/state"}}'
```

### Common Errors

| Error | Cause | Solution |
|-------|-------|----------|
| `Error acquiring lock` | Stale lock | Wait or force-unlock |
| `Backend config changed` | Config modified | `terraform init -reconfigure` |
| `State out of sync` | Manual changes | `terraform apply -refresh-only` |
| `Resource already exists` | Import needed | `terraform import` |

### Recovery Procedures
```bash
# Backup current state
terraform state pull > backup.json

# Verify backup
cat backup.json | jq '.resources | length'

# Restore if needed
terraform state push backup.json
```

## Best Practices

1. **Always use remote state** - Never commit .tfstate
2. **Enable state locking** - Prevent concurrent modifications
3. **Enable encryption** - Protect sensitive data
4. **Version state bucket** - S3 versioning for recovery
5. **Separate state per environment** - Isolation and safety

## Usage

```python
Skill("terraform-state")
```

## Related

- **Agent**: 03-terraform-state (PRIMARY_BOND)
- **Skill**: terraform-workspace (SECONDARY_BOND)
