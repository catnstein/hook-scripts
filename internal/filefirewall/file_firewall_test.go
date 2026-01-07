package filefirewall

import (
	"os"
	"testing"
)

const sampleRead = "./assets/sample-read.json"
const sampleGrep = "./assets/sample-grep.json"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestFileFirewall_ReadFilePathContainsEnv(t *testing.T) {
	data, err := os.ReadFile(sampleRead)
	check(err)

	payload, err := DeserialzeReadPayload(data)

	res, err := AnalyzeContainsEnv(payload)
	if !res || err != nil {
		t.Errorf(`AnalyzeContainsEnv should return true and no error for file containing ".env". Returned %t and %v`, res, err)
	}
}

func TestFileFirewall_GrepFilePathContainsEnv(t *testing.T) {
	data, err := os.ReadFile(sampleGrep)
	check(err)

	payload, err := DeserialzeReadPayload(data)

	res, err := AnalyzeContainsEnv(payload)
	if !res || err != nil {
		t.Errorf(`AnalyzeContainsEnv should return true and no error for file containing ".env". Returned %t and %v`, res, err)
	}
}
