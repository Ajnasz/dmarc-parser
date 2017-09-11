package main

// RecordAuthResults is a struct
type RecordAuthResults struct {
	Dkim AuthResults `xml:"dkim"`
	Spf  AuthResults `xml:"spf"`
}
