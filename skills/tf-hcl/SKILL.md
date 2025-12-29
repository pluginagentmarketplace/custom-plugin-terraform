---
name: tf-hcl
description: HCL (HashiCorp Configuration Language) syntax
sasmp_version: "1.3.0"
bonded_agent: tf-fundamentals
bond_type: PRIMARY_BOND
---

# HCL Syntax Skill

## Blocks

```hcl
# Block syntax
block_type "label1" "label2" {
  argument = "value"
  
  nested_block {
    argument = "value"
  }
}

# Resource block
resource "aws_instance" "example" {
  ami           = "ami-123456"
  instance_type = "t2.micro"
}

# Data block
data "aws_ami" "ubuntu" {
  most_recent = true
  owners      = ["099720109477"]
}
```

## Expressions

```hcl
# References
aws_instance.example.id
aws_instance.example.public_ip
var.instance_type
local.common_tags
data.aws_ami.ubuntu.id

# String interpolation
name = "web-${var.environment}"

# Conditional
instance_type = var.env == "prod" ? "t2.large" : "t2.micro"

# For expressions
tags = { for k, v in var.tags : k => upper(v) }
ids  = [for s in var.subnets : s.id]
```

## Types

```hcl
# Primitive
string  = "hello"
number  = 42
bool    = true

# Complex
list    = ["a", "b", "c"]
map     = { key = "value" }
set     = toset(["a", "b"])
object  = { name = string, age = number }
tuple   = [string, number, bool]
```

## Operators

```hcl
# Arithmetic
a + b, a - b, a * b, a / b, a % b

# Comparison
a == b, a != b, a < b, a > b, a <= b, a >= b

# Logical
a && b, a || b, !a
```

## Quick Reference

| Block | Purpose |
|-------|---------|
| resource | Create infrastructure |
| data | Query existing resources |
| variable | Input variables |
| output | Output values |
| locals | Local values |
| module | Reusable modules |
| provider | Provider config |

## Related
- tf-basics - Core concepts
- tf-variables - Variables
