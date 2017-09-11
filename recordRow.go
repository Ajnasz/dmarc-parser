package main

import "net"

// RecordRow is a struct to represent a row in xml
type RecordRow struct {
	SourceIP        net.IP          `xml:"source_ip"`
	Count           int             `xml:"count"`
	PolicyEvaluated PolicyEvaluated `xml:"policy_evaluated"`
}
