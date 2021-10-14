// +build ignore

package main

import (
	"log"
	"net/http"

	"github.com/shurcool/vfsgen"
)

var fs http.FileSystem = http.Dir("./assets/")

func main() {
	err := vfsgen.Generate(fs, vfsgen.Options{
		PackageName:  "stage",
		VariableName: "Assets",
		BuildTags:    "!designer",
	})
	if err != nil {
		log.Fatalf("could not generate assets: %v", err)
	}
}
