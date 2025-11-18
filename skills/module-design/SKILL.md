---
name: terraform-modules
description: Design and build reusable, production-grade Terraform modules. Master module structure, composition, testing, and best practices.
---

# Terraform Modules Mastery

## Quick Start

```hcl
# Module structure
# ./modules/vpc/
# ├── main.tf         (resources)
# ├── variables.tf    (inputs)
# ├── outputs.tf      (outputs)
# └── README.md       (documentation)

# Using a module
module "vpc" {
  source = "./modules/vpc"
  version = "~> 1.0"

  name               = "main"
  cidr_block         = "10.0.0.0/16"
  enable_nat_gateway = true

  tags = local.common_tags
}

# Accessing module outputs
output "vpc_id" {
  value = module.vpc.vpc_id
}
```

## Module Design

### Structure
- Root vs child modules
- Flat vs nested hierarchies
- Standard layout
- Documentation requirements

### Best Practices
- Single responsibility
- Reusable variables
- Clear outputs
- Consistent naming
- Type validation
- Variable defaults

### Composition
- Module dependencies
- Variable flow
- Output chaining
- Complex configurations
- Provider passing

### Testing Modules
- Terratest framework
- Unit testing
- Integration testing
- Edge cases
- Validation

### Registry & Versioning
- Semantic versioning
- Changelog management
- Documentation standards
- README requirements
- Module certification

## Common Patterns

- VPC modules
- Compute clusters
- Database modules
- Load balancer modules
- Monitoring modules
- Security modules
