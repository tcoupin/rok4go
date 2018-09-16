// +build ignore

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {

	err := vfsgen.Generate(http.Dir("ui/dist"), vfsgen.Options{
		PackageName:  "resources",
		BuildTags:    "!dev,!generate,!test",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
