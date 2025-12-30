---
description: Generate and analyze Terraform execution plan
allowed-tools: Bash, Read, Grep, Glob
sasmp_version: "1.3.0"
version: "2.0.0"
---

# Terraform Plan Command

Generate an execution plan showing what Terraform will do when you apply changes.

## Usage

```bash
/tf-plan [path] [options]
```

## Arguments

| Argument | Type | Default | Description |
|----------|------|---------|-------------|
| `path` | string | `.` | Directory containing Terraform files |

## Options

| Option | Short | Description |
|--------|-------|-------------|
| `--var-file` | `-f` | Variable definitions file (.tfvars) |
| `--var` | `-v` | Set a variable (key=value) |
| `--target` | `-t` | Resource to target (can be repeated) |
| `--out` | `-o` | Save plan to file for apply |
| `--destroy` | `-d` | Generate destroy plan |
| `--refresh-only` | | Only refresh state, no changes |
| `--replace` | | Force replacement of resource |
| `--lock` | | Lock state file (default: true) |
| `--lock-timeout` | | State lock timeout |
| `--parallelism` | `-p` | Limit concurrent operations |
| `--detailed-exitcode` | | Return detailed exit codes |
| `--json` | `-j` | Output plan in JSON format |
| `--compact` | `-c` | Compact diff output |

## Input Validation

```hcl
validation {
  initialized     = true   # Must run tf-init first
  valid_config    = true   # terraform validate passes
  state_readable  = true   # Can access state backend
  vars_exist      = true   # All var-files exist
}
```

## Exit Codes

### Standard
| Code | Status | Description |
|------|--------|-------------|
| `0` | SUCCESS | Plan generated (may have changes) |
| `1` | ERROR | Plan generation failed |

### Detailed (--detailed-exitcode)
| Code | Status | Description |
|------|--------|-------------|
| `0` | NO_CHANGES | No infrastructure changes |
| `1` | ERROR | Plan generation failed |
| `2` | CHANGES | Plan has changes to apply |

## Examples

### Basic Plan
```bash
# Plan current directory
/tf-plan

# Plan specific path
/tf-plan infrastructure/prod
```

### With Variables
```bash
# Use tfvars file
/tf-plan --var-file=prod.tfvars

# Set inline variables
/tf-plan --var="environment=prod" --var="instance_count=3"

# Multiple var files
/tf-plan --var-file=common.tfvars --var-file=prod.tfvars
```

### Save Plan for Apply
```bash
# Save plan to file
/tf-plan --out=tfplan

# Apply saved plan later
/tf-apply tfplan
```

### Targeted Planning
```bash
# Plan specific resource
/tf-plan --target=aws_instance.web

# Plan module
/tf-plan --target=module.vpc

# Multiple targets
/tf-plan --target=aws_instance.web --target=aws_security_group.web
```

### Destroy Plan
```bash
# Preview what will be destroyed
/tf-plan --destroy

# Save destroy plan
/tf-plan --destroy --out=destroy.tfplan
```

### CI/CD Pipeline
```bash
# Plan with detailed exit code for automation
/tf-plan --detailed-exitcode --json > plan.json
echo "Exit code: $?"

# Limit parallelism for rate-limited APIs
/tf-plan --parallelism=5 --out=tfplan
```

### Force Replacement
```bash
# Force recreate a resource
/tf-plan --replace=aws_instance.web

# Replace multiple resources
/tf-plan --replace=aws_instance.web --replace=aws_instance.api
```

## Output Format

### Standard Output
```
Terraform Plan Results
======================
Path: ./infrastructure/prod

Terraform will perform the following actions:

  # aws_instance.web will be created
  + resource "aws_instance" "web" {
      + ami           = "ami-0123456789"
      + instance_type = "t3.medium"
      + tags          = {
          + "Name" = "prod-web"
        }
    }

  # aws_security_group.web will be updated in-place
  ~ resource "aws_security_group" "web" {
      ~ ingress = [
          + {
              + cidr_blocks = ["10.0.0.0/8"]
              + from_port   = 443
              + to_port     = 443
            }
        ]
    }

Plan: 1 to add, 1 to change, 0 to destroy.
```

### JSON Output
```json
{
  "format_version": "1.2",
  "terraform_version": "1.6.0",
  "planned_values": {
    "root_module": {
      "resources": [...]
    }
  },
  "resource_changes": [
    {
      "address": "aws_instance.web",
      "mode": "managed",
      "type": "aws_instance",
      "change": {
        "actions": ["create"],
        "before": null,
        "after": {...}
      }
    }
  ],
  "summary": {
    "add": 1,
    "change": 1,
    "destroy": 0
  }
}
```

## Change Types

| Symbol | Action | Description |
|--------|--------|-------------|
| `+` | Create | New resource will be created |
| `-` | Destroy | Resource will be destroyed |
| `~` | Update | Resource will be updated in-place |
| `-/+` | Replace | Destroy then create (forces new) |
| `+/-` | Replace | Create then destroy (create before) |
| `<=` | Read | Data source will be read |

## Cost Estimation

```bash
# Generate plan JSON for cost tools
/tf-plan --json > plan.json

# Use with Infracost
infracost breakdown --path plan.json
```

## Troubleshooting

| Error | Cause | Solution |
|-------|-------|----------|
| `No configuration files` | Wrong directory | Verify path has .tf files |
| `Backend not initialized` | Missing init | Run `/tf-init` first |
| `State lock timeout` | Concurrent access | Wait or check who has lock |
| `Provider error` | API/auth issue | Check credentials and limits |
| `Variable not set` | Missing required var | Add `--var` or `--var-file` |
| `Cycle detected` | Circular dependency | Review resource dependencies |
| `Resource not found` | Drifted state | Run `terraform refresh` |

## State Refresh Options

```bash
# Refresh-only mode (sync state with infrastructure)
/tf-plan --refresh-only

# Skip refresh for faster planning
/tf-plan --refresh=false

# Note: --refresh=false may show incorrect diff
```

## Sensitive Values

Terraform masks sensitive values in output:
```
~ password = (sensitive value)
```

To see sensitive values:
```bash
# Output as JSON (includes sensitive)
/tf-plan --json | jq '.resource_changes'
```

## Related

- **Command**: tf-init (must init first)
- **Command**: tf-apply (apply the plan)
- **Agent**: 01-terraform-fundamentals
- **Skill**: terraform-fundamentals, terraform-state
