package test

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"strings"
	"testing"
)

func TestUnitExampleInfrastructure(t *testing.T) {
	t.Parallel()

	temporaryFolder := structure.CopyTerraformFolderToTemp(t, "../", "infrastructure/modules/s3-bucket")
	terraformPlanFilePath := filepath.Join(temporaryFolder, "plan.out")
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../infrastructure/modules/s3-bucket",
		Vars: map[string]interface{}{
			"s3_bucket_name": fmt.Sprintf("terratest-unit-%s", strings.ToLower(random.UniqueId())),
		},
		PlanFilePath: terraformPlanFilePath,
	})

	terraformPlanOutput := terraform.InitAndPlanAndShowWithStruct(t, terraformOptions)

	terraform.RequirePlannedValuesMapKeyExists(t, terraformPlanOutput, "aws_s3_bucket.this")
	assert.True(t, extractS3BucketVersioningEnabledAttribute(terraformPlanOutput))
}

func extractS3BucketVersioningEnabledAttribute(terraformPlanOutput *terraform.PlanStruct) bool {
	plannedS3Bucket := terraformPlanOutput.ResourcePlannedValuesMap["aws_s3_bucket.this"]
	plannedS3BucketVersioningConfiguration := plannedS3Bucket.AttributeValues["versioning"].([]interface{})[0].(map[string]interface{})
	plannedS3BucketVersioning := plannedS3BucketVersioningConfiguration["enabled"].(bool)
	return plannedS3BucketVersioning
}
