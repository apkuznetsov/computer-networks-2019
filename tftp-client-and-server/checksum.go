package tftp

import (
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s file...\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func ChecksumCmd() {
	flag.Parse()
	for _, file := range flag.Args() {
		hash := checksum(file)
		fmt.Printf("%s %s\n", hash, file)
	}
}

func checksum(file string) string {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return err.Error()
	}

	hash := sha512.Sum512_256(buf)

	return fmt.Sprintf("%x", hash)
}
