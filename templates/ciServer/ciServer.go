package simpleEmbedded

import (
	"github.com/objectpartners/continuous-deployment-templates/src/templates"
)

var template = &templates.Template{
	Name: "ciServer",
	Files: []*templates.TemplateFile{
		{
			Name:     "config.tf",
			Template: "templates/config.tf",
		},
		{
			Name:     "ci_server.tf",
			Template: "templates/ci_server.tf",
		},
	},
	Inputs: []*templates.UserInput{},
	DefaultInputs: []*templates.UserInput{
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
	},
}

// Adding a comment
func init() {
	templates.Add(template)
}
