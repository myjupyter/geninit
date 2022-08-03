package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/myjupyter/geninit/internal/entity"
	"github.com/myjupyter/geninit/internal/parser"
	"github.com/myjupyter/geninit/internal/render"
)

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "", "select file with structs")
	flag.Parse()

	rawFile, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	f := &entity.File{
		RawFile: rawFile,
	}
	if err := parser.ParseTo(filename, f); err != nil {
		panic(err)
	}

	b, err := render.Render(f)
	if err != nil {
		panic(err)
	}

	fmt.Println(b.String())
}
