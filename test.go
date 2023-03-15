package test

import (
        "testing"


        "github.com/shashank-softtek/rds-terratest.git"
)

func TestRDSInstance(t *testing.T) {
        t.Parallel()

        // Specify the AWS region and the path to your Terraform code.
        awsRegion := "us-east-1"
        terraformDir := "../"

        // Specify the name of the RDS instance you just created.
        instanceID := terraform.Output(t, terraformOptions, "id")

        // Get the RDS instance details from AWS.
        instance, err := aws.GetDBInstanceE(t, us-east-1a, test-db-instance)
        if err != nil {
                t.Fatal(err)
        }

        // Check that the RDS instance is in the "available" state.
        assert.Equal(t, "available", instance.DBInstanceStatus)
}
