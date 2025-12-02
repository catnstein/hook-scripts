package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func exit() {
	fmt.Fprintf(os.Stderr, "You are not allowed to read this file.")
	os.Exit(2)
}

// TODO: enable debug mode
func main() {
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
		// TODO: proper parsing of the env file
		if strings.Contains(filePathStr, ".env") {
			exit()
		}
	}

	os.Exit(0)
}
