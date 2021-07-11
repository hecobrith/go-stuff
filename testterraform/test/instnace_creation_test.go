package test

import (
	"testing"
	"strings"
	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformGcpInstanceCPU(t *testing.T) {

	// website::tag::1:: Get the Project Id to use
	projectId := "reclutachat-production"

	// website::tag::2:: Give the example instance a unique name
	instanceName := fmt.Sprintf("gcp-other-instance-%s", strings.ToLower(random.UniqueId()))

	// website::tag::6:: Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// website::tag::3:: The path to where our Terraform code is located
		TerraformDir: "../modules/instance",

		// website::tag::4:: Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"instance_name": instanceName,
		},

		// website::tag::5:: Variables to pass to our Terraform code using TF_VAR_xxx environment variables
		EnvVars: map[string]string{
			"GOOGLE_CLOUD_PROJECT": projectId,
		},
	})

	// website::tag::8:: At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// website::tag::7:: Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)
	instance_cpu := terraform.Output(t, terraformOptions, "instance_cpu")
	assert.Equal(t, "Intel Haswell", instance_cpu)
}

func TestTerraformGcpInstanceName(t *testing.T) {

	// website::tag::1:: Get the Project Id to use
	projectId := "reclutachat-production"

	// website::tag::2:: Give the example instance a unique name
	instanceName := fmt.Sprintf("gcp-hello-world-example-%s", strings.ToLower(random.UniqueId()))

	// website::tag::6:: Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// website::tag::3:: The path to where our Terraform code is located
		TerraformDir: "../modules/instance",

		// website::tag::4:: Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"instance_name": instanceName,
		},

		// website::tag::5:: Variables to pass to our Terraform code using TF_VAR_xxx environment variables
		EnvVars: map[string]string{
			"GOOGLE_CLOUD_PROJECT": projectId,
		},
	})

	// website::tag::8:: At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// website::tag::7:: Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)
	instance_name := terraform.Output(t, terraformOptions, "instance_name")
	assert.Equal(t, instanceName, instance_name)
}