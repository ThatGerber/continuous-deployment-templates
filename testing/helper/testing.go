package helper

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/logging"
	"github.com/objectpartners/continuous-deployment-templates/templates"
)

func Test(t *testing.T, spec TestCase) {

	logWriter, err := logging.LogOutput()
	if err != nil {
		t.Error(fmt.Errorf("error setting up logging: %s", err))
	}
	log.SetOutput(logWriter)

	if spec.Precheck != nil {
		err = spec.Precheck()
		if err != nil {
			t.Fatal(fmt.Errorf("error setting up spec: %s", err))
		}
	}
	dir, err := ioutil.TempDir("", spec.Template.Name)
	if err != nil {
		t.Fatal(fmt.Errorf("error creating temp dir for %s: %s", spec.Template.Name, err))
	}
	spec.dir = dir
	t.Logf("outputting to " + dir)

	err = spec.validateInput()
	if err != nil {
		t.Fatal(fmt.Errorf("error providing input to template %s: %s", spec.Template.Name, err))
	}
	err = spec.generate()
	if err != nil {
		t.Fatal(fmt.Errorf("error generating templates: %s", err))
	}

	//TODO run terraform

	err = spec.Check(&spec)
	if err != nil {
		t.Fatal(fmt.Errorf("error: %s", err))
	}

	//TODO destroy resources
	err = os.RemoveAll(dir)
	if err != nil {
		t.Fatal(fmt.Errorf("error cleaning up temp dir %s: %s", dir, err))
	}
}

type TestCase struct {
	Precheck PreCheckFunc
	Template *templates.Template
	Inputs   map[string]string
	Check    TestCheckFunc
	dir      string
}

func (tc *TestCase) validateInput() error {
	return tc.Template.ValidateInputs(tc.Inputs)
}

func (tc *TestCase) generate() error {
	err := tc.Template.TemplateFiles(tc.dir, tc.Inputs)
	if err != nil {
		return err
	}
	return nil
}

func AssertOutputFileContentsEqual(name, content string) TestCheckFunc {
	return func(tc *TestCase) error {
		data, err := ioutil.ReadFile(filepath.Join(tc.dir, name))
		if err != nil {
			return fmt.Errorf("Error loading file %s: %s", filepath.Join(tc.dir, name), err)
		}
		if strings.Compare(content, string(data)) != 0 {
			return fmt.Errorf("File %s does not match.\nExpected:\n%s\n\nActual:\n%s", filepath.Join(tc.dir, name), content, string(data))
		}
		return nil
	}
}

type PreCheckFunc func() error

type TestCheckFunc func(tc *TestCase) error

// ComposeTestCheckFunc lets you compose multiple TestCheckFuncs into
// a single TestCheckFunc.
//
// As a user testing their provider, this lets you decompose your checks
// into smaller pieces more easily.
func ComposeTestCheckFunc(fs ...TestCheckFunc) TestCheckFunc {
	return func(tc *TestCase) error {
		for i, f := range fs {
			if err := f(tc); err != nil {
				return fmt.Errorf("Check %d/%d error: %s", i+1, len(fs), err)
			}
		}

		return nil
	}
}
