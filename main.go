package main

import (
	"bufio"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
)

var xmlFile string

func init() {
	flag.StringVar(&xmlFile, "fileName", "", "XML file name. If omitted, stdin will be used")

	flag.Parse()
}

func fatal(err ...interface{}) {
	fmt.Fprintf(os.Stderr, fmt.Sprint(err...))
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}

func main() {
	v := Feedback{}

	var content []byte
	var err error

	if xmlFile == "" {
		reader := bufio.NewReader(os.Stdin)
		content, err = io.ReadAll(reader)

		if err != nil {
			fatal(err)
		}
	} else {
		content, err = os.ReadFile(xmlFile)

		if err != nil {
			fatal(err)
		}
	}

	err = xml.Unmarshal(content, &v)

	if err != nil {
		fatal(err)
	}

	for _, record := range v.Records {
		if record.AuthResults.Dkim.Result != "pass" {
			fatal(fmt.Sprintf("dkim %s", record.AuthResults.Dkim.GetStatus()))
		}

		if record.AuthResults.Spf.Result != "pass" {
			fatal(fmt.Sprintf("spf %s", record.AuthResults.Spf.GetStatus()))
		}
	}
}
