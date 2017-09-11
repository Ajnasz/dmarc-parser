package main

// Record is a struct to represent a record in xml
type Record struct {
	Row         []RecordRow       `xml:"row"`
	Identifiers RecordIdentifers  `xml:"identifiers"`
	AuthResults RecordAuthResults `xml:"auth_results"`
}
