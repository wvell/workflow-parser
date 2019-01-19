package main

import (
	"fmt"
	"os"

	"github.com/github/actions-parser/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  " + os.Args[0] + " filename.workflow...")
		os.Exit(1)
	}

	for _, fn := range(os.Args[1:]) {
		parseFile(fn)
	}
}

func parseFile(fn string) {
	file, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	config, err := parser.Parse(file)
	if err != nil {
		panic(err)
	}

	for _, err := range config.Errors {
		fmt.Printf("%s: %s\n", fn, err)
	}

	if len(config.Errors) == 0 {
		fmt.Println(fn, "is a valid file with", plural(len(config.Actions), "action"), "and", plural(len(config.Workflows), "workflow"))
	}
}

func plural(n int, s string) string {
	if n == 1 {
		return fmt.Sprintf("%d %s", n, s);
	} else {
		return fmt.Sprintf("%d %ss", n, s);
	}
}