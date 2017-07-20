/*
Package ciServer generates configuration for a Continuous Integration server.
*/
package ciServer

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
		Name:     "ci_inputs.tf",
		Template: "templates/ci_inputs.tf",
	},
	{
		Name:     "ci_server.tf",
		Template: "templates/ci_server.tf",
	},
}

var inputs = []*input.StringInput{
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
}

// Adding a comment
func init() {
	templates.Add(&templates.Template{
		Name:   "ciServer",
		Files:  files,
		Engine: goTemplate.New("generate"),
		Assets: &fileSystem{},
		Inputs: input.CollectionFromStrings(inputs, &input.Array{}),
	})
}
