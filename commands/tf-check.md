---
description: Run Terraform validation, formatting, and security checks
allowed-tools: Bash, Read, Grep, Glob
sasmp_version: "1.3.0"
version: "2.0.0"
---

# Terraform Check Command

Comprehensive validation of Terraform configurations including formatting, syntax, security, and best practices.

## Usage

```bash
/tf-check [path] [options]
```

## Arguments

| Argument | Type | Default | Description |
|----------|------|---------|-------------|
| `path` | string | `.` | Directory containing Terraform files |

## Options

| Option | Short | Description |
|--------|-------|-------------|
| `--format` | `-f` | Run terraform fmt check |
| `--validate` | `-v` | Run terraform validate |
| `--lint` | `-l` | Run tflint static analysis |
| `--security` | `-s` | Run security scans (tfsec/checkov) |
| `--all` | `-a` | Run all checks (default) |
| `--fix` | | Auto-fix formatting issues |
| `--json` | `-j` | Output results in JSON format |

## Input Validation

```hcl
validation {
  path_exists     = true
  contains_tf     = true  # Must have *.tf files
  readable        = true
  max_file_size   = "10MB"
}
```

## Exit Codes

| Code | Status | Description |
|------|--------|-------------|
| `0` | SUCCESS | All checks passed |
| `1` | FORMAT_ERROR | Formatting issues found |
| `2` | VALIDATE_ERROR | Syntax/validation errors |
| `3` | LINT_ERROR | Linting warnings/errors |
| `4` | SECURITY_ERROR | Security vulnerabilities found |
| `5` | NOT_INITIALIZED | Terraform not initialized |
| `126` | PERMISSION_ERROR | Cannot access path |
| `127` | TOOL_NOT_FOUND | Required tool not installed |

## Examples

### Basic Validation
```bash
# Check current directory
/tf-check

# Check specific module
/tf-check modules/vpc
```

### Format Check with Auto-Fix
```bash
# Check formatting only
/tf-check --format

# Auto-fix formatting issues
/tf-check --format --fix
```

### Security Scan
```bash
# Run security checks
/tf-check --security

# Full security audit with JSON output
/tf-check --security --json > security-report.json
```

### CI/CD Pipeline
```bash
# Complete check for pipelines
/tf-check --all --json

# Quick validation only
/tf-check --validate --format
```

## Check Details

### 1. Format Check (--format)
```bash
terraform fmt -check -recursive -diff
```
- Verifies HCL formatting standards
- Reports files requiring formatting
- Optional auto-fix with `--fix`

### 2. Validation (--validate)
```bash
terraform validate
```
- Syntax verification
- Provider schema validation
- Resource attribute checking

### 3. Linting (--lint)
```bash
tflint --recursive
```
- AWS/Azure/GCP rule validation
- Deprecated syntax detection
- Best practice enforcement

### 4. Security Scan (--security)
```bash
tfsec .
checkov -d .
```
- Misconfiguration detection
- CIS benchmark compliance
- OWASP vulnerability scanning

## Output Format

### Standard Output
```
Terraform Check Results
=======================
Path: ./modules/vpc

✓ Format Check    PASSED
✓ Validation      PASSED
⚠ Linting         2 warnings
✗ Security        1 high, 3 medium

Summary: 2 checks passed, 1 warning, 1 failed
```

### JSON Output
```json
{
  "path": "./modules/vpc",
  "timestamp": "2024-01-15T10:30:00Z",
  "checks": {
    "format": {"status": "passed", "files_checked": 12},
    "validate": {"status": "passed"},
    "lint": {"status": "warning", "warnings": 2},
    "security": {"status": "failed", "high": 1, "medium": 3}
  },
  "exit_code": 4
}
```

## Troubleshooting

| Error | Cause | Solution |
|-------|-------|----------|
| `No .tf files found` | Empty or wrong path | Verify path contains Terraform files |
| `terraform not found` | CLI not installed | Install Terraform CLI |
| `tflint not found` | Linter not installed | `brew install tflint` or skip with `--validate` |
| `Backend not initialized` | Missing init | Run `/tf-init` first |
| `Provider schema missing` | Old lock file | Run `terraform init -upgrade` |

## Related

- **Command**: tf-init (initialize before checking)
- **Agent**: 01-terraform-fundamentals
- **Skill**: terraform-fundamentals
