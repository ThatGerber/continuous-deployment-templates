package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Running generator")
	cwd, _ := os.Getwd()
	dir := filepath.Base(cwd)
	templateName := dir
	fmt.Println("Name: " + templateName)
	out, _ := os.Create("templates.go")
	out.Write([]byte("package " + templateName + "\n\nimport \"github.com/objectpartners/continuous-deployment-templates/templates\"\n\nconst (\n"))
	fs, _ := ioutil.ReadDir(".")
	tmpls := make(map[string]string)
	for _, f := range fs {
		if strings.HasSuffix(f.Name(), ".tmpl") {
			fileName := strings.TrimSuffix(f.Name(), ".tmpl")
			tokens := strings.Split(fileName, ".")
			constName := tokens[0] + strings.Title(tokens[1])
			tmpls[fileName] = constName
			fdata, _ := os.Open(f.Name())
			out.Write([]byte("\t" + constName + " = `"))
			io.Copy(out, fdata)
			out.Write([]byte("`\n\n"))
		}
	}
	out.Write([]byte(")\n"))

	out.Write([]byte("var templateFiles = []*templates.TemplateFile{\n"))
	for key, val := range tmpls {
		out.Write([]byte("\t&templates.TemplateFile{\n\t\tName:     \"" + key + "\",\n\t\tTemplate: " + val + ",\n\t},\n"))
	}
	out.Write([]byte("}\n"))

}
