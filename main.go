package a51

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type bit uint64

func xorBytes(a, b []byte) []byte {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	dst := make([]byte, n)
	for i := 0; i < n; i++ {
		dst[i] = a[i] ^ b[i]
	}
	return dst
}

func Crypt(msg string, key []byte) string {
	reader := strings.NewReader(msg)
	buf := make([]byte, 15)
	text := ""

	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				text += string(xorBytes(buf[:n], key))
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		t := make([]byte, 15)
		t = xorBytes(buf, key)
		text += string(t[:n])
	}

	return text
}
