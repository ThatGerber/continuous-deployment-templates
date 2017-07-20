package templates

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Add template to templates map

The templates.Template map contains the name and value of any generated
templates.
*/
func Add(template *Template) {
	Templates[template.Name] = template
}

/*
Select requests a selection from the user, returning a slice containing the
names of templates to be parsed, otherwise keeps looping until it gets an
acceptable result.

User can submit choices as a single integer ("1"), or they can pass a comma-
separated string of choices ("1,3,4") to be parsed through and created.
*/
func Select(tmplts []string) []string {
	var selected []string

	max := len(tmplts)

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
Templates To Be Generated:
%s
`
	fmt.Println(fmt.Sprintf(output, selected))
	return selected
}
