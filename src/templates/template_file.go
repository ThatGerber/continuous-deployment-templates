package templates

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
)

type TemplateFile struct {
	Name     string
	Template string
	Body     string
}

func (t *TemplateFile) ConsumeTemplateFile() {
	// 2 refers to how many jumps back to get to the original template.Add
	// Sorry, magic numbers suck.
	_, filename, _, _ := runtime.Caller(2)
	dir, err := filepath.Abs(filepath.Dir(filename))
	if err != nil {
		log.Fatal(err)
	}
	tmplFile := filepath.Join(dir, t.Template)

	b, err := ioutil.ReadFile(tmplFile)
	if err != nil {
		log.Fatal(err)
	}

	t.Body = string(b)
}
