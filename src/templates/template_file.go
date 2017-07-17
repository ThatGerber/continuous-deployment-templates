package templates

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	ttemplate "text/template"
	"time"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/hashicorp/hcl/hcl/printer"
)

const (
	/**
	 * Time to wait while generating HCL template and getting original file. 10
	 * seconds is probably too high, but setting it there as a sanity check.
	*/
	FILE_LOAD_TIMEOUT = 10
)

/**
 * The TemplateFile struct represents a single file in a template collection.
 *
 * A template file includes:
 *
 * * A name for the final parsed template.
 * * A Template file, which is a string representing a path to a template file.
 * * A body, which is the content of the Template file.
 * * The Raw bytes of the body. Because Go can't reference static files without
 *   or something to parse (you can't include files in a package and then
 *   use those files from a distributed binary), we use a separate tool to
 *   munge the file into a dynamically generated Golang assets.go file:
 *   `templates/**\/assets.go`
 * * The HCL, which is a Terraform ast representing contents of any file in
 *   the user's working directory that is named the same as the
 *   TemplateFile.Name **as well as** the contents from the generated template
 *   file.
 *
 * The HCL attribute needs to merge the content of the template and the original
 * file and write it out as a formatted (terraform fmt) HCL config.
 */
type TemplateFile struct {
	Name     string
	Template string
	Body     string
	RawBody  []byte
	Hcl      *ast.File
}

/**
 * TemplateFileMessenger
 *
 * Contains a Files attribute and an Errors attribute that hold channels. These
 * channels manage the flow of information in and out of gofuncs related to file
 * processing.
 */
type TemplateFileMessenger struct {
	Files  chan string
	Errors chan error
}

/**
 * Consume a template and parse into HCL.
 *
 * This takes a templating engine, feeds in a template file (i.e. TemplateFile)
 * and merges the result with the content of a similarly named file.
 */
func (t *TemplateFile) ConsumeTemplateFile(engine *ttemplate.Template, inputs map[string]string) error {
	var err error
	var files map[int]string
	var messenger TemplateFileMessenger

	messenger.Files = make(chan string, 2)
	messenger.Errors = make(chan error, 1)

	// Original File
	go func(t *TemplateFile, messenger TemplateFileMessenger) {
		messenger.Files <- t.destFileAbsPath()
	}(t, messenger)
	// Add Template File.
	go t.writeToTempfile(engine, inputs, messenger)

	files = make(map[int]string)

	for i := 0; i < 2; i++ {
		select {
		case msg := <-messenger.Errors:
			return fmt.Errorf("Error getting file %s: %s", t.Name, msg)
		case f := <-messenger.Files:
			files[i] = f
		case <-time.After(time.Second * FILE_LOAD_TIMEOUT):
			return fmt.Errorf("File loading timed out after %d seconds.", FILE_LOAD_TIMEOUT)
		}
	}

	t.Hcl, err = t.mergeHCL(files)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error loading hcl from %v", files))
		return fmt.Errorf("Error: %v", err)
	}

	return nil
}

/**
 * Write Template to temporary file.
 */
func (t *TemplateFile) writeToTempfile(engine *ttemplate.Template, inputs map[string]string, messenger TemplateFileMessenger) {
	t.Body = string(t.RawBody)

	tmpl, err := engine.Parse(t.Body)
	if err != nil {
		err = fmt.Errorf("Error parsing data from %s.", t.Name)

		messenger.Errors <- err
	}

	// Create temporary file from template.
	tmpFile, err := ioutil.TempFile("", "")
	if err != nil {
		err = fmt.Errorf("Error creating tmp output file %s.", t.Name)

		messenger.Errors <- err
	}

	err = tmpl.Execute(tmpFile, &Input{
		Variables: inputs,
	})
	if err != nil {
		err = fmt.Errorf("Error executing templating file %s.", t.Name)

		messenger.Errors <- err
	}

	// Pass temp file name to HCL
	messenger.Files <- tmpFile.Name()
}

/**
 * Merges the HCL from multiple files and merges them into a single ast config.
 */
func (t *TemplateFile) mergeHCL(files map[int]string) (*ast.File, error) {
	result := &ast.File{}
	// Parse it
	// Load all the regular files, append them to each other.
	for _, f := range files {
		fi, err := os.Stat(f)
		// (1) Does it exist and (2) have text to parse?
		if err != nil {
			continue
		}
		if fi.Size() < 1 {
			continue
		}

		b, err := ioutil.ReadFile(f)
		if err != nil {

			return nil, err
		}

		c, err := hcl.Parse(string(b))
		if err != nil {

			return nil, err
		}

		var ok bool
		// Empty file. Gotta go.
		if _, ok := c.Node.(*ast.ObjectList); !ok {
			continue
		}
		// First response. Can't possibly merge.
		if _, ok := result.Node.(*ast.ObjectList); !ok {
			result = c
			continue
		}
		// Now merge.
		if result, ok = t.mergeNode(result, c).(*ast.File); !ok {

			return nil, fmt.Errorf("Error merging files. %s", result)
		}
	}

	return result, nil
}

/**
 * Write out the HCL config into properly formatted HCL.
 *
 * Equivalent of running `terraform fmt` on the created file.
 */
func (t *TemplateFile) write(dest string) (err error) {
	var fileAbsPath string
	var file *os.File

	// Get user's current directory if empty string is passed.
	if dest == "" {
		dest, err = os.Getwd()
		if err != nil {

			return err
		}
	}

	fileAbsPath = filepath.Join(dest, t.Name)
	file, err = os.Create(fileAbsPath)

	if err = printer.Fprint(file, t.Hcl.Node); err != nil {

		return err
	}

	return nil
}

/**
 * Absolute path to the destination file for the HCL config.
 */
func (t *TemplateFile) destFileAbsPath() string {
	cwd, _ := os.Getwd()
	origFile := filepath.Join(cwd, t.Name)

	return origFile
}

/**
 * Merges two ASTs into a single tree.
 *
 * Check if Nodes are equal; If not, pass recursively to function.
 *
 * 	@TODO Make it work, I guess.
 */
func (t *TemplateFile) mergeNode(o ast.Node, n ast.Node) (result ast.Node) {
	var add bool

	result = o
	// The return value can either be the original node (o) or a new node
	// passed by reference (result)

	// Break now and return the original because the two nodes are different
	// types or they're the same anyways. If they're not the same type, return the
	// original, as well, because we don't want to
	if reflect.TypeOf(o) != reflect.TypeOf(n) || reflect.DeepEqual(o, n) {

		return
	}

	if reflect.DeepEqual(o, reflect.Zero(reflect.TypeOf(o)).Interface()) {
		result = n
		return
	}

	switch o := o.(type) {
	case *ast.File:
		// Merge root file
		n := reflect.ValueOf(n).Interface().(*ast.File)

		t.mergeNode(o.Node, n.Node)

	case *ast.ObjectList:
		// Compare all items in list against each other.
		n := reflect.ValueOf(n).Interface().(*ast.ObjectList)

		// Look at all new ast items and check if there are items with matching
		// keys. If there are, merge those two items.
		var nkeys []string
		for i := range n.Items {
			nkeys = make([]string, len(n.Items[i].Keys))
			for k := range n.Items[i].Keys {
				nkeys = append(nkeys, n.Items[i].Keys[k].Token.Value().(string))
			}
			oMatches := o.Filter(nkeys...)
			if len(oMatches.Items) >= 1 {
				for l := range oMatches.Items {
					t.mergeNode(n.Items[i], oMatches.Items[l])
				}
			}

			o.Add(n.Items[i])
		}

		// We've merged the new list in. Time to flatten the list.
		// Filter the list against itself. If the two object lists are the same
		// lengths, no changes were made and we can exit.
		var done bool
		var e *ast.ObjectList
		var okeys []string
		var matchers []*ast.ObjectItem

		done = false
		for !done {
			e = &ast.ObjectList{}

			for i := range o.Items {
				okeys = make([]string, len(o.Items[i].Keys))
				for d := range o.Items[i].Keys {
					okeys[d] = o.Items[i].Keys[d].Token.Value().(string)
				}

				if len(okeys) == 0 {
					continue
				}
				m := o.Filter(okeys...)
				for j := range m.Items {
					m.Items[j].Keys = o.Items[i].Keys
				}

				matchers = append(matchers, m.Items...)
			}

			for _, m := range matchers {
				add := true

				if len(e.Items) < 1 {
					e.Add(m)
					continue
				}

				for ani := range e.Items {
					matchedKeys := make(map[string]bool)
					if len(e.Items[ani].Keys) == len(m.Keys) {
						for ik := range m.Keys {
							matchedKeys[m.Keys[ik].Token.Value().(string)] = true
						}
						for ik := range e.Items[ani].Keys {
							matchedKeys[e.Items[ani].Keys[ik].Token.Value().(string)] = true
						}
					}

					if len(matchedKeys) == len(m.Keys) {
						add = false
					}
				}
				if add {
					e.Add(m)
				}
			}

			if len(o.Items) == len(e.Items) {
				done = true
			} else {
				o.Items = e.Items
			}
		}

	case *ast.ObjectItem:
		// Compare two list items and merge.
		n := reflect.ValueOf(n).Interface().(*ast.ObjectItem)

		t.mergeNode(o.Val, n.Val)

		// Merge any comments.
		t.mergeNode(o.LeadComment, n.LeadComment)
		t.mergeNode(o.LineComment, n.LineComment)

	case *ast.ObjectKey:
	// We shouldn't be merging object keys, really.

	case *ast.ObjectType:
		// An HCL Object
		n := reflect.ValueOf(n).Interface().(*ast.ObjectType)

		t.mergeNode(o.List, n.List)

	case *ast.LiteralType:
		// An HCL string, float, boolean, or number
		n := reflect.ValueOf(n).Interface().(*ast.LiteralType)

		o.Token = n.Token
		t.mergeNode(o.LeadComment, n.LeadComment)
		t.mergeNode(o.LineComment, n.LineComment)

	case *ast.ListType:
		// An HCL List
		n := reflect.ValueOf(n).Interface().(*ast.ListType)

		additions := []ast.Node{}

		for _, nlItem := range n.List {
			add = true
			for _, olItem := range o.List {
				if reflect.DeepEqual(olItem, nlItem) {
					add = false
				}
			}
			if add {
				additions = append(additions, nlItem)
			}
		}

		for _, addition := range additions {
			o.Add(addition)
		}

	case *ast.CommentGroup:
		// Group of Comments without a line break or another node.
		n := reflect.ValueOf(n).Interface().(*ast.CommentGroup)

		// @TODO Write the merge of comments so that higher up comments are inserted
		//       further up within the list. Or merged with earlier comments.
		additions := []*ast.Comment{}

		for _, newComment := range n.List {
			add = true

			for _, oldComment := range o.List {
				if newComment.Text == oldComment.Text {
					add = false
				}
			}

			if add {
				additions = append(additions, newComment)
			}
		}

		for _, addition := range additions {
			o.List = append(o.List, addition)
		}

	case *ast.Comment:
		// A comment
		n := reflect.ValueOf(n).Interface().(*ast.Comment)

		o.Text = o.Text + string('\n') + string(n.Text)

	default:
		// We shouldn't be here. But if we do get here, return the original because
		// hopefully it's correct.

	}

	return
}
