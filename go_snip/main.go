package main

/*
Example usages in go file:

//go:generate go_snip testhelper pkgname
//go:generate go_snip typedlist pkgname ObjectName
//go:generate go_snip typedlistmap pkgname ObjectName TargetName
*/

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var root string = "src/github.com/tro3/go_snippets/"

var router map[string]func() = map[string]func(){
	"testhelper":   handleTestHelper,
	"typedlist":    handleTypedList,
	"typedlistmap": handleTypedListMap,
}

func main() {
	fn, ok := router[os.Args[1]]
	if !ok {
		panic("Unknown snippet: " + os.Args[1])
	}
	fn()
}

func generateFile(in, out string, replace map[string]string) {
	c := getAndReplace(in, replace)
	_ = ioutil.WriteFile(out, c, 0666)
	fmt.Printf("Generated %s\n", out)
}

func getAndReplace(fname string, replace map[string]string) []byte {
	c, _ := ioutil.ReadFile(filepath.Join(os.Getenv("GOPATH"), root, fname))
	for key, val := range replace {
		c = bytes.Replace(c, []byte(key), []byte(val), -1)
	}
	return c
}

func handleTestHelper() {
	proj := os.Args[2] // Target project name

	generateFile(
		"testhelper.go",
		"testhelper.go",
		map[string]string{
			"snippets": proj,
		},
	)
}

func handleTypedList() {
	proj := os.Args[2] // Target project name
	kind := os.Args[3] // Target object name

	generateFile(
		"typedlist.go",
		fmt.Sprintf("%s_list.go", strings.ToLower(kind)),
		map[string]string{
			"snippets": proj,
			"TYPEA":    kind,
		},
	)
}

func handleTypedListMap() {
	proj := os.Args[2]  // Target project name
	kind := os.Args[3]  // Target object name
	mkind := os.Args[4] // Target resulting object name

	generateFile(
		"typedlistmap.go",
		fmt.Sprintf("%s_list_map.go", strings.ToLower(kind)),
		map[string]string{
			"snippets": proj,
			"TYPEA":    kind,
			"TYPEB":    mkind,
		},
	)
}
