package test

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestIntegrationExampleInfrastructure(t *testing.T) {
	awsRegion := "us-east-1"
	s3BucketName := fmt.Sprintf("terratest-unit-%s", strings.ToLower(random.UniqueId()))
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../infrastructure/example",
		Vars: map[string]interface{}{
			"s3_bucket_name": s3BucketName,
		},
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	require.NotEmpty(t, terraform.Output(t, terraformOptions, "s3_bucket_arn"))
	require.Equal(t, fmt.Sprintf("%s.s3.amazonaws.com", s3BucketName), terraform.Output(t, terraformOptions, "s3_bucket_domain_name"))
	aws.AssertS3BucketExists(t, awsRegion, s3BucketName)
	aws.AssertS3BucketVersioningExists(t, awsRegion, s3BucketName)
}
