package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/objectpartners/continuous-deployment-templates/src/generate"
)

var buildID = "undefined"

/*
ExitCode is the application's current ExitCode during processing. Changed
dynamically as the application is run depending on errors encountered.
*/
var ExitCode int

/*
generate

@TODO Template collection should gather inputs from all templates and merge.
*/
func main() {
	var output string
	var templateNames []string
	var selected []string

	output = `
Generate - Build %s

Select Template(s) to Generate:
If selecting multiple templates, separate choices with comma. E.g.: 0,2,3

Templates
------------------`
	fmt.Println(fmt.Sprintf(output, buildID))

	templateNames = generate.TemplateNames()
	sort.Strings(templateNames)
	for i, t := range templateNames {
		output = fmt.Sprintf("%d: %s", i, t)
		fmt.Println(output)
	}

	selected = generate.Select()
	sort.Strings(selected)

	for _, name := range selected {
		tmpl := generate.Get(name)

		output = fmt.Sprintf("Generating files for template: %s", tmpl.Name)
		fmt.Println(output)

		err := tmpl.Run()
		if err != nil {
			ExitCode = 1
			output = fmt.Sprintf("Fatal Error:\n%s", err)
			fmt.Println(output)
		}
	}

	output = "Complete."
	fmt.Println(output)

	os.Exit(ExitCode)
}
