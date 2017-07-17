package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/objectpartners/continuous-deployment-templates/src/templates"
	_ "github.com/objectpartners/continuous-deployment-templates/src/templates/all"
)

/*
Application Exit Code. Changed dynamically as the application is run depending
on errors encountered.
*/
var ExitCode int = 0

func main() {
	var output string
	var templateNames []string
	var tmpl *templates.Template
	var selected []string

	output = `
Select Template(s) to Generate

If selecting multiple templates, separate choices with comma.
E.g.: 0,2,3

Templates
------------------`
	fmt.Println(output)

	templateNames = templates.TemplateNames()
	sort.Strings(templateNames)
	for i, t := range templateNames {
		output = fmt.Sprintf("%d: %s", i, t)
		fmt.Println(output)
	}

	selected = templates.Select(templateNames)
	sort.Strings(selected)

	for _, name := range selected {
		tmpl = templates.Templates[name]

		output = fmt.Sprintf("Generating files for template: %s", tmpl.Name)
		fmt.Println(output)

		err := tmpl.RunTemplate()
		if err != nil {
			ExitCode = 1
			output = fmt.Sprintf("Error in : %s", err)
			fmt.Println(output)
		}
	}

	output = "Complete."
	fmt.Println(output)

	os.Exit(ExitCode)
}
