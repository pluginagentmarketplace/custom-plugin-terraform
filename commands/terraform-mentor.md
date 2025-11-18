# /terraform-mentor

Expert guidance on Terraform architecture, design decisions, and complex problems.

## Mentoring Topics

### Design & Architecture
- Module composition
- Monorepo vs polyrepo
- State strategy
- Provider selection
- Scaling patterns

### Implementation
- HCL patterns
- Complex expressions
- Dynamic blocks
- Provider selection
- Resource design

### Operations
- Deployment strategies
- State management
- Multi-environment
- Disaster recovery
- Cost optimization

### Advanced Topics
- Custom providers
- Plugin development
- Enterprise governance
- Large-scale operations
- Multi-cloud strategies

## How to Get Help

```
/terraform-mentor

Topic: "Should I use workspaces or separate projects?"

Your Context:
- Team size: 10 engineers
- Multiple environments: dev, staging, prod
- Multiple projects: 5 microservices
- Using Terraform Cloud

Expert Analysis:
Based on your setup...
[Detailed guidance]
```

## Common Questions

**Q: Workspaces vs separate state files?**
A: Workspaces for same code, separate for different code

**Q: How to structure large projects?**
A: Use modules, organize by component, shared modules

**Q: Should we version modules?**
A: Always! Semantic versioning for safety

**Q: Multi-cloud strategy?**
A: Abstract with modules, provider-specific layers

**Q: Cost optimization?**
A: Right-sizing, spot instances, lifecycle rules

## Real-World Scenarios

- Migrating from CloudFormation
- Multi-team coordination
- Performance optimization
- Disaster recovery
- Compliance requirements
- Cost reduction
- Team scaling

## Best Practices Guidance

- Code organization
- Testing strategies
- Documentation
- Version control
- CI/CD integration
- Security patterns
- Monitoring setup

Ask specific questions for best guidance!
