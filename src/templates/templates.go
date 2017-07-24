package templates

/*
Set is a collection of templates to be munged by the application.
*/
type Set struct {
	Templates []*Template
}

var templates *Set

/*
NewSet of Templates
*/
func NewSet() *Set {
	s := &Set{}

	return s
}

/*
Add template to Collection

The templates.Template map contains the name and value of any generated
templates.
*/
func (c *Set) Add(t *Template) {

	c.Templates = append(c.Templates, t)
}

/*
Map a function to a dataset.
*/
func (c *Set) Map(fn func(*Template) interface{}) []interface{} {
	var face []interface{}
	for _, t := range c.Templates {
		face = append(face, fn(t))
	}

	return face
}
