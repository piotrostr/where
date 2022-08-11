package main

import (
	"flag"
	"fmt"
)

func main() {
	port := flag.String("port", "80", "port to run on")
	flag.Parse()
	fmt.Printf("hello world on %s\n", *port)
}
