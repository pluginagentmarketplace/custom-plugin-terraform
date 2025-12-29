---
name: tf-basics
description: Terraform basics - installation, CLI, and workflow
sasmp_version: "1.3.0"
bonded_agent: tf-fundamentals
bond_type: PRIMARY_BOND
---

# Terraform Basics Skill

## Installation

```bash
# macOS
brew tap hashicorp/tap
brew install hashicorp/tap/terraform

# Linux
wget -O- https://apt.releases.hashicorp.com/gpg | sudo gpg --dearmor -o /usr/share/keyrings/hashicorp-archive-keyring.gpg
echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/hashicorp.list
sudo apt update && sudo apt install terraform

# Verify
terraform -version
```

## Core Workflow

```bash
# Initialize working directory
terraform init

# Preview changes
terraform plan

# Apply changes
terraform apply

# Destroy infrastructure
terraform destroy

# Format code
terraform fmt

# Validate configuration
terraform validate
```

## File Structure

```
project/
├── main.tf          # Main configuration
├── variables.tf     # Input variables
├── outputs.tf       # Output values
├── providers.tf     # Provider configuration
├── terraform.tfvars # Variable values
├── .terraform/      # Plugins and modules
└── terraform.tfstate # State file
```

## Basic Example

```hcl
# providers.tf
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

# main.tf
resource "aws_instance" "example" {
  ami           = "ami-0c55b159cbfafe1f0"
  instance_type = "t2.micro"

  tags = {
    Name = "ExampleInstance"
  }
}
```

## Quick Reference

| Command | Description |
|---------|-------------|
| terraform init | Initialize directory |
| terraform plan | Preview changes |
| terraform apply | Apply changes |
| terraform destroy | Destroy resources |
| terraform fmt | Format code |
| terraform validate | Validate config |

## Related
- tf-hcl - HCL syntax
- tf-fundamentals agent
