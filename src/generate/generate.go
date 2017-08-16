package generate

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/objectpartners/continuous-deployment-templates/src/templates"
	"github.com/objectpartners/continuous-deployment-templates/src/templates/input"

	"github.com/objectpartners/continuous-deployment-templates/templates/awsTerraform"
	"github.com/objectpartners/continuous-deployment-templates/templates/ciServer"
	"github.com/objectpartners/continuous-deployment-templates/templates/simpleEmbedded"
)

var TemplateCollection *templates.Set

// init
func init() {
	TemplateCollection = templates.NewSet()

	TemplateCollection.Add(&ciServer.Template)
	TemplateCollection.Add(&simpleEmbedded.Template)
	TemplateCollection.Add(&awsTerraform.Template)
}

/*
Templates as set.
*/
func Templates() *templates.Set {

	return TemplateCollection
}

/*
SetVars prompts user for input on all variables.

It looks through all templates and finds all user inputs.

If a duplicate input name is found, it will check if the original one has a
default value. If there is no default value, it replaces the input with the new
input.

If an input isn't already in the stack and doesn't have a default value it is
sent to the top of the list to be prompted first.
*/
func SetVars(names []string) {
	var uInputs = []*input.UserInput{}
	var skip bool

	for _, name := range names {
		t := Get(name)
		templateInputs := t.Inputs

		for _, val := range templateInputs.All() {
			skip = false

			for i, ui := range uInputs {
				if ui.Name == val.Name {
					if ui.GetValue() == "" {
						uInputs[i] = val
					}
					skip = true
				}
			}

			if !skip {
				if val.GetValue() == "" {
					uInputs = append([]*input.UserInput{val}, uInputs...)
				} else {
					uInputs = append(uInputs, val)
				}
			}
		}
	}

	var arr = &input.Array{Items: uInputs}

	arr.Map(input.PromptUser)

	TemplateCollection.Map(func(a *templates.Template) interface{} {
		a.Inputs = arr

		return a
	})
}

/*
Get a template by name.
*/
func Get(name string) *templates.Template {
	var template *templates.Template

	for _, template = range TemplateCollection.Templates {
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

	raw = TemplateCollection.Map(func(t *templates.Template) interface{} {

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
	sort.Strings(tmplts)

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
