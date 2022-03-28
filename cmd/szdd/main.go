package main

import (
	"fmt"
	"github.com/psilva261/szdd"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	i := flag.String("i", "", "input file")
	o := flag.String("o", "", "output file")
	flag.Parse()

	if *i == "" || *o == "" {
		fmt.Printf("usage: szdd -i inputfile -o outputfile\n")
		os.Exit(1)
	}

	bs, err := os.ReadFile(*i)
	if err != nil {
		log.Fatalf("read: %v", err)
	}
	data, err := szdd.Expand(bs)
	if err != nil {
		log.Fatalf("expand: %v", err)
	}
	err = ioutil.WriteFile(*o, data, 0644)
	if err != nil {
		log.Fatalf("write: %v", err)
	}
}

