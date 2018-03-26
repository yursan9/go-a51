package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/yursan9/a51"
)

type keySlice []uint64

func (k *keySlice) String() string {
	return fmt.Sprintf("%X", *k)
}

func (k *keySlice) Set(val string) error {
	sp := strings.Split(val, " ")
	for _, v := range sp {
		i, err := strconv.ParseUint(v, 16, 8)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(15)
		}
		*k = append(*k, i)
	}

	return nil
}

func writeFile(filename string, key []byte) error {
	err := ioutil.WriteFile(filename, key, 0644)
	return err
}

func main() {
	var key keySlice
	flag.Var(&key, "key", "Pass `key` for generating keystream")
	framePtr := flag.Uint("frame", 0x134, "Pass `frame` for generating keystream")
	flag.Parse()

	if key == nil {
		flag.PrintDefaults()
		os.Exit(2)
	}

	frame := uint64(*framePtr)
	reg := a51.NewRegister()

	a51.KeySetup(reg, key, frame)
	a, b := a51.GenKeystream(reg)

	err := writeFile("Akey", a)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	err = writeFile("Bkey", b)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
