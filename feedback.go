package main

import "encoding/xml"

// Feedback is the root sturct of the xml
type Feedback struct {
	XMLName         xml.Name        `xml:"feedback"`
	ReportMetadata  ReportMetadata  `xml:"report_metadata"`
	PolicyPublished PolicyPublished `xml:"policy_published"`
	Records         []Record        `xml:"record"`
}
