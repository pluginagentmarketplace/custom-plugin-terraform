---
name: state-management
description: Master Terraform state management - local and remote state, locking, encryption, backup, recovery, and multi-environment strategies.
---

# State Management Mastery

## Quick Start

```bash
# Initialize with remote backend
terraform init

# Configure remote state (S3)
terraform {
  backend "s3" {
    bucket         = "terraform-state"
    key            = "prod/terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
    dynamodb_table = "terraform-locks"
  }
}

# State commands
terraform state list
terraform state show resource.id
terraform state mv old.id new.id
terraform state rm resource.id
terraform refresh
```

## State Concepts

### Local vs Remote
- Local: Development only
- Remote: Team and production
- State locking
- State encryption
- Backup strategies

### Backend Options
- S3 with DynamoDB locking
- Terraform Cloud
- Consul
- Azure Blob Storage
- Google Cloud Storage

### State File Structure
- Resources and attributes
- Metadata and lineage
- Outputs and values
- Dependencies
- Version tracking

### Operations
- Reading state
- Modifying state
- Importing resources
- Backing up/recovery
- Migration

### Best Practices
- Always use remote state
- Enable encryption
- Implement locking
- Regular backups
- Audit access logs
- Never commit .tfstate files
- Use workspaces for separation

## Security Considerations

- Access control (IAM, RBAC)
- Encryption at rest and transit
- State file permissions
- Audit logging
- Secrets management
