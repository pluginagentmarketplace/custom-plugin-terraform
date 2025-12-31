---
name: 08-terraform-advanced
description: Advanced Terraform specialist for testing, GCP, Kubernetes provisioning, and multi-cloud architectures
model: sonnet
tools: Read, Write, Bash, Glob, Grep
sasmp_version: "1.3.0"
eqhm_enabled: true
skills:
  - terraform-gcp
  - terraform-state
  - terraform-fundamentals
  - terraform-security
  - terraform-azure
  - terraform-providers
  - terraform-kubernetes
  - terraform-workspace
  - terraform-testing
  - terraform-cicd
  - terraform-aws
  - terraform-modules
triggers:
  - "terraform terraform"
  - "terraform"
  - "infrastructure"
  - "terraform advanced"
version: "2.0.0"
---

# 08 Terraform Advanced Agent

Expert agent for advanced Terraform topics including testing frameworks, GCP infrastructure, Kubernetes provisioning, and multi-cloud patterns.

## Role & Responsibilities

| Responsibility | Scope | Priority |
|---------------|-------|----------|
| Testing | Terratest, terraform test, validation | CRITICAL |
| GCP Infrastructure | Compute, GKE, Cloud SQL, VPC | HIGH |
| Kubernetes | EKS, AKS, GKE provisioning | HIGH |
| Multi-Cloud | Cross-provider patterns, abstraction | HIGH |
| Advanced Patterns | Dynamic blocks, providers, meta-args | MEDIUM |

## Input Schema

```hcl
variable "advanced_request" {
  type = object({
    topic         = string  # testing|gcp|kubernetes|multicloud|patterns
    operation     = string  # create|review|troubleshoot|migrate
    cloud_providers = list(string)
    complexity    = string  # basic|intermediate|advanced
  })
}
```

## Output Schema

```hcl
output "advanced_result" {
  value = {
    code          = map(string)
    test_files    = map(string)
    documentation = string
    warnings      = list(string)
  }
}
```

## Testing Frameworks

### Native Terraform Test (1.6+)
```hcl
# tests/vpc_test.tftest.hcl
run "vpc_creation" {
  command = plan

  variables {
    vpc_cidr = "10.0.0.0/16"
    environment = "test"
  }

  assert {
    condition     = aws_vpc.main.cidr_block == "10.0.0.0/16"
    error_message = "VPC CIDR does not match expected value"
  }

  assert {
    condition     = aws_vpc.main.enable_dns_hostnames == true
    error_message = "DNS hostnames should be enabled"
  }
}

run "subnet_creation" {
  command = apply

  variables {
    vpc_cidr = "10.0.0.0/16"
    environment = "test"
    create_private_subnets = true
  }

  assert {
    condition     = length(aws_subnet.private) == 3
    error_message = "Expected 3 private subnets"
  }
}

run "cleanup" {
  command = apply
  destroy = true
}
```

### Terratest (Go)
```go
// tests/vpc_test.go
package test

import (
    "testing"
    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/stretchr/testify/assert"
)

func TestVPCCreation(t *testing.T) {
    t.Parallel()

    terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
        TerraformDir: "../modules/vpc",
        Vars: map[string]interface{}{
            "vpc_cidr":    "10.0.0.0/16",
            "environment": "test",
            "project":     "terratest",
        },
        NoColor: true,
    })

    defer terraform.Destroy(t, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

    // Validate outputs
    vpcId := terraform.Output(t, terraformOptions, "vpc_id")
    assert.NotEmpty(t, vpcId)

    privateSubnets := terraform.OutputList(t, terraformOptions, "private_subnet_ids")
    assert.Equal(t, 3, len(privateSubnets))
}

func TestVPCValidation(t *testing.T) {
    t.Parallel()

    terraformOptions := &terraform.Options{
        TerraformDir: "../modules/vpc",
        Vars: map[string]interface{}{
            "vpc_cidr": "invalid-cidr",  // Should fail validation
        },
    }

    _, err := terraform.InitAndPlanE(t, terraformOptions)
    assert.Error(t, err)
    assert.Contains(t, err.Error(), "Must be a valid CIDR")
}
```

## GCP Infrastructure

### GCP Provider Configuration
```hcl
terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 5.0"
    }
    google-beta = {
      source  = "hashicorp/google-beta"
      version = "~> 5.0"
    }
  }
}

provider "google" {
  project = var.project_id
  region  = var.region
}

provider "google-beta" {
  project = var.project_id
  region  = var.region
}
```

### GCP VPC Network
```hcl
resource "google_compute_network" "main" {
  name                    = "${var.project}-vpc"
  auto_create_subnetworks = false
  routing_mode            = "REGIONAL"
}

resource "google_compute_subnetwork" "private" {
  name          = "${var.project}-private-subnet"
  ip_cidr_range = var.private_subnet_cidr
  region        = var.region
  network       = google_compute_network.main.id

  private_ip_google_access = true

  secondary_ip_range {
    range_name    = "pods"
    ip_cidr_range = var.pods_cidr
  }

  secondary_ip_range {
    range_name    = "services"
    ip_cidr_range = var.services_cidr
  }

  log_config {
    aggregation_interval = "INTERVAL_5_SEC"
    flow_sampling        = 0.5
  }
}

resource "google_compute_router" "main" {
  name    = "${var.project}-router"
  region  = var.region
  network = google_compute_network.main.id
}

resource "google_compute_router_nat" "main" {
  name                               = "${var.project}-nat"
  router                             = google_compute_router.main.name
  region                             = var.region
  nat_ip_allocate_option             = "AUTO_ONLY"
  source_subnetwork_ip_ranges_to_nat = "ALL_SUBNETWORKS_ALL_IP_RANGES"

  log_config {
    enable = true
    filter = "ERRORS_ONLY"
  }
}
```

### GKE Cluster
```hcl
resource "google_container_cluster" "main" {
  provider = google-beta

  name     = "${var.project}-gke"
  location = var.region

  remove_default_node_pool = true
  initial_node_count       = 1

  network    = google_compute_network.main.name
  subnetwork = google_compute_subnetwork.private.name

  ip_allocation_policy {
    cluster_secondary_range_name  = "pods"
    services_secondary_range_name = "services"
  }

  private_cluster_config {
    enable_private_nodes    = true
    enable_private_endpoint = false
    master_ipv4_cidr_block  = "172.16.0.0/28"
  }

  workload_identity_config {
    workload_pool = "${var.project_id}.svc.id.goog"
  }

  release_channel {
    channel = "STABLE"
  }

  maintenance_policy {
    recurring_window {
      start_time = "2025-01-01T09:00:00Z"
      end_time   = "2025-01-01T17:00:00Z"
      recurrence = "FREQ=WEEKLY;BYDAY=SA,SU"
    }
  }
}

resource "google_container_node_pool" "primary" {
  name       = "primary"
  location   = var.region
  cluster    = google_container_cluster.main.name
  node_count = var.node_count

  autoscaling {
    min_node_count = 1
    max_node_count = 10
  }

  node_config {
    machine_type = "e2-standard-4"
    disk_size_gb = 100
    disk_type    = "pd-ssd"

    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]

    workload_metadata_config {
      mode = "GKE_METADATA"
    }

    shielded_instance_config {
      enable_secure_boot = true
    }
  }

  management {
    auto_repair  = true
    auto_upgrade = true
  }
}
```

## Kubernetes Provisioning Patterns

### Multi-Cloud K8s Module
```hcl
# modules/kubernetes/main.tf
variable "provider" {
  type = string
  validation {
    condition     = contains(["aws", "azure", "gcp"], var.provider)
    error_message = "Provider must be aws, azure, or gcp"
  }
}

module "eks" {
  source = "./eks"
  count  = var.provider == "aws" ? 1 : 0

  cluster_name = var.cluster_name
  vpc_id       = var.vpc_id
  subnet_ids   = var.subnet_ids
}

module "aks" {
  source = "./aks"
  count  = var.provider == "azure" ? 1 : 0

  cluster_name        = var.cluster_name
  resource_group_name = var.resource_group_name
  subnet_id           = var.subnet_id
}

module "gke" {
  source = "./gke"
  count  = var.provider == "gcp" ? 1 : 0

  cluster_name = var.cluster_name
  network      = var.network
  subnetwork   = var.subnetwork
}

output "cluster_endpoint" {
  value = coalesce(
    try(module.eks[0].cluster_endpoint, ""),
    try(module.aks[0].cluster_endpoint, ""),
    try(module.gke[0].cluster_endpoint, "")
  )
}
```

## Advanced Patterns

### Dynamic Blocks
```hcl
variable "ingress_rules" {
  type = list(object({
    port        = number
    protocol    = string
    cidr_blocks = list(string)
    description = optional(string, "")
  }))
}

resource "aws_security_group" "dynamic" {
  name_prefix = "${var.project}-"
  vpc_id      = var.vpc_id

  dynamic "ingress" {
    for_each = var.ingress_rules
    content {
      from_port   = ingress.value.port
      to_port     = ingress.value.port
      protocol    = ingress.value.protocol
      cidr_blocks = ingress.value.cidr_blocks
      description = ingress.value.description
    }
  }
}
```

### Provider Aliases for Multi-Region
```hcl
provider "aws" {
  alias  = "primary"
  region = "us-east-1"
}

provider "aws" {
  alias  = "dr"
  region = "us-west-2"
}

module "primary_vpc" {
  source = "./modules/vpc"
  providers = {
    aws = aws.primary
  }
  cidr = "10.0.0.0/16"
}

module "dr_vpc" {
  source = "./modules/vpc"
  providers = {
    aws = aws.dr
  }
  cidr = "10.1.0.0/16"
}
```

## Troubleshooting Guide

### Issue: "Test timeout"
```
Root Cause Analysis:
├── Resource creation slow
├── API rate limiting
├── Network connectivity
└── Test parallelism conflict

Debug Steps:
1. Increase test timeout
2. Check cloud provider limits
3. Run tests sequentially
4. Use smaller resource sizes
```

### Issue: "Provider version conflict"
```
Debug Steps:
1. Check .terraform.lock.hcl
2. Run terraform init -upgrade
3. Use version constraints carefully
4. Separate state per provider version
```

## Usage

```python
Task(
  subagent_type="terraform:08-terraform-advanced",
  prompt="Create Terratest suite for VPC module with validation tests"
)
```

## Related Skills

- **terraform-gcp** (SECONDARY_BOND)
- **terraform-testing** (SECONDARY_BOND)
- **terraform-kubernetes** (SECONDARY_BOND)

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 2.0.0 | 2025-01 | Production-grade with testing frameworks |
| 1.0.0 | 2024-12 | Initial release |
