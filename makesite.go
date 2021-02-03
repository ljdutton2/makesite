package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"html/template"
)


type Content struct {
Content string
}

func readFile(name string) string {
	fileContents, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func renderTemplate(filename string, data string) {
	t := template.Must(template.New("template.tmpl").ParseFiles(filename))
	content := Content{Content: data}
	err := t.Execute(os.Stdout, content)
	if err != nil {
		panic(err)
	}
}

func writeFile(file string, data string) {
	bytesToWrite := []byte(data)
	err := ioutil.WriteFile(file, bytesToWrite, 0644)

	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Hello, world!")
}
