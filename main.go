package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/objectpartners/continuous-deployment-templates/templates"
	_ "github.com/objectpartners/continuous-deployment-templates/templates/all"
)

var defaultInputs = []*templates.UserInput{
	&templates.UserInput{
		Name:        "region",
		Default:     "us-west-2",
		Description: "AWS Region",
	},
	&templates.UserInput{
		Name:        "profile",
		Description: "AWS Credential Profile",
	},
	&templates.UserInput{
		Name:        "sshPublicKeyPath",
		Default:     "~/.ssh/id_rsa.pub",
		Description: "SSH Public Key Path",
	},
	&templates.UserInput{
		Name:        "environment",
		Default:     "tools",
		Description: "Environment Name",
	},
	&templates.UserInput{
		Name:        "stack",
		Default:     "server",
		Description: "Stack Name",
	},
	&templates.UserInput{
		Name:        "networkCidr",
		Default:     "10.0.0.0/16",
		Description: "Network CIDR",
	},
	&templates.UserInput{
		Name:        "numPublicSubnets",
		Default:     "3",
		Description: "Number of public subnets",
	},
	&templates.UserInput{
		Name:        "numPrivateSubnets",
		Default:     "3",
		Description: "Number of private subnets",
	},
}

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

	inputs := gatherInputs(defaultInputs, t.Inputs)

	t.TemplateFiles(inputs)
}

func gatherInputs(global, template []*templates.UserInput) map[string]string {
	variables := make(map[string]string)
	for _, input := range global {
		promptForInput(variables, input)
	}
	for _, input := range template {
		promptForInput(variables, input)
	}
	return variables
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

func promptForInput(data map[string]string, input *templates.UserInput) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s [%s]: ", input.Description, input.Default)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(response)
	if input.Default != "" && response == "" {
		data[input.Name] = input.Default
	} else if response != "" {
		data[input.Name] = strings.TrimSpace(response)
	} else {
		fmt.Println("Must provide a value")
		promptForInput(data, input)
	}
}
