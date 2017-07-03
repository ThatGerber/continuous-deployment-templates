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
	templateNames := []string{}
	for key := range templates.Templates {
		templateNames = append(templateNames, key)
	}

	for i, t := range templateNames {
		fmt.Printf("%d: %s\n", i, t)
	}
	fmt.Print("Select template to render: ")
	selected := selectTemplate(len(templates.Templates))

	t := templates.Templates[templateNames[selected]]

	inputs := t.AcceptInputs()

	t.TemplateFiles("", inputs)
}

func selectTemplate(max int) int {
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(response)
	if response == "" {
		fmt.Println("Must select a template")
		return selectTemplate(max)
	}
	val, err := strconv.Atoi(response)
	if err != nil {
		fmt.Println("Invalid input")
		return selectTemplate(max)
	}
	if val >= max {
		fmt.Println("Invalid selection")
		return selectTemplate(max)
	}
	return val
}
