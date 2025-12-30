---
description: Apply Terraform changes to infrastructure with safety controls
allowed-tools: Bash, Read, Grep, AskUserQuestion
sasmp_version: "1.3.0"
version: "2.0.0"
---

# Terraform Apply Command

Apply the changes required to reach the desired state of the configuration with built-in safety controls.

## Usage

```bash
/tf-apply [plan-file|path] [options]
```

## Arguments

| Argument | Type | Default | Description |
|----------|------|---------|-------------|
| `plan-file` | string | - | Saved plan file from tf-plan |
| `path` | string | `.` | Directory (if no plan file) |

## Options

| Option | Short | Description |
|--------|-------|-------------|
| `--auto-approve` | `-y` | Skip interactive approval |
| `--var-file` | `-f` | Variable definitions file |
| `--var` | `-v` | Set a variable (key=value) |
| `--target` | `-t` | Resource to target |
| `--replace` | | Force replacement of resource |
| `--destroy` | `-d` | Destroy all resources |
| `--lock` | | Lock state (default: true) |
| `--lock-timeout` | | State lock timeout |
| `--parallelism` | `-p` | Limit concurrent operations |
| `--refresh` | | Refresh state (default: true) |
| `--backup` | | Path to backup state file |
| `--json` | `-j` | Output in JSON format |

## Input Validation

```hcl
validation {
  initialized     = true   # Must run tf-init first
  valid_config    = true   # terraform validate passes
  state_writable  = true   # Can write to state backend
  vars_exist      = true   # All var-files exist
  plan_valid      = true   # Plan file not expired (if provided)
}
```

## Exit Codes

| Code | Status | Description |
|------|--------|-------------|
| `0` | SUCCESS | Apply completed successfully |
| `1` | ERROR | Apply failed |
| `2` | PARTIAL | Some resources failed |
| `3` | LOCK_ERROR | State lock failed |
| `4` | STATE_ERROR | State save failed |
| `5` | CANCELLED | User cancelled apply |
| `126` | PERMISSION_ERROR | Insufficient permissions |
| `127` | CONFIG_ERROR | Invalid configuration |

## Examples

### Basic Apply
```bash
# Apply with confirmation prompt
/tf-apply

# Apply specific directory
/tf-apply infrastructure/prod
```

### Apply Saved Plan
```bash
# Generate plan first
/tf-plan --out=tfplan

# Apply saved plan (no additional confirmation)
/tf-apply tfplan
```

### Auto-Approve (CI/CD)
```bash
# Skip confirmation (use with caution!)
/tf-apply --auto-approve

# With variables
/tf-apply --auto-approve --var-file=prod.tfvars
```

### Targeted Apply
```bash
# Apply specific resource
/tf-apply --target=aws_instance.web

# Apply module
/tf-apply --target=module.vpc

# Multiple targets
/tf-apply --target=aws_instance.web --target=aws_security_group.web
```

### Destroy Infrastructure
```bash
# Destroy with confirmation
/tf-apply --destroy

# Auto-approve destroy (DANGEROUS)
/tf-apply --destroy --auto-approve
```

### Force Replacement
```bash
# Force recreate resource
/tf-apply --replace=aws_instance.web

# Common use: rotate EC2 instance
/tf-apply --replace=aws_instance.app --auto-approve
```

### Production Safety
```bash
# Limit parallelism to avoid rate limits
/tf-apply --parallelism=5

# Explicit backup location
/tf-apply --backup=./backups/terraform.tfstate.backup
```

## Safety Controls

### 1. Confirmation Prompt
```
Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes
```

### 2. Plan Review
Before applying, always review the plan:
```bash
# Review plan first
/tf-plan

# Then apply if satisfied
/tf-apply
```

### 3. Saved Plan Pattern
```bash
# Generate and save plan
/tf-plan --out=tfplan

# Review plan file
terraform show tfplan

# Apply exact plan
/tf-apply tfplan
```

### 4. State Locking
- Automatic state locking prevents concurrent modifications
- Use `--lock-timeout` in CI/CD for retry behavior

## Output Format

### Standard Output
```
Terraform Apply Results
=======================
Path: ./infrastructure/prod

aws_security_group.web: Creating...
aws_security_group.web: Creation complete after 2s [id=sg-0123456789]

aws_instance.web: Creating...
aws_instance.web: Still creating... [10s elapsed]
aws_instance.web: Still creating... [20s elapsed]
aws_instance.web: Creation complete after 25s [id=i-0123456789]

Apply complete! Resources: 2 added, 0 changed, 0 destroyed.

Outputs:
instance_ip = "10.0.1.50"
```

### JSON Output
```json
{
  "format_version": "1.0",
  "terraform_version": "1.6.0",
  "changes": {
    "add": 2,
    "change": 0,
    "destroy": 0
  },
  "resources": [
    {
      "address": "aws_instance.web",
      "action": "create",
      "status": "complete",
      "id": "i-0123456789"
    }
  ],
  "outputs": {
    "instance_ip": {
      "value": "10.0.1.50",
      "type": "string"
    }
  },
  "exit_code": 0
}
```

## Rollback Strategies

### Using State Backup
```bash
# Restore from automatic backup
terraform state push terraform.tfstate.backup

# Or restore specific version from S3
aws s3 cp s3://bucket/state/terraform.tfstate.123 terraform.tfstate
terraform state push terraform.tfstate
```

### Using Version Control
```bash
# Revert to previous commit
git checkout HEAD~1 -- *.tf

# Plan the revert
/tf-plan

# Apply previous state
/tf-apply
```

### Manual Resource Recovery
```bash
# Import existing resource to state
terraform import aws_instance.web i-existing123

# Remove failed resource from state
terraform state rm aws_instance.failed
```

## Troubleshooting

| Error | Cause | Solution |
|-------|-------|----------|
| `State locked` | Concurrent apply | Wait or force-unlock (careful!) |
| `Resource already exists` | Drift or import needed | Import resource or update config |
| `Timeout waiting` | Slow resource creation | Increase timeout in resource config |
| `Rate exceeded` | API throttling | Reduce `--parallelism` |
| `Access denied` | IAM permissions | Check and update credentials |
| `Plan expired` | Old plan file | Regenerate plan with `/tf-plan` |
| `Partial failure` | Some resources failed | Fix and re-apply remaining |

## Force Unlock (Emergency)

```bash
# Get lock ID from error message
# Error: Error acquiring state lock: ...Lock ID: abc-123...

# Force unlock (DANGEROUS - verify no one else is running)
terraform force-unlock abc-123

# Then retry apply
/tf-apply
```

## Post-Apply Verification

```bash
# Verify state matches infrastructure
terraform plan
# Should show: No changes

# Check outputs
terraform output

# Verify specific resource
terraform state show aws_instance.web
```

## CI/CD Best Practices

```bash
# 1. Generate and save plan in plan stage
/tf-plan --out=tfplan --detailed-exitcode
PLAN_EXIT=$?

# 2. Apply only if changes exist (exit code 2)
if [ $PLAN_EXIT -eq 2 ]; then
  /tf-apply tfplan
fi

# 3. Always use saved plans in pipelines
# Never use --auto-approve without saved plan
```

## Related

- **Command**: tf-plan (always plan first)
- **Command**: tf-check (validate before apply)
- **Agent**: 01-terraform-fundamentals
- **Skill**: terraform-fundamentals, terraform-state
