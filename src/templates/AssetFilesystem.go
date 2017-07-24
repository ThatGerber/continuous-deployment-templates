package templates

/*
ReadFile is a function that reads files within the source package and returns a
byte slice containing contents of a package template.
*/
type ReadFile func(string) ([]byte, error)
