// Defines words from the command line.  Simple.
package main

import (
	"flag"
	"fmt"
	"log"

	"golang.org/x/net/dict"
)

var word string

func init() {
	flag.Parse()
	word = flag.Arg(0)
}

func main() {
	dict, err := dict.Dial("tcp", "dict.org:dict")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer dict.Close()

	//dictList, err := dict.Dicts()
	//if err != nil {
	//log.Fatalln(err.Error())
	//}

	//for _, d := range dictList {
	//fmt.Println("Using dict", d.Name)
	//definitions, err := dict.Define(d.Name, word)
	definitions, err := dict.Define("wn", word)
	if err != nil {
		log.Fatalln(err.Error())
	}
	for _, def := range definitions {
		fmt.Println(string(def.Text))
	}
	//}

}
