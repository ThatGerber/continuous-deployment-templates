package simpleEmbedded

import (
	"github.com/objectpartners/continuous-deployment-templates/templates"
)

//go:generate go run ../../scripts/templates.go

func init() {
	templates.Add(&templates.Template{
		Name:   "simpleEmbedded",
		Inputs: []*templates.UserInput{},
		Files: []*templates.TemplateFile{
			&templates.TemplateFile{
				Name:     "deployment.tf",
				Template: deploymentTf,
			},
			&templates.TemplateFile{
				Name:     "variables.tfvars",
				Template: variablesTfvars,
			},
		},
	})
}
