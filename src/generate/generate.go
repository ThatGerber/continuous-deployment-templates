package generate

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/objectpartners/continuous-deployment-templates/src/templates"
	"github.com/objectpartners/continuous-deployment-templates/templates/ciServer"
	"github.com/objectpartners/continuous-deployment-templates/templates/simpleEmbedded"
)

var set *templates.Set

func init() {
	set = templates.NewSet()

	set.Add(&ciServer.Template)
	set.Add(&simpleEmbedded.Template)
}

/*
Templates as set.
*/
func Templates() *templates.Set {

	return set
}

/*
Get a template by name.
*/
func Get(name string) *templates.Template {
	var template *templates.Template

	for _, template = range set.Templates {
		if template.Name == name {
			break
		}
	}

	return template
}

/*
TemplateNames is a slice of of the names of queued templates.
*/
func TemplateNames() []string {
	var raw []interface{}
	var tmplNames []string

	raw = set.Map(func(t *templates.Template) interface{} {

		return t.Name
	})

	for i := range raw {
		tmplNames = append(tmplNames, raw[i].(string))
	}

	return tmplNames
}

/*
Select requests a selection from the user, returning a slice containing the
names of templates to be parsed, otherwise keeps looping until it gets an
acceptable result.

User can submit choices as a single integer ("1"), or they can pass a comma-
separated string of choices ("1,3,4") to be parsed through and created.
*/
func Select() []string {
	var tmplts []string
	var selected []string

	tmplts = TemplateNames()

	max := len(tmplts)

	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(response)

	if response == "" {
		fmt.Println("Error - Must select a template.")
		return Select()
	}

	for _, key := range strings.Split(response, ",") {
		index, err := strconv.Atoi(strings.TrimSpace(key))
		if err != nil {
			fmt.Printf("Error - Invalid input: %s\n", key)
			return Select()
		}

		if index >= max {
			fmt.Printf("Error - Invalid Key Selection: %d\n", index)
			return Select()
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
