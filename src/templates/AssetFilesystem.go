package templates

/*
The AssetFilesystem would be an array of template file names and their paths.
In the case of this array, we're using the template file name to reference a
generated "asset" that contains a byte slice of the value of that template.
*/
type AssetFilesystem interface {
	ReadFile(string) ([]byte, error)
}
