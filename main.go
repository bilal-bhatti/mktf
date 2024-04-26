package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		// read all files with extension *.tf.json and convert
		// them to hcl and save as *.tf
		files, err := filepath.Glob("*.tf.json")
		if err != nil {
			panic(err)
		}

		for _, in := range files {
			if err := processFile(in); err != nil {
				panic(err)
			}
		}
	} else if len(args) == 1 && args[0] == "-" {
		// read file from stdin, convert and print to stdout
		var data []byte

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			data = append(data, scanner.Bytes()...)
		}
		err := scanner.Err()
		if err != nil {
			panic(err)
		}

		hclbuffer, err := toHCL(data)
		if err != nil {
			panic(err)
		}

		_, err = fmt.Fprint(os.Stdout, hclbuffer.String())
		if err != nil {
			panic(err)
		}
	} else {
		// assume all args are files and process them
		for _, in := range args {
			// TODO: file extension matches *.tf.json
			if err := processFile(in); err != nil {
				panic(err)
			}
		}
	}
}

func processFile(in string) error {
	data, err := os.ReadFile(in)
	if err != nil {
		return err
	}

	hclbuffer, err := toHCL(data)
	if err != nil {
		return err
	}

	fmt.Println(in, "=>", strings.TrimRight(in, ".json"))
	err = os.WriteFile(strings.TrimRight(in, ".json"), hclbuffer.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func toHCL(newData []byte) (*bytes.Buffer, error) {
	var data interface{}
	err := json.Unmarshal(newData, &data)
	if err != nil {
		return nil, err
	}

	mapsObjects := map[string]struct{}{"config": {}, "tags": {}, "tags_all": {}}
	formatted, err := terraformutils.Print(data, mapsObjects, "hcl")
	if err != nil {
		return nil, err
	}

	s := string(formatted)
	s = strings.ReplaceAll(s, " tags {", " tags = {")
	// s = strings.ReplaceAll(s, " tags_all {", " tags_all = {")
	// s = strings.ReplaceAll(s, " config {", " config = {")
	outbytes := []byte(s)

	return bytes.NewBuffer(outbytes), nil
}
