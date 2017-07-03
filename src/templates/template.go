package templates

import (
	"fmt"
	"os"
	"path/filepath"
	ttemplate "text/template"
)

// Templates is the map of available templates by name
var Templates = map[string]*Template{}

type Template struct {
	Name          string
	DefaultInputs []*UserInput
	Inputs        []*UserInput
	Files         []*TemplateFile
}

func (t *Template) TemplateFiles(output string, inputs map[string]string) error {
	t.mergeDefaultInputs(inputs)
	engine := ttemplate.New("infrastructure")
	for _, file := range t.Files {

		dest, err := os.Create(filepath.Join(output, file.Name))
		if err != nil {
			fmt.Println("Error creating output file: " + file.Name)
			os.Exit(1)
		}

		tmpl, err := engine.Parse(file.Body)
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
	for _, i := range t.DefaultInputs {
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
	for _, input := range t.DefaultInputs {
		promptForInput(variables, input)
	}
	for _, input := range t.Inputs {
		promptForInput(variables, input)
	}
	return variables
}

func (t *Template) ValidateInputs(inputs map[string]string) error {
	for _, i := range t.DefaultInputs {
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
