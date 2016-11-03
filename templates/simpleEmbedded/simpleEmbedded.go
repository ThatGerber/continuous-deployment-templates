package simpleEmbedded

import (
	"github.com/objectpartners/continuous-deployment-templates/templates"
)

//go:generate go run ../../scripts/templates.go

var template = &templates.Template{
	Name:   "simpleEmbedded",
	Inputs: []*templates.UserInput{},
	Files:  templateFiles,
}

func init() {
	templates.Add(template)
}
