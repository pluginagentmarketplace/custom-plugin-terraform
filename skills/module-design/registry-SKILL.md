---
name: module-registry
description: Publish and manage Terraform modules on public and private registries. Learn module versioning, documentation, and registry operations.
---

# Module Registry Operations

## Quick Start

```bash
# Module directory structure for registry
my-module/
├── main.tf
├── variables.tf
├── outputs.tf
├── README.md
├── CHANGELOG.md
├── LICENSE
├── examples/
│   └── basic/
│       ├── main.tf
│       └── variables.tf
└── tests/
    └── module_test.go
```

## Registry Operations

### Publishing Requirements
- Module name: provider-module-name
- GitHub repository naming
- Documentation standards
- README format
- Variable documentation
- Output documentation
- Example configurations

### Versioning
- Semantic versioning (SemVer)
- Version tags
- Release management
- Changelog tracking
- Backward compatibility

### Documentation
- Module description
- Input variables
- Output values
- Examples and usage
- Dependencies
- Notes and warnings

### Private Registries
- Terraform Cloud
- GitLab modules
- GitHub packages
- Self-hosted registries
- Authentication

### Best Practices
- Document all variables
- Provide examples
- Use consistent naming
- Maintain changelog
- Test before release
- Semantic versioning
- Deprecation notices

## Registry Discovery

- Search functionality
- Module ratings
- Download counts
- Documentation quality
- Maintenance status
- Community feedback
