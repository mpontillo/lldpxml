package main

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"

	"github.com/mpontillo/lldpxml"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		fmt.Println(arg)
		contents, err := os.ReadFile(arg)
		if err != nil {
			panic(err)
		}
		results := lldpxml.Parse(contents)
		spew.Dump(results)
	}
}
