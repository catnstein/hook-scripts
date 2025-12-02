package filefirewall

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const (
	NOT_ALLOWED_ERROR_MESSAGE = "You are not allowed to read this file unfortunately."
)

func Fail() {
	fmt.Fprintf(os.Stderr, NOT_ALLOWED_ERROR_MESSAGE)
	os.Exit(2)
}

func Succeed() {
	os.Exit(0)
}

func ReadPayload() (map[string]any, error) {
	input, err := io.ReadAll(os.Stdin)

	// LogDebug(string(input))
	if err != nil {
		return nil, err
	}

	var jsonMap map[string]any
	if err := json.Unmarshal(input, &jsonMap); err != nil {
		return nil, err
	}

	return jsonMap, nil
}
