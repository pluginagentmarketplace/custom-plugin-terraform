---
description: Initialize Terraform working directory with providers and modules
allowed-tools: Bash, Read, Grep
sasmp_version: "1.3.0"
version: "2.0.0"
---

# Terraform Init Command

Initialize a Terraform working directory by downloading providers, modules, and configuring backends.

## Usage

```bash
/tf-init [path] [options]
```

## Arguments

| Argument | Type | Default | Description |
|----------|------|---------|-------------|
| `path` | string | `.` | Directory containing Terraform files |

## Options

| Option | Short | Description |
|--------|-------|-------------|
| `--upgrade` | `-u` | Upgrade providers and modules to latest |
| `--reconfigure` | `-r` | Reconfigure backend, ignoring saved config |
| `--migrate-state` | `-m` | Migrate state to new backend |
| `--backend-config` | `-b` | Backend configuration file or key=value |
| `--lock` | | Lock state file (default: true) |
| `--lock-timeout` | | State lock timeout (default: 0s) |
| `--json` | `-j` | Output results in JSON format |

## Input Validation

```hcl
validation {
  path_exists       = true
  contains_tf       = true
  writable          = true   # Needs to create .terraform/
  network_access    = true   # Downloads providers
}
```

## Exit Codes

| Code | Status | Description |
|------|--------|-------------|
| `0` | SUCCESS | Initialization complete |
| `1` | CONFIG_ERROR | Invalid Terraform configuration |
| `2` | BACKEND_ERROR | Backend configuration failed |
| `3` | PROVIDER_ERROR | Provider download/verification failed |
| `4` | MODULE_ERROR | Module download failed |
| `5` | LOCK_ERROR | State lock acquisition failed |
| `126` | PERMISSION_ERROR | Cannot write to directory |
| `127` | NETWORK_ERROR | Cannot reach registry |

## Examples

### Basic Initialization
```bash
# Initialize current directory
/tf-init

# Initialize specific path
/tf-init infrastructure/prod
```

### Upgrade Providers
```bash
# Upgrade all providers to latest compatible versions
/tf-init --upgrade

# Upgrade specific module path
/tf-init modules/vpc --upgrade
```

### Backend Configuration
```bash
# Configure S3 backend
/tf-init --backend-config="bucket=my-tf-state" \
         --backend-config="key=prod/terraform.tfstate" \
         --backend-config="region=us-east-1"

# Use backend config file
/tf-init --backend-config=backend.hcl
```

### State Migration
```bash
# Migrate state to new backend
/tf-init --migrate-state

# Reconfigure backend (discard existing state path)
/tf-init --reconfigure
```

### CI/CD Pipeline
```bash
# Non-interactive init with lock timeout
/tf-init --lock-timeout=5m --json

# Init without state locking (for read-only plans)
/tf-init --lock=false
```

## Initialization Steps

### 1. Backend Initialization
```
Initializing the backend...
  - Configuring remote state storage
  - Acquiring state lock
  - Downloading existing state
```

### 2. Provider Installation
```
Initializing provider plugins...
  - Finding hashicorp/aws versions matching "~> 5.0"...
  - Installing hashicorp/aws v5.31.0...
  - Installed hashicorp/aws v5.31.0 (signed by HashiCorp)
```

### 3. Module Download
```
Initializing modules...
  - vpc in modules/vpc
  - eks in modules/eks
```

### 4. Lock File Update
```
Terraform has created a lock file .terraform.lock.hcl
```

## Output Format

### Standard Output
```
Terraform Init Results
======================
Path: ./infrastructure/prod

✓ Backend        S3 (my-tf-state)
✓ Providers      3 installed
  - aws          v5.31.0
  - kubernetes   v2.24.0
  - helm         v2.12.0
✓ Modules        2 downloaded
✓ Lock File      Updated

Terraform has been successfully initialized!
```

### JSON Output
```json
{
  "path": "./infrastructure/prod",
  "timestamp": "2024-01-15T10:30:00Z",
  "backend": {
    "type": "s3",
    "bucket": "my-tf-state",
    "status": "configured"
  },
  "providers": [
    {"name": "aws", "version": "5.31.0", "source": "hashicorp/aws"},
    {"name": "kubernetes", "version": "2.24.0", "source": "hashicorp/kubernetes"}
  ],
  "modules": [
    {"name": "vpc", "source": "modules/vpc"},
    {"name": "eks", "source": "modules/eks"}
  ],
  "lock_file_updated": true,
  "exit_code": 0
}
```

## Backend Config File

```hcl
# backend.hcl
bucket         = "company-terraform-state"
key            = "prod/terraform.tfstate"
region         = "us-east-1"
dynamodb_table = "terraform-locks"
encrypt        = true
```

## Troubleshooting

| Error | Cause | Solution |
|-------|-------|----------|
| `Backend configuration changed` | Backend settings modified | Use `--reconfigure` or `--migrate-state` |
| `Provider not found` | Registry unreachable | Check network, try `--plugin-dir` |
| `Module not found` | Invalid source | Verify module source path/URL |
| `State lock failed` | Another process running | Wait or use `--lock=false` (careful!) |
| `Permission denied` | No write access | Check directory permissions |
| `Checksum mismatch` | Corrupted download | Delete `.terraform` and re-init |
| `Version constraints` | Incompatible versions | Update version constraints in config |

## Lock File Management

```bash
# Generate lock file for multiple platforms
terraform providers lock \
  -platform=linux_amd64 \
  -platform=darwin_amd64 \
  -platform=darwin_arm64

# Commit lock file to version control
git add .terraform.lock.hcl
```

## Related

- **Command**: tf-check (validate after init)
- **Command**: tf-plan (plan after init)
- **Agent**: 01-terraform-fundamentals
- **Skill**: terraform-fundamentals, terraform-providers
