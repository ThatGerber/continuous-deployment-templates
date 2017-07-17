package templates

// Add template to templates map.
func Add(template *Template) {
	Templates[template.Name] = template
}
