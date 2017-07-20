/*
Package templates manages the application's template files.
*/
package templates

import (
	"fmt"
	"io/ioutil"
	"os"
	goTemplate "text/template"

	"github.com/objectpartners/continuous-deployment-templates/src/templates/input"
)

// Templates is the map of available templates by name
var Templates = map[string]*Template{}

/*
Template holds a name, a set of inputs, as well as a set of files to be
parsed and interpolated (as TemplateFiles).
*/
type Template struct {
	Name   string
	Inputs input.Collection
	Files  []*TemplateFile
	Assets AssetFilesystem
	Engine *goTemplate.Template
}

/*
Run through a Template struct and generate file files from templates.

Will return nothing if the files are written out, otherwise will return an error
if it runs into an irrecoverable error during the template file generation
process.
*/
func (t *Template) Run() error {
	var err error
	var file *TemplateFile

	// Get Inputs
	t.Inputs.Map(input.PromptUser)

	// Generate Files from Templates
	for _, file = range t.Files {
		err = t.Generate(file)
		if err != nil {
			err = fmt.Errorf("In %s: %s", file.Name, err)
			return err
		}
	}

	return err
}

/*
Generate the final template from a TemplateFile. Combines source file and
template into a single HCL file.
*/
func (t *Template) Generate(file *TemplateFile) error {
	var err error
	var srcs []string
	var tmpf string

	srcs = append(srcs, file.DestAbsPath())

	tmpf, err = t.generateTempFile(file)
	defer os.Remove(tmpf)
	if err != nil {
		return err
	}

	srcs = append(srcs, tmpf)

	err = file.Amalgamate(srcs)
	if err != nil {
		return err
	}

	err = file.Write("")

	return err
}

/*
TemplateNames is a slice of of the names of queued templates.
*/
func TemplateNames() (tmplNames []string) {
	for key := range Templates {
		tmplNames = append(tmplNames, key)
	}
	return
}

/*
generateTempFile wil write a TemplateFile to a file within the system's temp
directory, returning the name of the temporary file path.
*/
func (t *Template) generateTempFile(file *TemplateFile) (string, error) {
	var err error
	var buf []byte
	var filename string

	buf, err = t.Assets.ReadFile(file.Template)

	if err != nil {
		err = fmt.Errorf("error reading bytes from %s", file.Template)
		return filename, err
	}

	file.Body = string(buf)
	t.Engine, err = t.Engine.Parse(file.Body)
	if err != nil {
		err = fmt.Errorf("error parsing data from %s", file.Name)
		return filename, err
	}

	tmpFile, err := ioutil.TempFile("", "")
	if err != nil {
		err = fmt.Errorf("error creating tmp output file %s", file.Name)
		return filename, err
	}
	hash := t.Inputs.Hash()
	err = t.Engine.Execute(tmpFile, hash)
	if err != nil {
		err = fmt.Errorf("error executing templating file %s", t.Name)
		return filename, err
	}

	filename = tmpFile.Name()

	return filename, err
}
