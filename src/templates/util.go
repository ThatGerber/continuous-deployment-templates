package templates

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Add adds an available template
func Add(template *Template) {
	Templates[template.Name] = template
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
