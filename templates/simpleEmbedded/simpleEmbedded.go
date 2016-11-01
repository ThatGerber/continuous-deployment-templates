package simpleEmbedded

import (
	"os"

	"github.com/objectpartners/continuous-deployment-templates/templates"
	. "github.com/objectpartners/continuous-deployment-templates/utils"
)

//go:generate go run ../../scripts/templates.go

func init() {
	templates.Add(&templates.Template{
		Name:   "simpleEmbedded",
		Inputs: []*templates.UserInput{},
		Files: func() []*os.File {
			return LoadFiles([]string{
				"deployment.tf",
				"variables.tfvars",
			})
		},
	})
}
