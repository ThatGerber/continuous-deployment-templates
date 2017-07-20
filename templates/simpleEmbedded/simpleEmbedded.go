/*
Package simpleEmbedded installs a rancher server with an embedded DB.
*/
package simpleEmbedded

import (
	goTemplate "text/template"

	"github.com/objectpartners/continuous-deployment-templates/src/templates"
	"github.com/objectpartners/continuous-deployment-templates/src/templates/input"
)

type fileSystem struct{}

func (a *fileSystem) ReadFile(filename string) ([]byte, error) {

	return Asset(filename)
}

var files = []*templates.TemplateFile{
	{
		Name:     "deployment.tf",
		Template: "files/deployment.tf",
	},
	{
		Name:     "terraform.tfvars",
		Template: "files/terraform.tfvars",
	},
}

var inputs = []*input.StringInput{
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
}

// Adding a comment
func init() {
	templates.Add(&templates.Template{
		Name:   "simpleEmbedded",
		Files:  files,
		Engine: goTemplate.New("generate"),
		Assets: &fileSystem{},
		Inputs: input.CollectionFromStrings(inputs, &input.Array{}),
	})
}
