package templates

import (
	"fmt"
	ttemplate "text/template"
)

// Templates is the map of available templates by name
var Templates = map[string]*Template{}

/**
Template

A template holds a name, a set of inputs, as well as a set of files to be
parsed and interpolated (as TemplateFiles).

@TODO I believe Inputs is not being used for anything. Look into removing (or
      implementing).
*/
type Template struct {
	Name          string
	DefaultInputs []*UserInput
	Inputs        []*UserInput
	Files         []*TemplateFile
}

/**
Run through a Template struct and generate file files from templates.

Will return nothing if the files are written out, otherwise will return an
error if it runs into an irrecoverable error during the template file
generation process.
*/
func (t *Template) RunTemplate() error {
	inputs := t.AcceptInputs()

	err := t.TemplateFiles("", inputs)

	return err
}

/**
Create the final files from the "Files" attribute, from a slice of
TemplateFiles. Runs Files through a tex/template object to parse the template
file and writes them out to a destination.

This feels hacky. It's not a greatly intuitive API and feels like it has too
many responsibilities. It will need to be reworked, renamed, and deprecated in
the future.
*/
func (t *Template) TemplateFiles(destDir string, inputs map[string]string) error {
	var err error

	t.mergeDefaultInputs(inputs)

	engine := ttemplate.New("generate")

	for _, file := range t.Files {
		err = file.ConsumeTemplateFile(engine, inputs)
		if err != nil {

			return fmt.Errorf("In %s: %s", file.Name, err)
		}

		err = file.write(destDir)
		if err != nil {

			return fmt.Errorf("In %s: %s", file.Name, err)
		}
	}

	return nil
}

/**
Returns a slice of template names from the "Templates" slice.
*/
func TemplateNames() (tmplNames []string) {
	for key := range Templates {
		tmplNames = append(tmplNames, key)
	}

	return tmplNames
}
