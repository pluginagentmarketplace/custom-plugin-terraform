---
name: tf-best-practices
description: Terraform best practices and patterns
sasmp_version: "1.3.0"
bonded_agent: tf-advanced
bond_type: PRIMARY_BOND
---

# Terraform Best Practices Skill

## Code Organization

```
infrastructure/
├── modules/
│   ├── networking/
│   ├── compute/
│   └── database/
├── environments/
│   ├── dev/
│   ├── staging/
│   └── prod/
└── global/
    └── iam/
```

## Naming Conventions

```hcl
# Resources: lowercase with underscores
resource "aws_instance" "web_server" {}

# Variables: descriptive names
variable "instance_count" {}

# Modules: kebab-case directories
module "web-cluster" {}

# Tags
tags = {
  Name        = "web-server-prod"
  Environment = "production"
  ManagedBy   = "terraform"
}
```

## Security

```hcl
# Never hardcode secrets
variable "db_password" {
  sensitive = true
}

# Use remote state with encryption
backend "s3" {
  encrypt = true
}

# Mark sensitive outputs
output "db_password" {
  value     = random_password.db.result
  sensitive = true
}
```

## Recommendations

1. Use remote state with locking
2. Pin provider versions
3. Use modules for reusability
4. Implement consistent tagging
5. Use workspaces or directories for environments
6. Run terraform fmt and validate in CI
7. Review plans before apply
8. Use -target sparingly

## Quick Reference

| Practice | Why |
|----------|-----|
| Remote state | Team collaboration |
| Version pins | Reproducibility |
| Modules | Reusability |
| Workspaces | Isolation |

## Related
- tf-modules - Module patterns
- tf-state - State management
