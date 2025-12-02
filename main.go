package main

import (
	filefirewall "github.com/go/hook-scripts/file_firewall"
)

func main() {
	res, err := filefirewall.AnalyzeContainsEnv()

	if err != nil {
		filefirewall.Fail()
	}
	if res == true {
		filefirewall.Fail()
	}

	filefirewall.Succeed()
}
