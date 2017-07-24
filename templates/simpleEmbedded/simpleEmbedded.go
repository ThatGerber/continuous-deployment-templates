/*
Package simpleEmbedded installs a rancher server with an embedded DB.
*/
package simpleEmbedded

import (
	goTemplate "text/template"

	"github.com/objectpartners/continuous-deployment-templates/src/templates"
	"github.com/objectpartners/continuous-deployment-templates/src/templates/input"
)

/*
Template struct, carries template to be added and doesn't require any imports.
*/
var Template = templates.Template{
	Name:     "simpleEmbedded",
	ReadFile: func(f string) ([]byte, error) { return Asset(f) },
	Engine:   goTemplate.New("simpleEmbedded"),
	Files: []*templates.TemplateFile{
		{
			Name:     "deployment.tf",
			Template: "files/deployment.tf",
		},
		{
			Name:     "terraform.tfvars",
			Template: "files/terraform.tfvars",
		},
		{
			Name:     "deployment_inputs.tf",
			Template: "files/deployment_inputs.tf",
		},
		{
			Name:     "deployment_outputs.tf",
			Template: "files/deployment_outputs.tf",
		},
	},
	Inputs: input.CollectionFromStrings(&input.Array{}, []*input.StringInput{
		{
			Name:        "profile",
			Description: "AWS Credential Profile",
		},
		{
			Name:        "awsAccountId",
			Description: "AWS Account ID",
		},
		{
			Name:        "region",
			Default:     "us-west-2",
			Description: "AWS Region",
		},
		{
			Name:        "sshPublicKeyPath",
			Default:     "~/.ssh/id_rsa.pub",
			Description: "SSH Public Key Path",
		},
		{
			Name:        "environment",
			Default:     "tools",
			Description: "Environment name",
		},
		{
			Name:        "stack",
			Default:     "server",
			Description: "Stack name",
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
	}),
}
