---
name: provider-sdk
description: Master Terraform Provider SDK - build custom providers for any API or cloud platform. Learn resource implementation, schemas, and provider architecture.
---

# Provider SDK Development

## Quick Start

```go
package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("API_KEY", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"example_thing": resourceExampleThing(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"example_thing": dataSourceExampleThing(),
		},
	}
}

func resourceExampleThing() *schema.Resource {
	return &schema.Resource{
		Create: resourceExampleThingCreate,
		Read:   resourceExampleThingRead,
		Update: resourceExampleThingUpdate,
		Delete: resourceExampleThingDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
```

## Provider Development

### SDK Concepts
- Terraform Plugin SDK v2 (Go)
- Resource and data source types
- Schemas and attributes
- CRUD operations
- Type validation

### Resource Implementation
- Create operation
- Read operation
- Update operation
- Delete operation
- State management

### Schema Definition
- Argument types
- Default values
- Validation rules
- Computed attributes
- Sensitive data

### Provider Configuration
- Authentication
- Client initialization
- Connection pooling
- Retry logic
- Error handling

### SDK Features
- Auto-generated documentation
- State migration
- Timeouts and retries
- Nested resources
- Computed values

## Best Practices

- Comprehensive error handling
- Proper logging
- Resource import support
- Validation at schema level
- Documentation strings
- Backward compatibility
