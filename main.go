package main

import (
	filefirewall "github.com/go/hook-scripts/file_firewall"
)

func main() {
	jsonMap, err := filefirewall.ReadPayload()
	if err != nil {
		filefirewall.Fail()
	}

	res, err := filefirewall.AnalyzeContainsEnv(jsonMap)

	if err != nil {
		filefirewall.Fail()
	}
	if res == true {
		filefirewall.Fail()
	}

	filefirewall.Succeed()
}
