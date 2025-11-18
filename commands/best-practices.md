# /best-practices

Production-grade Terraform best practices and optimization patterns.

## Core Best Practices

### Code Organization
```
├── modules/              # Reusable modules
│   ├── vpc/
│   ├── compute/
│   └── database/
├── environments/         # Environment configs
│   ├── dev/
│   ├── staging/
│   └── prod/
├── global/              # Shared resources
└── docs/                # Documentation
```

### State Management
✅ DO:
- Use remote state (S3, Terraform Cloud)
- Enable state locking
- Encrypt state files
- Regular backups
- Version control for configs (not state files)
- Separate state for environments

❌ DON'T:
- Commit state files to git
- Manually edit state
- Share state across teams without locking
- Use default local state in production

### Resource Naming
```hcl
# Consistent naming convention
resource "aws_vpc" "main" {
  # Prefix: environment-application-component
}

resource "aws_instance" "app_server" {
  tags = {
    Name        = "${var.environment}-app-server"
    Environment = var.environment
    Project     = var.project
  }
}
```

### Variables & Outputs
- Use descriptive names
- Always validate input
- Provide defaults
- Document all variables
- Output important values
- Mark sensitive data

### Module Best Practices
- Single responsibility
- Reusable and flexible
- Comprehensive inputs
- Clear outputs
- Example configurations
- Automated tests
- Semantic versioning

### Error Handling
- Input validation
- Proper error messages
- Graceful degradation
- Rollback procedures
- Testing edge cases

### Security
✅ Secure Practices:
- Use IAM roles (not keys)
- Encrypt secrets
- Network isolation
- Security groups/NACLs
- VPC endpoints
- Audit logging
- Regular backups

### Performance Optimization
- Parallel execution (terraform parallelism)
- Dependency management
- Resource batching
- State file size
- Plan optimization
- Cache modules

### CI/CD Integration
- Automated planning
- Code review process
- Automated testing
- Terraform formatting (terraform fmt)
- Linting (tflint)
- Security scanning (checkov)
- Automated apply

### Team Practices
- Code reviews required
- Change approval process
- Documentation requirements
- Consistent formatting
- Version constraints
- Team training
- Runbooks

### Monitoring & Alerting
- Track resource changes
- Cost monitoring
- Compliance audits
- Change logs
- Performance metrics
- Alert on failures

## Anti-Patterns to Avoid

❌ Monolithic configuration
❌ Manual state editing
❌ No testing
❌ Hardcoded values
❌ No documentation
❌ Large state files
❌ No version constraints
❌ Manual deployments

## Production Checklist

- [ ] Remote state configured
- [ ] State locking enabled
- [ ] Encryption enabled
- [ ] Backups in place
- [ ] Modules tested
- [ ] Documentation complete
- [ ] CI/CD configured
- [ ] Team trained
- [ ] Runbooks written
- [ ] Monitoring setup
- [ ] Change approval process
- [ ] Disaster recovery plan

Get help with `/terraform-mentor best-practices`
