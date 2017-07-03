package simpleEmbedded

import (
	"github.com/objectpartners/continuous-deployment-templates/src/templates"
)

var template = &templates.Template{
	Name: "ciServer",
	Files: []*templates.TemplateFile{
		{
			Name:     "development.tf",
			Template: "templates/devops_vpc.tf",
		},
		{
			Name:     "staging.tf",
			Template: "templates/stage_vpc.tf",
		},
		{
			Name:     "production.tf",
			Template: "templates/prod_vpc.tf",
		},
	},
	Inputs: []*templates.UserInput{},
	DefaultInputs: []*templates.UserInput{
		{
			Name:        "environment",
			Description: "Environment Name (alnum only)",
		},
		{
			Name:        "networkCidr",
			Default:     "10.0",
			Description: "First two bytes of CIDR",
		},
		{
			Name:        "ciType",
			Default:     "jenkins",
			Description: "Which CI server would you like to use? [Jenkins, Drone, Concourse]",
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
