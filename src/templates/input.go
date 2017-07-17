package templates

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	Variables map[string]string
}

/**
Send user prompt for template choices. Returns a slice containing the names of
templates to be parsed, otherwise keeps looping until it gets an acceptable
result.

User can submit choices as a single integer ("1"), or they can pass a comma-
separated string of choices ("1,3,4") to be parsed through and created.
*/
func Select(tmplts []string) []string {
	var max int
	var selected []string

	max = len(tmplts)

	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(response)

	if response == "" {
		fmt.Println("Error - Must select a template.")

		return Select(tmplts)
	}

	for _, key := range strings.Split(response, ",") {

		index, err := strconv.Atoi(strings.TrimSpace(key))
		if err != nil {
			fmt.Printf("Error - Invalid input: %s\n", key)

			return Select(tmplts)
		}

		if index >= max {
			fmt.Printf("Error - Invalid Key Selection: %d\n", index)

			return Select(tmplts)
		}

		selected = append(selected, tmplts[index])
	}
	output := `
Templates Being Generated:
%s
`
	fmt.Println(fmt.Sprintf(output, selected))
	return selected
}

// Prompt user for input. If a value is not passed, set to the default value.
func promptForInput(data map[string]string, input *UserInput) {
	var inputFormat string
	inputFormat = `%s [%s]:
> `

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(inputFormat, input.Description, input.Default)
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

/**
Merge a TemplateFile's default inputs with the Inputs slice.
*/
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

/**
Create a series of prompts for user input based on the DefaultInputs and
UserInput slices.
*/
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

/**
Validate the supplied default inputs as well as user supplied inputs.
*/
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
