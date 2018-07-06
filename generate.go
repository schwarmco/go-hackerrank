package main

import (
	"flag"
	"io"
	"os"
	"path/filepath"
)

var domain = flag.String("domain", "algorithms", "hackerrank category/domain")
var name = flag.String("name", "", "name of the challenge")

func main() {

	flag.Parse()

	path := filepath.Join(".", *domain, *name)
	os.MkdirAll(path, os.ModePerm)

	copyTemplate("main", path)
	copyTemplate("main_test", path)
}

func copyTemplate(name, dest string) {
	from, err := os.Open(filepath.Join(".", name+".tpl"))
	if err != nil {
		panic(err)
	}
	defer from.Close()

	to, err := os.Create(filepath.Join(dest, name+".go"))
	if err != nil {
		panic(err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		panic(err)
	}

	return
}
