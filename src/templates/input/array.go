package input

/*
Collection is a grouping of user inputs.
*/
type Collection interface {
	VarHasher
	All() []UserInput
	Get(string) UserInput
	Add(UserInput)
	Update(UserInput)
	Delete(string) int
	Map(func(UserInput) UserInput)
}

/*
VarHasher is an interface for describing a struct that can communicate with a
template. The interface must at minimum produce a hash table containing key =>
value of some variables to be referenced within a templating engine, whatever
that is.

@TODO
In addition to a Hash(), it should also produce a series of functions that may
be used within the templating engine to retrieve more dynamic information.
*/
type VarHasher interface {
	Hash() map[string]interface{}
}

/*
Array is a collection of non-duplicated items that can be referenced by a key.
In this case, the Array contains a collection of UserInputs that are unique
based on their 'name'.
*/
type Array struct {
	Items []UserInput
}

/*
GetRange returns a range of UserInputs from the supplied collection.

If the start index is greater than the number of items (e.g. start of "5" and
there are only 5 items), it will return an empty slice.

If the the sum of the start index and and length are greater than the number of
items (e.g. start "3" and length "3" and there are only 5 items), then it will
return all of the items. Passing a negative length returns everything.
*/
func GetRange(c Collection, start int, length int) []UserInput {
	var list []UserInput
	var count int
	var end int

	list = c.All()
	count = len(list)
	if start >= count {
		return list
	}

	if length < 0 || (start+length) >= count {
		end = count
	} else {
		end = start + length
	}

	list = list[start:end]

	return list
}

/*
Get returns a single UserInput from a collection, referencing the variable name
of the user input based on a string key.
*/
func (c *Array) Get(key string) UserInput {
	if len(c.Items) < 1 {
		return nil
	}

	for i := range c.Items {
		if c.Items[i].GetName() == key {
			return c.Items[i]
		}
	}

	return nil
}

/*
All returns all of the items in the collection as a slice.
*/
func (c *Array) All() []UserInput {
	return c.Items
}

/*
Add will append additional UserInput to the collection.
*/
func (c *Array) Add(item UserInput) {
	var match bool

	for _, currentItem := range c.Items {
		match = false
		if item.GetName() == currentItem.GetName() {
			match = true
			if currentItem.GetValue() == "" {
				c.Update(item)
			}
		}
	}
	if !match {
		c.Items = append(c.Items, item)
	}
}

/*
Update item in list with contents of new item. Looks through list searches for
an input with a matching name. If found, replaces that item with input item.
*/
func (c *Array) Update(item UserInput) {
	for i, currentItem := range c.Items {
		if item.GetName() == currentItem.GetName() {
			c.Items[i] = item
		}
	}
}

/*
Map runs a supplied function against all items in the collection.

Map will edit in place, altering the inputs in the original collection.
*/
func (c *Array) Map(fn func(UserInput) UserInput) {
	for i, v := range c.Items {
		c.Items[i] = fn(v)
	}
}

/*
Delete user input from the collection by the variable name.

Returns the number of items in the list if succesful, nil if it doesn't find
anything to delete.

@TODO Write a test for this. Not sure what works.
*/
func (c *Array) Delete(key string) int {
	var count int

	for i := range c.Items {
		if c.Items[i].GetName() == key {
			copy(c.Items[i:], c.Items[i+1:])
			c.Items[len(c.Items)-1] = nil // or the zero value of T
			c.Items = c.Items[:len(c.Items)-1]
			count = len(c.Items)
		}
	}

	return count
}

/*
Hash generates a key => value hash map of the collection. Useful for preparing
the inputs to be sent to a template.
*/
func (c *Array) Hash() map[string]interface{} {
	var h map[string]interface{}

	h = make(map[string]interface{})

	for i := range c.Items {
		key := c.Items[i].GetName()
		val := c.Items[i].GetValue()
		h[key] = (interface{})(val)
	}

	return h
}
