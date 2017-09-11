package main

import (
	"encoding/xml"
	"flag"
	"io/ioutil"
	"log"
)

var xmlFile string

func init() {
	flag.StringVar(&xmlFile, "fileName", "", "XML file name")

	flag.Parse()
}

func main() {
	v := Feedback{}
	content, err := ioutil.ReadFile(xmlFile)

	if err != nil {
		log.Fatal(err)
	}

	err = xml.Unmarshal(content, &v)

	if err != nil {
		log.Fatal(err)
	}

	for _, record := range v.Records {
		if record.AuthResults.Dkim.Result != "pass" {
			log.Fatal("dkim", " ", record.AuthResults.Dkim.GetStatus())
		}

		if record.AuthResults.Spf.Result != "pass" {
			log.Fatal("spf", " ", record.AuthResults.Spf.GetStatus())
		}
	}
}
