---
description: Run Terraform validation and format checks
allowed-tools: Bash, Read, Grep
---

# Terraform Check Command

Validate and format Terraform configuration.

## Usage
```
/tf-check [path]
```

## Checks
1. terraform fmt -check
2. terraform validate
3. tflint (if available)
4. terraform plan (dry-run)
