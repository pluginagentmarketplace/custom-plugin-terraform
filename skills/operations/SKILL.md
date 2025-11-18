---
name: terraform-operations
description: Master Terraform operations at scale - deployments, state management, multi-environment setups, and production patterns.
---

# Terraform Operations Mastery

## Quick Start - Safe Deployment

```bash
# Plan before apply
terraform plan -out=tfplan
terraform show tfplan  # Review changes

# Apply with confidence
terraform apply tfplan

# Track changes with VCS
git add terraform.tfstate
git commit -m "Infrastructure update"

# Monitor state
terraform state list
terraform state show resource
```

## Operations Essentials

### Deployment Workflow
- Planning and review
- Change validation
- Apply strategies
- Rollback procedures
- Testing environments

### Multi-Environment Setup
- Workspace organization
- Environment isolation
- Configuration management
- Secrets handling
- Consistent deployment

### Terraform Cloud/Enterprise
- VCS integration
- Cost estimation
- Sentinel policies
- Team management
- API-driven workflows

### State Management
- Remote state setup
- Locking mechanism
- Encryption
- Backup strategy
- Recovery procedures

### Monitoring & Auditing
- Resource tracking
- Change logging
- Cost monitoring
- Performance metrics
- Compliance audits

### Troubleshooting
- Common issues
- State corruption
- Lock file issues
- Provider errors
- Dependency resolution
- Performance optimization

## Best Practices

- Automate deployments
- Comprehensive planning
- Team code reviews
- Incremental changes
- Regular backups
- Disaster recovery
- Continuous monitoring
- Documentation
