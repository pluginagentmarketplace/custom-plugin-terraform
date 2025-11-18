# /module-builder

Build production-grade Terraform modules step-by-step.

## Module Structure

```
modules/vpc/
├── main.tf           # Resource definitions
├── variables.tf      # Input variables
├── outputs.tf        # Output values
├── locals.tf         # Local values
├── README.md         # Documentation
└── examples/
    └── basic/
        ├── main.tf
        └── terraform.tfvars
```

## Step-by-Step Guide

### Step 1: Define Scope (1 day)
- What does the module do?
- What variables will it accept?
- What outputs will it provide?
- What are edge cases?

### Step 2: Create Structure (1 day)
- Create main.tf with resources
- Define all variables
- Create outputs
- Add documentation

### Step 3: Develop & Test (3-5 days)
- Implement resources
- Create examples
- Manual testing
- Refine based on testing

### Step 4: Add Tests (2-3 days)
- Write Terratest
- Test all scenarios
- Test edge cases
- Achieve 100% coverage

### Step 5: Documentation (1 day)
- Write README
- Document variables
- Document outputs
- Create examples

### Step 6: Publish (1 day)
- Semantic versioning
- Release tags
- Registry publication
- Announce release

## Module Best Practices

- Single responsibility
- Clear input/output
- Sensible defaults
- Comprehensive validation
- Complete documentation
- Example configurations
- Automated tests
- Version constraints

## Common Modules

- VPC and networking
- Compute (EC2, Fargate)
- Database (RDS, DynamoDB)
- Load balancers
- Security groups
- Monitoring
- Storage (S3, EBS)

Get help with `/terraform-mentor module`
