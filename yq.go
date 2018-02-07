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

func tooManyArguments() {
	println("Too many arguments provided. Maximum number of arguments is 2.")
}

func tooFewArguments() {
	println("Too few arguments provided. Requires at least 1 argument.")
}

func invalidQueryString() {
	println("Invalid query string.")
}

func validateArgumentCount(count int) {
	if count < 1 {
		tooFewArguments()
		os.Exit(-1)
	}
	if count > 2 {
		tooManyArguments()
		os.Exit(-1)
	}
}

func validateQueryString(queryString string) {
	if queryString[0] != '.' {
		invalidQueryString()
		os.Exit(-1)
	}
}

func getIoType(args []string) (io.Reader, error) {
	if len(args) == 2 {
		f, err := os.Open(args[1])
		if err != nil {
			return nil, err
		}
		return f, nil
	}
	return os.Stdin, nil
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
	validateArgumentCount(len(flag.Args()))
	queryString := flag.Args()[0]
	validateQueryString(queryString)
	buf := make([]byte, 0, 4*1024)
	io_type, err := getIoType(flag.Args())
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		os.Exit(-1)
	}

	reader := bufio.NewReader(io_type)
	s, err := reader.Read(buf[:cap(buf)])
	if err != nil {
		fmt.Printf("Error reading text: %v", err)
		os.Exit(-1)
	}
	text := string(buf[:s])
	parsed_map := parse_yaml(text, queryString)
	fmt.Printf("%v\n", parsed_map)
}
