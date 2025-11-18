---
name: provider-testing
description: Master Terraform provider testing - unit tests, acceptance tests, and test automation. Learn testing strategies and coverage requirements.
---

# Provider Testing Mastery

## Quick Start - Acceptance Testing

```go
package provider

import (
	"testing"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccExampleThing_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckExampleThingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccExampleThingConfig_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckExampleThingExists("example_thing.test"),
					resource.TestCheckResourceAttr("example_thing.test", "name", "foo"),
				),
			},
		},
	})
}
```

## Testing Strategies

### Unit Testing
- Schema validation
- Input validation
- Function logic
- Error handling

### Acceptance Testing (ACC)
- Real API calls
- Create, Read, Update, Delete
- Import functionality
- Error scenarios
- Clean up on failure

### Test Framework
- Helper functions
- Test case structure
- Assertions
- Cleanup and teardown
- Parallel execution

### Coverage Requirements
- 100% of resources
- Data source coverage
- Error paths
- Edge cases
- Provider configuration

### CI/CD Integration
- GitHub Actions
- Test execution
- Coverage reporting
- Registry requirements
- Automated releases

## Testing Best Practices

- Comprehensive coverage
- Test edge cases
- Test error scenarios
- Clean up resources
- Parallel testing
- Idempotency testing
- Import testing
- Documentation of tests
