/*
Package input contains a collection of structs and interfaces for managing user
input, including default values and values supplied by the user.
*/
package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
PromptUser sends the user a prompt on the console and stores the response.

Takes a UserInput and generates a prompt to the user. The result is passed to
set. If there is an error, it will print the error and return the prompt.
*/
func PromptUser(val *UserInput) *UserInput {
	// 1. Prompt
	inputFormat := `
%s:
[%s]: `
	fmt.Printf(inputFormat, val.Description, val.GetValue())

	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(response)

	val.Value = response

	if val.GetValue() == "" {
		fmt.Println(fmt.Sprintf("Err: %s", val.Value))
		return PromptUser(val)
	}

	return val
}

/*
UserInput describes any type of user input, whether the value is a string, map,
or list.
*/
type UserInput struct {
	Name        string
	Description string
	Default     string
	Value       string
}

/*
GetValue looks through the user supplied value and the default to determine what
the returned value should be. Ideally, all values except "Name" would be changed
to methods so that they could be interfaced.
*/
func (u *UserInput) GetValue() string {
	if u.Value == "" {
		return u.Default
	} else {
		return u.Value
	}
}

/*
NewArray generates a new collection matching the type of the passed variable,
and adds an array of inputs to the Collection.
*/
func NewArray(inputs []*UserInput) *Array {
	a := &Array{}

	for i := range inputs {
		a.Add(inputs[i])
	}

	return a
}
