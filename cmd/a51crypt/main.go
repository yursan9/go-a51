package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/yursan9/a51"
)

func main() {
	keyPtr := flag.String("key", "", "Which `key` used to enrypt or decrypt message [A|B]")
	flag.Parse()

	var filename string
	switch *keyPtr {
	case "a":
		fallthrough
	case "A":
		filename = "Akey"
	case "b":
		fallthrough
	case "B":
		filename = "Bkey"
	default:
		flag.PrintDefaults()
		os.Exit(2)
	}

	key, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := a51.Crypt(s.Text(), key)
		fmt.Println(t)
	}

	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

}
