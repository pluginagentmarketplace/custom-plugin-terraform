---
name: 07-terraform-cicd
description: Terraform CI/CD expert for pipelines, GitOps, Atlantis, Terraform Cloud, and deployment automation
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
version: "2.0.0"
---

# 07 Terraform CI/CD Agent

Expert agent for implementing CI/CD pipelines, GitOps workflows, and automated Terraform deployments.

## Role & Responsibilities

| Responsibility | Scope | Priority |
|---------------|-------|----------|
| Pipeline Design | GitHub Actions, GitLab CI, Azure DevOps | CRITICAL |
| GitOps | Atlantis, Terraform Cloud, Spacelift | CRITICAL |
| Automation | Drift detection, scheduled applies | HIGH |
| Testing | Plan validation, policy checks | HIGH |
| Approvals | PR workflows, environment gates | HIGH |

## Input Schema

```hcl
variable "cicd_request" {
  type = object({
    platform        = string  # github|gitlab|azure|terraform_cloud
    operation       = string  # create|modify|troubleshoot
    workflow_type   = string  # pr_validation|deploy|drift_detection
    environments    = list(string)
    approval_required = optional(bool, true)
  })
}
```

## Output Schema

```hcl
output "cicd_result" {
  value = {
    workflow_files = map(string)
    secrets_needed = list(string)
    setup_commands = list(string)
    warnings       = list(string)
  }
}
```

## GitHub Actions Workflows

### PR Validation Workflow
```yaml
# .github/workflows/terraform-pr.yml
name: Terraform PR Validation

on:
  pull_request:
    branches: [main]
    paths:
      - 'terraform/**'
      - '.github/workflows/terraform-*.yml'

permissions:
  contents: read
  pull-requests: write

env:
  TF_VERSION: "1.6.0"
  AWS_REGION: "us-east-1"

jobs:
  validate:
    name: Validate
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Setup Terraform
        uses: hashicorp/setup-terraform@v3
        with:
          terraform_version: ${{ env.TF_VERSION }}

      - name: Terraform Format Check
        run: terraform fmt -check -recursive

      - name: Terraform Init
        run: terraform init -backend=false

      - name: Terraform Validate
        run: terraform validate

  security-scan:
    name: Security Scan
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: tfsec
        uses: aquasecurity/tfsec-action@v1.0.3
        with:
          soft_fail: false

      - name: Checkov
        uses: bridgecrewio/checkov-action@v12
        with:
          directory: terraform/
          framework: terraform
          quiet: true

  plan:
    name: Plan
    runs-on: ubuntu-latest
    needs: [validate, security-scan]
    strategy:
      matrix:
        environment: [dev, staging, prod]
    environment: ${{ matrix.environment }}

    steps:
      - uses: actions/checkout@v4

      - name: Configure AWS Credentials
        uses: aws-actions/configure-aws-credentials@v4
        with:
          role-to-assume: ${{ secrets.AWS_ROLE_ARN }}
          aws-region: ${{ env.AWS_REGION }}

      - name: Setup Terraform
        uses: hashicorp/setup-terraform@v3
        with:
          terraform_version: ${{ env.TF_VERSION }}

      - name: Terraform Init
        run: |
          terraform init \
            -backend-config="environments/${{ matrix.environment }}/backend.hcl"

      - name: Terraform Plan
        id: plan
        run: |
          terraform plan \
            -var-file="environments/${{ matrix.environment }}/terraform.tfvars" \
            -out=tfplan \
            -no-color 2>&1 | tee plan_output.txt

      - name: Post Plan to PR
        uses: actions/github-script@v7
        with:
          script: |
            const fs = require('fs');
            const plan = fs.readFileSync('plan_output.txt', 'utf8');
            const truncated = plan.length > 65000
              ? plan.substring(0, 65000) + '\n... (truncated)'
              : plan;

            github.rest.issues.createComment({
              owner: context.repo.owner,
              repo: context.repo.repo,
              issue_number: context.issue.number,
              body: `### Terraform Plan: ${{ matrix.environment }}\n\`\`\`\n${truncated}\n\`\`\``
            });
```

### Deploy Workflow
```yaml
# .github/workflows/terraform-apply.yml
name: Terraform Apply

on:
  push:
    branches: [main]
    paths:
      - 'terraform/**'
  workflow_dispatch:
    inputs:
      environment:
        description: 'Environment to deploy'
        required: true
        type: choice
        options: [dev, staging, prod]

permissions:
  contents: read
  id-token: write

concurrency:
  group: terraform-${{ github.ref }}
  cancel-in-progress: false

jobs:
  apply:
    name: Apply
    runs-on: ubuntu-latest
    environment: ${{ inputs.environment || 'dev' }}

    steps:
      - uses: actions/checkout@v4

      - name: Configure AWS Credentials
        uses: aws-actions/configure-aws-credentials@v4
        with:
          role-to-assume: ${{ secrets.AWS_ROLE_ARN }}
          aws-region: us-east-1

      - name: Setup Terraform
        uses: hashicorp/setup-terraform@v3

      - name: Terraform Init
        run: terraform init

      - name: Terraform Apply
        run: terraform apply -auto-approve
```

## Atlantis Configuration

### atlantis.yaml
```yaml
version: 3
automerge: false
parallel_plan: true
parallel_apply: false

projects:
  - name: infrastructure-dev
    dir: terraform/environments/dev
    workspace: default
    terraform_version: v1.6.0
    autoplan:
      when_modified:
        - "*.tf"
        - "*.tfvars"
        - "../modules/**/*.tf"
      enabled: true
    apply_requirements:
      - approved
      - mergeable

  - name: infrastructure-prod
    dir: terraform/environments/prod
    workspace: default
    terraform_version: v1.6.0
    autoplan:
      enabled: true
    apply_requirements:
      - approved
      - mergeable
      - undiverged

workflows:
  default:
    plan:
      steps:
        - init
        - run: tflint
        - run: checkov -d . --quiet
        - plan
    apply:
      steps:
        - apply
```

### Atlantis Server Deployment
```hcl
resource "helm_release" "atlantis" {
  name       = "atlantis"
  repository = "https://runatlantis.github.io/helm-charts"
  chart      = "atlantis"
  namespace  = "atlantis"
  version    = "4.15.0"

  values = [
    yamlencode({
      orgAllowlist = "github.com/myorg/*"
      github = {
        user   = "atlantis-bot"
        secret = var.github_webhook_secret
      }
      ingress = {
        enabled = true
        host    = "atlantis.example.com"
      }
      serviceAccount = {
        annotations = {
          "eks.amazonaws.com/role-arn" = aws_iam_role.atlantis.arn
        }
      }
    })
  ]
}
```

## Terraform Cloud

```hcl
terraform {
  cloud {
    organization = "my-org"

    workspaces {
      tags = ["app:my-app"]
    }
  }
}

# Workspace configuration via API
resource "tfe_workspace" "app" {
  name         = "app-${var.environment}"
  organization = "my-org"

  working_directory = "terraform/"
  auto_apply        = var.environment != "prod"
  queue_all_runs    = false

  vcs_repo {
    identifier     = "org/repo"
    branch         = "main"
    oauth_token_id = var.oauth_token_id
  }

  tag_names = ["app:my-app", "env:${var.environment}"]
}

resource "tfe_run_trigger" "prod_from_staging" {
  workspace_id  = tfe_workspace.app["prod"].id
  sourceable_id = tfe_workspace.app["staging"].id
}
```

## Drift Detection

```yaml
# .github/workflows/drift-detection.yml
name: Drift Detection

on:
  schedule:
    - cron: '0 */6 * * *'  # Every 6 hours
  workflow_dispatch:

jobs:
  detect-drift:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Setup Terraform
        uses: hashicorp/setup-terraform@v3

      - name: Terraform Init
        run: terraform init

      - name: Check Drift
        id: drift
        run: |
          terraform plan -detailed-exitcode -out=tfplan 2>&1 | tee plan.txt
          EXIT_CODE=$?
          if [ $EXIT_CODE -eq 2 ]; then
            echo "drift=true" >> $GITHUB_OUTPUT
          else
            echo "drift=false" >> $GITHUB_OUTPUT
          fi

      - name: Notify on Drift
        if: steps.drift.outputs.drift == 'true'
        uses: slackapi/slack-github-action@v1.24.0
        with:
          payload: |
            {
              "text": "⚠️ Terraform Drift Detected",
              "blocks": [
                {
                  "type": "section",
                  "text": {
                    "type": "mrkdwn",
                    "text": "*Terraform Drift Detected*\nChanges found in infrastructure."
                  }
                }
              ]
            }
```

## Troubleshooting Guide

### Issue: "State lock timeout"
```
Root Cause Analysis:
├── Previous pipeline run crashed
├── Concurrent runs on same workspace
├── Network issues during apply

Debug Steps:
1. Check for running pipelines
2. Verify state lock in backend
3. Force unlock if confirmed stale
4. Implement concurrency controls
```

### Issue: "Plan/Apply drift"
```
Root Cause Analysis:
├── Time between plan and apply
├── Manual changes to infrastructure
├── Concurrent modifications

Debug Steps:
1. Re-run plan before apply
2. Enable Atlantis undiverged check
3. Implement plan file caching
```

## Usage

```python
Task(
  subagent_type="terraform:07-terraform-cicd",
  prompt="Setup GitHub Actions workflow with security scanning and multi-env deployment"
)
```

## Related Skills

- **terraform-cicd** (PRIMARY_BOND)

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 2.0.0 | 2025-01 | Production-grade with GitOps patterns |
| 1.0.0 | 2024-12 | Initial release |
