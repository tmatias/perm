package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tmatias/perm/perm"
)

func main() {
	f := flag.String("file", "", "file for which to describe permissions")
	flag.Parse()

	fi, err := os.Stat(*f)
	if err != nil {
		log.Fatal(err)
	}
	bits := fi.Mode().Perm()
	owner, group, other := perm.Permissions(int(bits))

	printPerm("owner", owner)
	printPerm("group members", group)
	printPerm("others", other)
}

func printPerm(who string, bits int) {
	what, err := perm.Description(bits)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s %s\n", who, what)
}
