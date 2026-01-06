package filefirewall

import (
	"os"
	"testing"
)

const path = "./assets/sample-read.json"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestFileFirewall_FilePathContainsEnv(t *testing.T) {
	data, err := os.ReadFile(path)
	check(err)

	payload, err := DeserialzeReadPayload(data)

	res, err := AnalyzeContainsEnv(payload)
	if !res || err != nil {
		t.Errorf(`AnalyzeContainsEnv should return true and no error for file containing ".env". Returned %t and %v`, res, err)
	}
}
