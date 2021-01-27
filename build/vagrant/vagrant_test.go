package testing

import (
	"fmt"
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestSpelVagrant(t *testing.T) {
	resourceName := random.UniqueId()

	terraformOptions := &terraform.Options{
		// The path to where your Terraform code is located
		TerraformDir: "./infra",

		// Disable color output
		NoColor: true,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"resource_name":          fmt.Sprintf("packer-spel-vagrant-%s", resourceName),
			"source_cidr":            os.Getenv("SOURCE_CIDR"),
			"aws_region":             os.Getenv("AWS_REGION"),
			"ssm_vagrantcloud_token": os.Getenv("SSM_VAGRANTCLOUD_TOKEN"),
			"vagrantcloud_user":      os.Getenv("VAGRANTCLOUD_USER"),
			"spel_identifier":        os.Getenv("SPEL_IDENTIFIER"),
			"spel_version":           os.Getenv("SPEL_VERSION"),
			"spel_ci":                os.Getenv("SPEL_CI"),
			"packer_url":             os.Getenv("PACKER_URL"),
			"artifact_location":      os.Getenv("ARTIFACT_LOCATION"),
			"kms_key":                os.Getenv("KMS_KEY"),
			"code_repo":              os.Getenv("CODEBUILD_SOURCE_REPO_URL"),
			"source_commit":          os.Getenv("CODEBUILD_SOURCE_VERSION"),
			"iso_url_centos7":        os.Getenv("ISO_URL_CENTOS7"),
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)
}
