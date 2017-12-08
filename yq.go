package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

func usage() {
	println("Bad usage")
}

func parse_yaml(text string, queryString string) {
	m := make(map[interface{}]interface{})
	err := yaml.Unmarshal([]byte(text), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("%v", m[queryString[1:]])
}

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 || len(flag.Args()) > 2 {
		usage()
		os.Exit(-1)
	}
	queryString := flag.Args()[0]
	if queryString[0] != '.' {
		usage()
		os.Exit(-1)
	}
	if len(flag.Args()) == 2 {
		dat, err := ioutil.ReadFile(flag.Args()[1])
		if err != nil {
			fmt.Printf("Error reading text: %v", err)
			os.Exit(-1)
		}
		text := string(dat)
		parse_yaml(text, queryString)
	} else {
		buf := make([]byte, 0, 4*1024)
		reader := bufio.NewReader(os.Stdin)
		s, err := reader.Read(buf[:cap(buf)])
		if err != nil {
			fmt.Printf("Error reading text: %v", err)
			os.Exit(-1)
		}
		text := string(buf[:s])
		parse_yaml(text, queryString)
	}
}
