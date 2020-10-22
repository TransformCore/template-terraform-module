package test

import (
	"fmt"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformHelloWorldExample(t *testing.T) {
	t.Parallel()

	awsRegion := "eu-west-2"
	terraformDir := "../../examples/complete"

	test_structure.RunTestStage(t, "setup", func() {
		terraformOptions := &terraform.Options{
			TerraformDir: terraformDir,
			Upgrade:      true,
			Vars: map[string]interface{}{
				"region": awsRegion,
			},
			EnvVars: map[string]string{
				"AWS_DEFAULT_REGION": awsRegion,
			},
		}
		test_structure.SaveTerraformOptions(t, terraformDir, terraformOptions)

		terraform.InitAndApply(t, terraformOptions)
	})

	test_structure.RunTestStage(t, "validate", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)

		publicIp := terraform.Output(t, terraformOptions, "public_ip")

		url := fmt.Sprintf("http://%s:8080", publicIp)
		http_helper.HttpGetWithRetry(t, url, nil, 200, "Hello, ENGINE", 30, 5*time.Second)
	})

	test_structure.RunTestStage(t, "teardown", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
		terraform.Destroy(t, terraformOptions)
	})
}
