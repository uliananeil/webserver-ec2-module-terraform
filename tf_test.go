package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAws(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "infrastructure/",
	})

	terraform.InitAndApply(t, terraformOptions)

	http_id_ip := terraform.OutputMap(t, terraformOptions, "http_ip")
	db_id_ip := terraform.OutputMap(t, terraformOptions, "db_ip")
	// Check if instances exist
	for _, ip := range http_id_ip {
		assert.NotEmpty(t, ip)
	}
	for id, _ := range db_id_ip {
		assert.NotEmpty(t, id)
		publicIP := aws.GetPublicIpOfEc2Instance(t, id, "eu-central-1")
		assert.Empty(t, publicIP)
	}

	vpc_cidr_block := terraform.Output(t, terraformOptions, "vpc_cidr_block")
	assert.Equal(t, "192.168.0.0/16", vpc_cidr_block)

	http_subnet_cidr_block := terraform.Output(t, terraformOptions, "http_subnet_cidr_block")
	assert.Equal(t, "192.168.1.0/24", http_subnet_cidr_block)

	db_subnet_cidr_block := terraform.Output(t, terraformOptions, "db_subnet_cidr_block")
	assert.Equal(t, "192.168.2.0/24", db_subnet_cidr_block)

}
