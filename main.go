package main

import (
	"bufio"
	"bytes"
	"encoding/json"
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
	var b []byte
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
		var strs []string
		if len(args) == 0 {
			sc := bufio.NewScanner(os.Stdin)
			if sc.Scan() {
				b := sc.Bytes()
				strs = strings.Split(string(b), " ")
			}
		} else {
			strs = args
		}
		b, err = buildStructJSON(strs)
		if err != nil {
			return err
		}
		res = string(b)
	}

	fmt.Printf("%s\n", res)
	return nil
}

func buildStructJSON(args []string) ([]byte, error) {
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
		return nil, err
	}

	return b, nil
}

func printPretty(args []string) (string, error) {
	var strs []string
	if len(args) == 1 { // pretty optionのみならば、標準入力待ちになる
		sc := bufio.NewScanner(os.Stdin)
		if sc.Scan() {
			b := sc.Bytes()
			strs = strings.Split(string(b), " ")
		}
	} else {
		strs = args
	}
	b, err := buildStructJSON(strs)
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
	var strs []string
	if len(args) == 0 {
		sc := bufio.NewScanner(os.Stdin)
		if sc.Scan() {
			b := sc.Bytes()
			strs = strings.Split(string(b), " ")
		}
	} else {
		strs = args
	}
	b, err := json.Marshal(strs)
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
