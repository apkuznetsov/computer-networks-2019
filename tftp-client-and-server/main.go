package tftp

import (
	"flag"
	"io/ioutil"
	"log"
)

var (
	address = flag.String("a", "127.0.0.1:69", "listen address")
	payload = flag.String("p", "payload.png", "file to serve to clients")
)

func main() {
	flag.Parse()

	p, err := ioutil.ReadFile(*payload)
	if err != nil {
		log.Fatal(err)
	}
	
	s := Server{Payload: p}

	err = s.ListenAndServe(*address)
	log.Fatal(err)
}
