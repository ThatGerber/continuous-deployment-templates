/*
Package templates_test
*/
package templates_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/objectpartners/continuous-deployment-templates/src/templates"
)

// Variables
var (
	testFnString string
	testFn       = func(t *templates.Template) interface{} {
		var blank interface{}

		testFnString = fmt.Sprintf("%s%s", testFnString, t.Name)

		blank = interface{}(t)

		return blank
	}
)

/*
Example data and Fixtures
*/
func newTemplateSet() *templates.Set {

	return &templates.Set{}
}

// Test templates.Set
func TestSet(t *testing.T) {
	t.Run("Group: templates.Set", func(t *testing.T) {
		t.Run("NewSet", NewSet)
		t.Run("Set::Add", setAdd)
		t.Run("Set::Map", setMap)
	})
}

// Test templates.NewSet() method.
func NewSet(t *testing.T) {
	var (
		templateSet *templates.Set
		newSet      *templates.Set
	)

	templateSet = newTemplateSet()
	newSet = templates.NewSet()

	if !reflect.DeepEqual(newSet, templateSet) {
		t.Errorf("Expected %T, got %T", templateSet, newSet)
	}
}

// Test templateSet.Add(t)
func setAdd(t *testing.T) {
	t.Skip("Not yet written.")
}

// Test templateSet.Map(fn)
func setMap(t *testing.T) {
	var (
		res         []interface{}
		templateSet *templates.Set
	)

	templateSet = newTemplateSet()

	res = templateSet.Map(testFn)

	newString := ""
	for i := range res {
		aVar := res[i].(templates.Template)
		newString = fmt.Sprintf("%s%s", newString, aVar.Name)
	}

	if newString != "" {
		t.Errorf("Expected %s, got %s", "", newString)
	}
}
