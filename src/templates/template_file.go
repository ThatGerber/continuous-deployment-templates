package templates

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/hashicorp/hcl/hcl/printer"
)

/*
The TemplateFile struct represents a single file in a template collection.

A template file includes:

* A name for the final parsed template.
* A Template file, which is a string representing a path to a template file.
* A body, which is the content of the Template file.
* The Raw bytes of the body. Because Go can't reference static files without
  or something to parse (you can't include files in a package and then
  use those files from a distributed binary), we use a separate tool to
  munge the file into a dynamically generated Golang assets.go file:
  `templates/**\/assets.go`

The HCL attribute needs to merge the content of the template and the original
file and write it out as a formatted (terraform fmt) HCL config.
*/
type TemplateFile struct {
	Name     string
	Template string
	Body     []byte
}

/*
Amalgamate consumes a slice of file paths and parse the file out into an HCL
ast.

This takes a templating engine, feeds in a template file (i.e. TemplateFile)
and merges the result with the content of a similarly named file.
*/
func (t *TemplateFile) Amalgamate(files []string) error {
	var (
		err         error
		isHCL       bool
		result      []byte
		newBytes    []byte
		tfFileNames []string
	)

	tfFileNames = []string{
		"tf",
		"tfvars",
	}

	// Load all the regular files, append them to each other.
	for _, f := range files {
		isHCL = false

		fi, err := os.Stat(f)
		// (1) Does it exist and (2) have text to parse?
		if err != nil || fi.Size() < 1 {
			continue
		}

		newBytes, err = ioutil.ReadFile(f)
		if err != nil {
			return err
		}

		for i := range tfFileNames {
			if t.Name[len(t.Name)-len(tfFileNames[i]):] == tfFileNames[i] {
				isHCL = true
			}
		}

		if isHCL {
			result, err = parseHCL(result, newBytes)
			if err != nil {
				return err
			}
		} else {
			result, err = parseText(result, newBytes)
			if err != nil {
				return err
			}
		}
	}

	t.Body = result

	return err
}

/*
Write out the HCL config into properly formatted HCL file.

Equivalent of running `terraform fmt` on the created file.
*/
func (t *TemplateFile) Write(dest string) error {
	var fileAbsPath string
	var file *os.File
	var err error

	// Get user's current directory if empty string is passed.
	if dest == "" {
		dest, err = os.Getwd()
		if err != nil {
			return err
		}
		fileAbsPath = filepath.Join(dest, t.Name)
	}
	// Open and clear file.
	file, err = os.Create(fileAbsPath)
	defer file.Close()
	if err != nil {
		return err
	}

	// Format bytes (if possible
	_, err = file.Write(t.Body)
	if err != nil {
		return fmt.Errorf("unable to write all bytes for %s", file.Name())
	}
	file.Sync()

	// The HCL package doesn't add a newline to the end of the file, so we'll
	// append one ourselves.
	fi, err := file.Stat()
	if err != nil {
		return err
	}
	endOfFile := []byte("\n")
	newLine := []byte("\n")

	_, err = file.ReadAt(endOfFile, fi.Size()-int64(len(newLine)))

	if !bytes.Equal(endOfFile, newLine) {
		_, err = file.Write(newLine)
	}

	return err
}

/*
DestAbsPath returns the absolute path to the destination file for the HCL
config. In this case, that is the current directory of the user and the template
name.
*/
func (t *TemplateFile) DestAbsPath() string {
	cwd, _ := os.Getwd()
	origFile := filepath.Join(cwd, t.Name)

	return origFile
}

/*
ParseHCLFile
*/
func parseHCL(src []byte, contents []byte) ([]byte, error) {
	var (
		o   *ast.File
		n   *ast.File
		buf *bytes.Buffer
		by  []byte
		ok  bool
		err error
	)

	n, err = hcl.Parse(string(contents))
	if err != nil {
		return src, err
	}

	// Empty file.
	if _, ok := n.Node.(*ast.ObjectList); !ok {
		return by, err
	}
	// First non-empty response. Nothing to merge with.
	if src != nil {
		// Now merge.
		o, err = hcl.Parse(string(src))
		if err != nil {
			return src, err
		}

		if n, ok = mergeNode(o, n).(*ast.File); !ok {
			err = fmt.Errorf("Error merging files. %s", n)
			return src, err
		}
	}
	buf = bytes.NewBuffer(by)
	err = printer.Fprint(buf, n)
	by = []byte(buf.String())

	return by, err
}

/*
parseText

@TODO Write this so that it can merge the two slices. For now, we'll just take
the new one, if it's not empty.
*/
func parseText(src []byte, contents []byte) ([]byte, error) {

	if len(contents) != 0 {
		src = contents
	}

	return src, nil
}

/*
MergeNode weaves together two ASTs into a single tree.

Check if Nodes are equal; If not, pass recursively to function.

@TODO Really needs to be cleaned up, I suppose. But hey, Hashicorp did it....
https://github.com/hashicorp/hcl/blob/master/hcl/printer/nodes.go#L109
*/
func mergeNode(o ast.Node, n ast.Node) ast.Node {
	// This "add" variable is used in several places. So I'll create it here.
	var add bool

	// Break now and return the original because the two nodes are different
	// types or they're the same anyways. If they're not the same type, return the
	// original, as well, because we don't want to try and merge disparate types.
	if reflect.TypeOf(o) != reflect.TypeOf(n) || reflect.DeepEqual(o, n) {

		return o
	}

	if reflect.DeepEqual(o, reflect.Zero(reflect.TypeOf(o)).Interface()) {
		return n
	}

	switch o := o.(type) {
	case *ast.File:
		// Merge root file
		n := reflect.ValueOf(n).Interface().(*ast.File)

		mergeNode(o.Node, n.Node)

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
					mergeNode(n.Items[i], oMatches.Items[l])
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

			// Filter out array of matches
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

			// Add those matches to new list if it doesn't already exist.
			for _, m := range matchers {
				add = true

				if len(e.Items) < 1 {
					e.Add(m)
					continue
				}

				// Creates a map of [string]bool. All of the keys are added to the map.
				// Because duplicate keys would override each other, if we have more
				// keys in our map than we have keys in just one list, we know these two
				// items are not equal (key-wise).
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

					// They match, so we'll do a merge instead.
					if len(matchedKeys) == len(m.Keys) {
						add = false
						mergeNode(e.Items[ani], m)
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

		mergeNode(o.Val, n.Val)

		// Merge any comments.
		mergeNode(o.LeadComment, n.LeadComment)
		mergeNode(o.LineComment, n.LineComment)

	case *ast.ObjectKey:
	// We shouldn't be merging object keys, really.

	case *ast.ObjectType:
		// An HCL Object
		n := reflect.ValueOf(n).Interface().(*ast.ObjectType)

		mergeNode(o.List, n.List)

	case *ast.LiteralType:
		// An HCL string, float, boolean, or number
		n := reflect.ValueOf(n).Interface().(*ast.LiteralType)

		o.Token = n.Token
		mergeNode(o.LeadComment, n.LeadComment)
		mergeNode(o.LineComment, n.LineComment)

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

	return o
}
