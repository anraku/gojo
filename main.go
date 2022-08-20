package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args
	err := run(args[1:])
	if err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	var pretty bool
	var array bool
	for _, v := range args {
		if v == "-p" {
			pretty = true
		} else if v == "-a" {
			array = true
		}
	}

	var res string
	var err error
	if array {
		res, err = printArray(args)
		if err != nil {
			return err
		}
	} else if pretty {
		res, err = printPretty(args)
		if err != nil {
			return err
		}
	} else {
		return errors.New("no option specified")
	}

	fmt.Printf("%s\n", res)
	return nil
}

func printPretty(args []string) (string, error) {
	jsonMap := make(map[string]string)
	for _, v := range args {
		kv := strings.Split(v, "=")
		if len(kv) != 2 {
			continue
		}
		jsonMap[kv[0]] = kv[1]
	}

	b, err := json.Marshal(jsonMap)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = json.Indent(&buf, b, "", "  ")
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func printArray(args []string) (string, error) {
	var idx int
	for i, v := range args {
		if v == "-p" {
			idx = i
			break
		}
	}
	args = append(args[:idx], args[idx+1:]...)

	b, err := json.Marshal(args)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = json.Indent(&buf, b, "", "  ")
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
