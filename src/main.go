package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/objectpartners/continuous-deployment-templates/src/templates"
	_ "github.com/objectpartners/continuous-deployment-templates/src/templates/all"
)

func main() {
	var templateNames []string
	var tmpl *templates.Template
	var selectedTemplates []string
	var exitCode int

	exitCode = 0

	templateNames = templates.TemplateNames()

	for i, t := range templateNames {
		fmt.Printf("%d: %s\n", i, t)
	}
	fmt.Println("Select template(s) to render:")
	fmt.Println("If selecting multiple templates, separate choices with comma.")
	fmt.Print(": ")

	selectedTemplates = selectTemplates(templateNames)

	for _, name := range selectedTemplates {
		tmpl = templates.Templates[name]

		fmt.Println(fmt.Sprintf("Generating files for %v template.", string(tmpl.Name)))

		err := tmpl.RunTemplate()

		if err != nil {
			fmt.Print(err)
			exitCode = 1
		}
	}

	os.Exit(exitCode)
}

/**
Send user prompt for template choices. Returns a slice containing the names of
templates to be parsed, otherwise keeps looping until it gets an acceptable
result.

User can submit choices as a single integer ("1"), or they can pass a comma-
separated string of choices ("1,3,4") to be parsed through and created.
*/
func selectTemplates(tmplts []string) (selected []string) {
	var max int

	max = len(tmplts)

	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(response)

	if response == "" {
		fmt.Println("Must select a template")

		return selectTemplates(tmplts)
	}

	for _, key := range strings.Split(response, ",") {

		index, err := strconv.Atoi(strings.TrimSpace(key))
		if err != nil {
			fmt.Printf("Invalid input: %s\n", key)

			return selectTemplates(tmplts)
		}

		if index >= max {
			fmt.Printf("Invalid selection for key %s: %d\n", key, index)

			return selectTemplates(tmplts)
		}
		selected = append(selected, tmplts[index])
	}

	return selected
}
