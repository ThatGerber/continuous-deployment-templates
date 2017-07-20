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
func PromptUser(val UserInput) UserInput {
	var err error

	// 1. Prompt
	inputFormat := `
%s:
[%s]: `
	fmt.Printf(inputFormat, val.GetDescription(), val.GetValue())

	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(response)

	err = val.SetValue(response)
	if err != nil {
		fmt.Println(fmt.Sprintf("Err: %s", err))
		return PromptUser(val)
	}

	return val
}
