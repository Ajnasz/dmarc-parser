package main

// PolicyPublished is a struct to represent policy_published in xml
type PolicyPublished struct {
	Domain string `xml:"domain"`
	Adkim  string `xml:"adkim"`
	Aspf   string `xml:"aspf"`
	P      string `xml:"p"`
	Sp     string `xml:"sp"`
	Pct    int    `xml:"pct"`
}
