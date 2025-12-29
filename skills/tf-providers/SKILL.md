---
name: tf-providers
description: Terraform providers - configuration and multi-cloud
sasmp_version: "1.3.0"
bonded_agent: tf-providers
bond_type: PRIMARY_BOND
---

# Terraform Providers Skill

## Provider Configuration

```hcl
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 3.0"
    }
  }
}

# AWS Provider
provider "aws" {
  region  = "us-east-1"
  profile = "production"
}

# Provider alias
provider "aws" {
  alias  = "west"
  region = "us-west-2"
}

resource "aws_instance" "west" {
  provider = aws.west
  # ...
}
```

## Authentication

```hcl
# AWS - Environment variables
# AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY

# AWS - Shared credentials
provider "aws" {
  shared_credentials_files = ["~/.aws/credentials"]
  profile                  = "production"
}

# Azure
provider "azurerm" {
  features {}
  subscription_id = var.subscription_id
  tenant_id       = var.tenant_id
}

# GCP
provider "google" {
  project     = var.project_id
  region      = "us-central1"
  credentials = file("account.json")
}
```

## Quick Reference

| Provider | Source |
|----------|--------|
| AWS | hashicorp/aws |
| Azure | hashicorp/azurerm |
| GCP | hashicorp/google |
| Kubernetes | hashicorp/kubernetes |
| Docker | kreuzwerker/docker |

## Related
- tf-resources - Resource creation
- tf-providers agent
