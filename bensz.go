package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ardnew/bensz/walk"
)

func main() {

	if len(os.Args) > 0 {
		var w walk.Walker
		for _, p := range os.Args[1:] {
			w.Add(walk.Walk(p))
		}
		fmt.Printf("%s\n", strings.Join(w.Strings(), "\n"))
	}
}
