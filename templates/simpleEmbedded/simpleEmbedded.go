package simpleEmbedded

import (
	"github.com/objectpartners/continuous-deployment-templates/src/templates"
)

var template = &templates.Template{
	Name: "simpleEmbedded",
	Files: []*templates.TemplateFile{
		{
			Name:     "deployment.tf",
			Template: "files/deployment.tf",
		},
		{
			Name:     "terraform.tfvars",
			Template: "files/terraform.tfvars",
		},
	},
	Inputs: []*templates.UserInput{},
	DefaultInputs: []*templates.UserInput{
		{
			Name:        "region",
			Default:     "us-west-2",
			Description: "AWS Region",
		},
		{
			Name:        "profile",
			Description: "AWS Credential Profile",
		},
		{
			Name:        "aws_account_id",
			Description: "AWS Account ID",
		},
		{
			Name:        "sshPublicKeyPath",
			Default:     "~/.ssh/id_rsa.pub",
			Description: "SSH Public Key Path",
		},
		{
			Name:        "environment",
			Default:     "tools",
			Description: "Environment Name",
		},
		{
			Name:        "stack",
			Default:     "server",
			Description: "Stack Name",
		},
		{
			Name:        "networkCidr",
			Default:     "10.0.0.0/16",
			Description: "Network CIDR",
		},
		{
			Name:        "numPublicSubnets",
			Default:     "3",
			Description: "Number of public subnets",
		},
		{
			Name:        "numPrivateSubnets",
			Default:     "3",
			Description: "Number of private subnets",
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
