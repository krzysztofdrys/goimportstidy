package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/krzysztofdrys/goimportstidy/format"
)

var local = flag.String("local", "", "local package name, used for grouping")
var write = flag.Bool("w", false, "write changes")

func usage() {
	fmt.Fprintf(os.Stderr, "usage: goimportstidy [flags] [path]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func errAndExit(f string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, f, args...)
	os.Exit(2)
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		usage()
	}
	file := flag.Arg(0)

	s, err := os.Stat(file)
	if err != nil {
		errAndExit("failed to stat file: %v", err)
	}
	f, err := ioutil.ReadFile(file)
	if err != nil {
		errAndExit("failed to read file: %v", err)
	}

	output := format.File(string(f), *local)

	if !*write {
		fmt.Print(string(output))
	}

	if err := ioutil.WriteFile(file, []byte(output), s.Mode()); err != nil {
		errAndExit("failed to format file: %v", err)
	}
}
