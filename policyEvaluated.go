package main

// PolicyEvaluated is a struct to represent policy_evaluated in xml
type PolicyEvaluated struct {
	Disposition string `xml:"disposition"`
	Dkim        string `xml:"dkim"`
	Spf         string `xml:"spf"`
}
