package simpleEmbedded

import (
	"fmt"
	"log"

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
	var err error

	for _, file := range template.Files {
		file.RawBody, err = Asset(file.Template)

		if err != nil {
			log.Fatal(fmt.Sprintf("Error on template %s => %s", file.Name, err))
		}
	}
	templates.Add(template)
}
