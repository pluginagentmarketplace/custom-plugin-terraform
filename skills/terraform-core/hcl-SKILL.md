---
name: hcl-language
description: Master HCL syntax, expressions, functions, and advanced patterns. Deep dive into dynamic blocks, loops, conditionals, and complex transformations.
---

# HCL Language Mastery

## Quick Start - Complex Expressions

```hcl
# Conditional expression
environment = var.is_prod ? "production" : "development"

# For expressions - create map
resource_tags = {
  for name in var.resource_names :
  name => {
    Name = name
    Environment = var.environment
  }
}

# Dynamic blocks
resource "aws_security_group" "main" {
  dynamic "ingress" {
    for_each = var.ingress_rules
    content {
      from_port   = ingress.value.from_port
      to_port     = ingress.value.to_port
      protocol    = ingress.value.protocol
      cidr_blocks = ingress.value.cidr_blocks
    }
  }
}

# Try-catch error handling
locals {
  safe_value = try(var.optional_value, "default")
}

# Merge and complex operations
merged_config = merge(local.default_config, var.custom_config)
```

## Core Features

### Types
- string, number, bool
- list, map, set, tuple, object
- Type constraints and conversion

### Operators
- Arithmetic: +, -, *, /, %
- Comparison: ==, !=, <, >, <=, >=
- Logical: &&, ||, !

### Functions (150+)
- String: join, split, format, regex
- Collection: keys, values, merge, flatten
- Math: min, max, sum, ceil, floor
- Encoding: base64, jsonencode, yamlparse

### Advanced Patterns
- Recursive operations
- Complex type transformations
- Conditional logic chains
- Custom validation rules
- Performance optimization

## Best Practices

- Keep expressions readable
- Use locals for reusability
- Avoid deep nesting
- Comment complex logic
- Validate input types
