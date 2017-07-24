package input

import (
	"fmt"
	"strings"
)

/*
UserInput describes any type of user input, whether the value is a string, map,
or list.
*/
type UserInput interface {
	GetName() string
	GetDescription() string
	GetValue() interface{}
	SetValue(UserValue) error
}

/*
UserValue is any value supplied by the user. There is no implementation yet, but
I wanted to setup the groundwork for a future where there is. A wonderful f
uture.
*/
type UserValue interface{}

/*
StringInput contains information passed-to and received-from the user, with the
value of the variable being stored as a string.

* Name is the name of the variable and is used to reference it within templates.
* Description is a description of the variables intended use.
* Default is a default value.
* Value refers to the value of the Input, either the default of a value supplied
  by the user
*/
type StringInput struct {
	Name        string
	Description string
	Default     string
	Value       string
}

/*
CollectionFromStrings generates a new collection matching the type of the passed
variable, and adds an array of inputs to the Collection
*/
func CollectionFromStrings(c Collection, strs []*StringInput) Collection {
	for i := range strs {
		c.Add(strs[i])
	}

	return c
}

/*
GetName returns the variable name of the input.
*/
func (i *StringInput) GetName() string {

	return i.Name
}

/*
GetDescription returns the description of the string input.
*/
func (i *StringInput) GetDescription() string {

	return i.Description
}

/*
GetValue returns the value of the input. This is either the default value or
it's a value that has been set by user input. If there is no default, it's
empty, which in this case, it's an empty string.

Cast result as string
*/
func (i *StringInput) GetValue() interface{} {
	var v string
	if i.Value == "" {
		if i.Default != "" {
			v = i.Default
		} else {
			v = i.Default
		}
	} else {
		v = i.Value
	}
	return v
}

/*
SetValue assigns a user input to the StringInput. If the value supplied is
empty, and there is no default value, return the error.

This would be the place to perform validation on the user supplied value.
*/
func (i *StringInput) SetValue(val UserValue) error {
	var err error

	val = val.(string)
	if val != "" {
		i.Value = strings.TrimSpace(val.(string))
	} else if val == "" && i.Default == "" {
		err = fmt.Errorf("need to provide a value for %s", i.Name)
	}

	return err
}
