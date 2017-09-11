package main

// DateRange is a struct to define date range field in xml
type DateRange struct {
	Begin int `xml:"begin"`
	End   int `xml:"end"`
}
