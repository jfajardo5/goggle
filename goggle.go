package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
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
	err := openBrowser(url + query)
	if err != nil {
		log.Fatal(err)
	}
}

func openBrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("cmd", "/c", "start", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	return err
}
