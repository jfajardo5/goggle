package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var help = flag.Bool("help", false, "Type in your search arguments after the goggle command, for example:\n\ngoggle food near me")

func main() {
	flag.Parse()
	query := strings.Join(flag.Args(), "+")
	url := "https://www.google.com/search?q="
	if *help || query == "" {
		flag.Usage()
		os.Exit(0)
	}
	fmt.Println(url + query)
}
