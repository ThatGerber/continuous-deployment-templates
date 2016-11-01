package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	templateName := os.Args[1]
	files := os.Args[2:]
	out, _ := os.Create("templates.go")
	out.Write([]byte("package " + templateName + "\n\nconst (\n"))
	for _, file := range files {
		tokens := strings.Split(file, ".")
		constName := tokens[0] + strings.Title(tokens[1])
		f, err := os.Open(templateName)
		if err != nil {
			fmt.Println("Error reading file: " + file)
			os.Exit(1)
		}
		out.Write([]byte(constName + " = `"))
		io.Copy(out, f)
		out.Write([]byte("`\n"))
	}
	out.Write([]byte(")\n"))

}
