package templates

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

// Templates is the map of available templates by name
var Templates = map[string]*Template{}

var defaultInputs = []*UserInput{
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

// Add adds an available template
func Add(template *Template) {
	Templates[template.Name] = template
}

type Template struct {
	Name   string
	Inputs []*UserInput
	Files  []*TemplateFile
}

type TemplateFile struct {
	Name     string
	Template string
}

func (t *Template) TemplateFiles(output string, inputs map[string]string) error {
	t.mergeDefaultInputs(inputs)
	engine := template.New("infrastructure")
	for _, file := range t.Files {
		dest, err := os.Create(filepath.Join(output, file.Name))
		if err != nil {
			fmt.Println("Error creating output file: " + file.Name)
			os.Exit(1)
		}
		tmpl, err := engine.Parse(file.Template)
		if err != nil {
			fmt.Println("Error parsing data from: " + file.Name)
			os.Exit(1)
		}
		err = tmpl.Execute(dest, &Input{
			Variables: inputs,
		})
		if err != nil {
			fmt.Println("Error executing templating file: " + file.Name)
			os.Exit(1)
		}
	}
	return nil
}

func (t *Template) mergeDefaultInputs(inputs map[string]string) {
	for _, i := range defaultInputs {
		if i.Default != "" {
			if _, ok := inputs[i.Name]; !ok {
				inputs[i.Name] = i.Default
			}
		}
	}
	for _, i := range t.Inputs {
		if i.Default != "" {
			if _, ok := inputs[i.Name]; !ok {
				inputs[i.Name] = i.Default
			}
		}
	}
}

func (t *Template) AcceptInputs() map[string]string {
	variables := make(map[string]string)
	for _, input := range defaultInputs {
		promptForInput(variables, input)
	}
	for _, input := range t.Inputs {
		promptForInput(variables, input)
	}
	return variables
}

func (t *Template) ValidateInputs(inputs map[string]string) error {
	for _, i := range defaultInputs {
		if i.Default == "" {
			if _, ok := inputs[i.Name]; !ok {
				return fmt.Errorf("Required input '%s' missing", i.Name)
			}
		}
	}
	for _, i := range t.Inputs {
		if i.Default == "" {
			if _, ok := inputs[i.Name]; !ok {
				return fmt.Errorf("Required input '%s' missing", i.Name)
			}
		}
	}
	return nil
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
