---
name: tf-variables
description: Terraform variables - input, local, and validation
sasmp_version: "1.3.0"
bonded_agent: tf-fundamentals
bond_type: SECONDARY_BOND
---

# Terraform Variables Skill

## Input Variables

```hcl
# variables.tf
variable "instance_type" {
  description = "EC2 instance type"
  type        = string
  default     = "t2.micro"
}

variable "environment" {
  type = string
  validation {
    condition     = contains(["dev", "staging", "prod"], var.environment)
    error_message = "Must be dev, staging, or prod."
  }
}

variable "tags" {
  type = map(string)
  default = {
    Project = "MyApp"
  }
}

# Usage
resource "aws_instance" "web" {
  instance_type = var.instance_type
  tags          = var.tags
}
```

## Variable Assignment

```bash
# terraform.tfvars
instance_type = "t2.small"
environment   = "prod"

# Command line
terraform apply -var="instance_type=t2.large"

# Environment variable
export TF_VAR_instance_type="t2.large"

# .auto.tfvars (auto-loaded)
# prod.auto.tfvars
```

## Local Values

```hcl
locals {
  common_tags = {
    Environment = var.environment
    Project     = var.project_name
    ManagedBy   = "Terraform"
  }
  
  name_prefix = "${var.project}-${var.environment}"
}

resource "aws_instance" "web" {
  tags = merge(local.common_tags, {
    Name = "${local.name_prefix}-web"
  })
}
```

## Quick Reference

| Type | Syntax |
|------|--------|
| string | `"hello"` |
| number | `42` |
| bool | `true` |
| list | `["a", "b"]` |
| map | `{ key = "val" }` |
| object | `{ name = string }` |

## Related
- tf-outputs - Output values
- tf-hcl - HCL syntax
