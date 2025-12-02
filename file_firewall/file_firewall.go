package filefirewall

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	NOT_ALLOWED_ERROR_MESSAGE = "You are not allowed to read this file unfortunately."
)

func exit() {
	fmt.Fprintf(os.Stderr, NOT_ALLOWED_ERROR_MESSAGE)
	os.Exit(2)
}

func Analyze() {
	input, err := io.ReadAll(os.Stdin)

	if err != nil {
		exit()
	}

	var jsonMap map[string]any
	if err := json.Unmarshal(input, &jsonMap); err != nil {
		exit()
	}

	toolInput, ok := jsonMap["tool_input"].(map[string]any)
	if ok {
		filePath := toolInput["file_path"]
		filePathStr := fmt.Sprintf("%v", filePath)

		if strings.Contains(filePathStr, ".env") {
			exit()
		}
	}

	os.Exit(0)
}
