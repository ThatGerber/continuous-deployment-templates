/*
Package input_test
*/
package input_test

import (
	"reflect"
	"testing"

	"github.com/objectpartners/continuous-deployment-templates/src/templates/input"
)

var testInputs = []*input.UserInput{
	{
		Name:        "aVar",
		Description: "Description of aVar",
	},
	{
		Name:        "bVar",
		Description: "Description of bVar",
		Default:     "value b",
	},
	{
		Name:        "cVar",
		Description: "Description of CVar",
		Default:     "value c",
	},
	{
		Name:        "dVar",
		Description: "Description of dVar",
		Default:     "value d",
	},
}

// Generate empty array.
func newArray() *input.Array {

	return &input.Array{}
}

// Generate new array with user inputs. Set the count to 0 to get all items.
func newArrayWithItems(count int) *input.Array {
	var (
		a     *input.Array
		items []*input.UserInput
	)

	if count == 0 {
		count = len(testInputs)
	}

	for i := 0; i < count; i++ {
		items = append(items, testInputs[i])
	}

	a = newArray()
	a.Items = items

	return a
}

// Test Group for input.Array
func TestArray(t *testing.T) {
	t.Run("group", func(t *testing.T) {
		t.Run("Input::Cut", Input_Cut)
		t.Run("Array::All", Array_All)
		t.Run("Array::Get", Array_Get)
		t.Run("Array::Add", Array_Add)
		t.Run("Array::Update", Array_Update)
		t.Run("Array::Delete", Array_Delete)
		t.Run("Array::Hash", Array_Hash)
	})
}

// input.Cut
func Input_Cut(t *testing.T) {
	var inputs *input.Array

	inputs = newArrayWithItems(0)

	expectedCount := 2
	items := input.Cut(inputs, 1, expectedCount)

	if len(items) != expectedCount {
		t.Errorf("Array::Cut expected %d items, got %d", expectedCount, len(items))
	}

	bVar := items[0]
	if bVar.Name != "bVar" {
		t.Errorf("Array::Cut[0] expected var %s, got %s", "bVar", bVar.Name)
	}
	cVar := items[1]
	if cVar.Name != "cVar" {
		t.Errorf("Array::Cut[1] expected var %s, got %s", "cVar", bVar.Name)
	}

}

// Array_All
func Array_All(t *testing.T) {
	var array *input.Array

	array = newArrayWithItems(0)

	result := array.All()

	if !reflect.DeepEqual(testInputs, result) {
		errorFormat := "Array::All() expected %s, got %s."
		t.Errorf(errorFormat, testInputs, result)
	}

}

// Array_Get
func Array_Get(t *testing.T) {
	var actual, expected *input.UserInput

	array := newArrayWithItems(0)

	expected = testInputs[3]
	actual = array.Get("dVar")

	if !reflect.DeepEqual(expected, actual) {
		errorFormat := "Array::Get() expected %s, got %s."
		t.Errorf(errorFormat, expected, actual)
	}

	array = newArray()
	expected = nil
	actual = array.Get("dVar")

	if expected != actual {
		errorFormat := "Array::Get() expected %s, got %s."
		t.Errorf(errorFormat, expected, actual)
	}
}

// Array_Add
func Array_Add(t *testing.T) {
	var actual, expected []*input.UserInput

	array := newArrayWithItems(3)
	array.Add(testInputs[3])

	expected = testInputs

	actual = array.Items

	if !reflect.DeepEqual(expected, actual) {
		errorFormat := "Array::Add() expected %s, got %s."
		t.Errorf(errorFormat, expected, actual)
	}
}

// Array_Update
func Array_Update(t *testing.T) {
	var actual, expected *input.UserInput

	array := newArrayWithItems(0)

	expected = &input.UserInput{
		Name:        "dVar",
		Description: "Description of dVar",
		Default:     "value not d",
	}

	array.Update(expected)

	actual = array.Items[3]

	if !reflect.DeepEqual(expected, actual) {
		errorFormat := "Array::Update() expected %s, got %s."
		t.Errorf(errorFormat, expected, actual)
	}
}

// Array_Delete
func Array_Delete(t *testing.T) {
	var (
		actual, expected int
		err              error
	)

	array := newArrayWithItems(0)

	actual, _ = array.Delete("cVar")
	expected = 3

	if expected != actual {
		errorFormat := "Array::Delete() expected %s, got %s."
		t.Errorf(errorFormat, expected, actual)
	}

	actual, _ = array.Delete("bVar")
	expected = 2
	if expected != actual {
		errorFormat := "Array::Delete() expected %s, got %s."
		t.Errorf(errorFormat, expected, actual)
	}

	actual, err = array.Delete("notVar")
	if err == nil {
		t.Error("Array::Delete() expected err, got nothing.")
	}
	if expected != actual {
		errorFormat := "Array::Delete() expected %s, got %s."
		t.Errorf(errorFormat, expected, actual)
	}

	actual, _ = array.Delete("dVar")
	expected = 1
	if expected != actual {
		errorFormat := "Array::Delete() expected %s, got %s."
		t.Errorf(errorFormat, expected, actual)
	}

	actual, _ = array.Delete("aVar")
	expected = 0
	if expected != actual {
		errorFormat := "Array::Delete() expected %s, got %s."
		t.Errorf(errorFormat, expected, actual)
	}
}

// Array.Hash()
func Array_Hash(t *testing.T) {
	var actual, expected map[string]interface{}

	array := newArrayWithItems(0)

	expected = map[string]interface{}{}

	expected["aVar"] = ""
	expected["bVar"] = "value b"
	expected["cVar"] = "value c"
	expected["dVar"] = "value d"

	actual = array.Hash()

	if !reflect.DeepEqual(expected, actual) {
		errorFormat := "Array::Hash() expected %s, got %s."
		t.Errorf(errorFormat, expected, actual)
	}
}
