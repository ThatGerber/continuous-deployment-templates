package awsTerraform

import (
	goTemplate "text/template"

	"github.com/objectpartners/continuous-deployment-templates/src/templates"
	"github.com/objectpartners/continuous-deployment-templates/src/templates/input"
)

/*
Template ...
*/
var Template = templates.Template{
	Name:     "awsTerraform",
	ReadFile: func(f string) ([]byte, error) { return Asset(f) },
	Engine:   goTemplate.New("awsTerraform"),
	Files: []*templates.TemplateFile{
		{
			Name:     "backend.tfvars",
			Template: "templates/terraform.tfvars",
		},
		{
			Name:     "config.tf",
			Template: "templates/config.tf",
		},
		{
			Name:     "ci_server.tf",
			Template: "templates/ci_server.tf",
		},
		{
			Name:     "Makefile",
			Template: "templates/Makefile",
		},
		{
			Name:     "terraform.tf",
			Template: "templates/terraform.tf",
		},
	},
	Inputs: input.NewArray([]*input.UserInput{
		{
			Name:        "environment",
			Description: "Environment Name (alnum only)",
		},
		{
			Name:        "awsRegion",
			Description: "AWS Region",
			Default:     "us-west-2",
		},
		{
			Name:        "awsProfile",
			Description: "AWS IAM Profile",
		},
		{
			Name:        "networkCidr",
			Default:     "10.0.0.0/16",
			Description: "CIDR of CI Server's Network",
		},
		{
			Name:        "tfBackend",
			Default:     "local",
			Description: "Type of Terraform Backend to Initiate. [local, s3]",
		},
		{
			Name:        "tfStateBucket",
			Description: "Name of Terraform State Bucket and Key",
		},
		{
			Name:        "tfStateRegion",
			Default:     "us-west-2",
			Description: "Region to place Terraform State Bucket",
		},
		{
			Name:        "ciType",
			Default:     "jenkins",
			Description: "Which CI server would you like to use? [jenkins, drone, concourse]",
		},
		{
			Name:        "moduleSource",
			Default:     "github.com/objectpartners/continuous-deployment-templates",
			Description: "Location of module templates",
		},
	}),
}
