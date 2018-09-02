// +build ignore

package main

import (
	"github.com/shurcooL/vfsgen"
	"log"
	"net/http"
)

func main() {

	err := vfsgen.Generate(http.Dir("ui/dist"), vfsgen.Options{
		PackageName:  "resources",
		BuildTags:    "!dev",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
