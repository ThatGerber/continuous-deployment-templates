package ciServer

import (
	"fmt"
	"log"

	"github.com/objectpartners/continuous-deployment-templates/src/templates"
)

var template = &templates.Template{
	Name: "ciServer",
	Files: []*templates.TemplateFile{
		{
			Name:     "ci_inputs.tf",
			Template: "templates/ci_inputs.tf",
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
			Name:        "networkCidr",
			Default:     "10.0.0.0/16",
			Description: "CIDR of CI Server's Network",
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
	var err error

	for _, file := range template.Files {
		file.RawBody, err = Asset(file.Template)

		if err != nil {
			log.Fatal(fmt.Sprintf("Error on template %s => %s", file.Name, err))
		}
	}
	templates.Add(template)
}
