package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	templates := loadTemplates()
	for i, template := range templates {
		fmt.Printf("%d: %s\n", i, template)
	}
	fmt.Print("Select template to render: ")
	selected := selectTemplate(len(templates))

	templateDir := "templates/" + templates[selected]
	sourceFiles, err := ioutil.ReadDir(templateDir)
	check(err)

	inputs := []*UserInput{
		&UserInput{
			Name:        "region",
			Default:     "us-west-2",
			Description: "AWS Region",
		},
		&UserInput{
			Name:        "profile",
			Description: "AWS Credential Profile",
		},
		&UserInput{
			Name:        "sshPublicKeyPath",
			Default:     "~/.ssh/id_rsa.pub",
			Description: "SSH Public Key Path",
		},
		&UserInput{
			Name:        "environment",
			Default:     "tools",
			Description: "Environment Name",
		},
		&UserInput{
			Name:        "stack",
			Default:     "server",
			Description: "Stack Name",
		},
		&UserInput{
			Name:        "networkCidr",
			Default:     "10.0.0.0/16",
			Description: "Network CIDR",
		},
		&UserInput{
			Name:        "numPublicSubnets",
			Default:     "3",
			Description: "Number of public subnets",
		},
		&UserInput{
			Name:        "numPrivateSubnets",
			Default:     "3",
			Description: "Number of private subnets",
		},
	}
	variables := make(map[string]string)
	for _, input := range inputs {
		promptForInput(variables, input)
	}
	templateFiles(templateDir, sourceFiles, variables)
}

func templateFiles(templateDir string, files []os.FileInfo, vars map[string]string) {
	t := template.New("infrastructure")
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fmt.Println("Name: " + file.Name())
		data, err := ioutil.ReadFile(templateDir + "/" + file.Name())
		check(err)
		dest, err := os.Create(file.Name())
		check(err)
		tmpl, err := t.Parse(string(data))
		check(err)
		err = tmpl.Execute(dest, &Input{
			Variables: vars,
		})
		check(err)
	}
}

func check(err error) {
	if err != nil {
		os.Exit(1)
	}
}

func loadTemplates() []string {
	var templates []string
	files, _ := ioutil.ReadDir("templates")
	for _, file := range files {
		if file.IsDir() {
			templates = append(templates, file.Name())
		}
	}
	return templates
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
	check(err)
	if val >= max {
		fmt.Println("Invalid selection")
		return selectTemplate(max)
	}
	return val
}

func promptForInput(data map[string]string, input *UserInput) {
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

type UserInput struct {
	Name        string
	Default     string
	Description string
}

type Input struct {
	Variables map[string]string
}
