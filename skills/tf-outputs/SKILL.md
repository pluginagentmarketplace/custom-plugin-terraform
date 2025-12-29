---
name: tf-outputs
description: Terraform outputs - values and module returns
sasmp_version: "1.3.0"
bonded_agent: tf-fundamentals
bond_type: SECONDARY_BOND
---

# Terraform Outputs Skill

## Output Syntax

```hcl
output "instance_ip" {
  description = "Public IP of the instance"
  value       = aws_instance.web.public_ip
}

output "db_connection" {
  description = "Database connection string"
  value       = aws_db_instance.main.endpoint
  sensitive   = true
}

output "all_instance_ids" {
  value = aws_instance.web[*].id
}
```

## Output Features

```hcl
# Conditional output
output "load_balancer_dns" {
  value = var.create_lb ? aws_lb.main[0].dns_name : null
}

# Complex output
output "server_info" {
  value = {
    id         = aws_instance.web.id
    public_ip  = aws_instance.web.public_ip
    private_ip = aws_instance.web.private_ip
  }
}

# precondition
output "instance_ip" {
  value = aws_instance.web.public_ip
  precondition {
    condition     = aws_instance.web.public_ip != ""
    error_message = "Instance has no public IP."
  }
}
```

## Access Outputs

```bash
# Show all outputs
terraform output

# Show specific output
terraform output instance_ip

# JSON format
terraform output -json

# Raw value
terraform output -raw instance_ip
```

## Quick Reference

| Argument | Purpose |
|----------|---------|
| value | Output value |
| description | Documentation |
| sensitive | Hide value |
| depends_on | Dependencies |
| precondition | Validation |

## Related
- tf-variables - Input variables
- tf-modules - Module outputs
