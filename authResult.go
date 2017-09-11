package main

import "fmt"

// AuthResults is a struct
type AuthResults struct {
	Domain      string `xml:"domain"`
	Result      string `xml:"result"`
	Selector    string `xml:"selector"`
	HumanResult string `xml:"human_result"`
}

// GetStatus is a method
func (r AuthResults) GetStatus() string {
	return fmt.Sprintf("%s result %s", r.Domain, r.Result)
}
