package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func usage() {
	println("Bad usage")
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		usage()
		os.Exit(-1)
	}
	queryString := flag.Args()[0]
	if queryString[0] != '.' {
		usage()
		os.Exit(-1)
	}
	buf := make([]byte, 0, 4*1024)
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.Read(buf[:cap(buf)])
	if err != nil {
		fmt.Printf("Error reading text: %v", err)
		os.Exit(-1)
	}
	text := string(buf[:s])
	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(text), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("%v", m[queryString[1:]])
}
