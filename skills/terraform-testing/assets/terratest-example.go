package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformVpc(t *testing.T) {
	opts := &terraform.Options{
		TerraformDir: "../modules/vpc",
	}
	defer terraform.Destroy(t, opts)
	terraform.InitAndApply(t, opts)
}
