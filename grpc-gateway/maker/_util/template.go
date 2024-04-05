package _util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

func Render(tmplFile string, outputFile string, data any) {

	funcMap := template.FuncMap{
		"Time":  Time,
		"Upper": CapitalizeFirstChar,
		"Kebab": KebabCase,
	}

	var buffer bytes.Buffer
	tmpl, err := template.New(filepath.Base(tmplFile)).Funcs(funcMap).ParseFiles(tmplFile)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		os.Exit(1)
	}
	err = tmpl.Execute(&buffer, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		os.Exit(1)
	}
	output := buffer.Bytes()
	err = ioutil.WriteFile(outputFile, output, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}

	fmt.Println("::Output written to", outputFile)
}
