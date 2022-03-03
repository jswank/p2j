package main

// simple cli to convert a POM to json
//
// $ p2j pomfile > jsonfile
// $ cat pomfile | p2j > jsonfile

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"regexp"

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

	b, err := json.MarshalIndent(conventionalMarshaller{parsedPom}, "", "  ")
	if err != nil {
		log.Printf("unable to marshal parsed pom: %s", err)
		os.Exit(1)
	}

	os.Stdout.Write(b)

}

// Regexp definitions
var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
var wordBarrierRegex = regexp.MustCompile(`(\w)([A-Z])`)

type conventionalMarshaller struct {
	Value interface{}
}

func (c conventionalMarshaller) MarshalJSON() ([]byte, error) {
	marshalled, err := json.Marshal(c.Value)

	converted := keyMatchRegex.ReplaceAllFunc(marshalled, bytes.ToLower)

	return converted, err
}
