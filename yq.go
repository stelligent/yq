package main

import (
	"bufio"
	"io"
	"flag"
	"fmt"
	"log"
	"os"
	"gopkg.in/yaml.v2"
)

func usage() {
	println("Bad usage")
}

func parse_yaml(text string, queryString string) interface{} {
	m := make(map[interface{}]interface{})
	err := yaml.Unmarshal([]byte(text), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return (m[queryString[1:]])
}

func main() {
	flag.Parse()
	var io_type io.Reader
	if len(flag.Args()) < 1 || len(flag.Args()) > 2 {
		usage()
		os.Exit(-1)
	}
	queryString := flag.Args()[0]
	if queryString[0] != '.' {
		usage()
		os.Exit(-1)
	}

	buf := make([]byte, 0, 4*1024)
	if len(flag.Args()) == 2 {
		f, err := os.Open(flag.Args()[1])
		if err != nil {
			fmt.Printf("Error reading file: %v", err)
			os.Exit(-1)
		}
		io_type = f
	} else {
		io_type = os.Stdin
	}
	reader := bufio.NewReader(io_type)
	s, err := reader.Read(buf[:cap(buf)])
	if err != nil {
		fmt.Printf("Error reading text: %v", err)
		os.Exit(-1)
	}
	text := string(buf[:s])
	parsed_map := parse_yaml(text, queryString)
	fmt.Printf("%v", parsed_map)
}
