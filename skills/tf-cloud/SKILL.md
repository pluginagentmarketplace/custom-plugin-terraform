---
name: tf-cloud
description: Terraform Cloud and Enterprise
sasmp_version: "1.3.0"
bonded_agent: tf-cloud
bond_type: PRIMARY_BOND
---

# Terraform Cloud Skill

## Configuration

```hcl
terraform {
  cloud {
    organization = "my-org"
    
    workspaces {
      name = "my-app-prod"
    }
  }
}

# Or with tags
terraform {
  cloud {
    organization = "my-org"
    
    workspaces {
      tags = ["app:myapp", "env:prod"]
    }
  }
}
```

## CLI Workflow

```bash
# Login
terraform login

# Init with cloud
terraform init

# Run operations remotely
terraform plan
terraform apply
```

## Features

- Remote state management
- VCS integration (GitHub, GitLab)
- Policy as Code (Sentinel)
- Cost estimation
- Private registry
- Team management
- SSO/SAML

## Quick Reference

| Feature | Description |
|---------|-------------|
| Workspaces | Environment isolation |
| Runs | Plan/apply execution |
| Policies | Governance rules |
| Modules | Private registry |

## Related
- tf-state - Remote state
- tf-cloud agent
