package main

// simple cli to convert a POM to json
//
// $ p2j pomfile > jsonfile
// $ cat pomfile | p2j > jsonfile

import (
	"encoding/json"
	"log"
	"os"

	"github.com/2000Slash/gopom"
)

// p2j <filename> > jsonfile
// cat xml | x2j > jsonfile

var inputfile = "/dev/stdin"

func init() {
	if len(os.Args) > 1 {
		inputfile = os.Args[1]
	}
}

func main() {

	parsedPom, err := gopom.Parse(inputfile)
	if err != nil {
		log.Printf("unable to parse file (%s): %s", inputfile, err)
		os.Exit(1)
	}

	b, err := json.MarshalIndent(parsedPom, "", "  ")
	if err != nil {
		log.Printf("unable to marshal parsed pom: %s", err)
		os.Exit(1)
	}

	os.Stdout.Write(b)

}
