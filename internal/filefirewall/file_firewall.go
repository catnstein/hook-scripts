package filefirewall

import (
	"fmt"
	"strings"
)

// TODO: implementation for Grep missing
func AnalyzeContainsEnv(jsonMap map[string]any) (bool, error) {
	// TODO: better type
	toolInput, ok := jsonMap["tool_input"].(map[string]any)

	if ok {
		filePath := toolInput["file_path"]
		filePathStr := fmt.Sprintf("%v", filePath)

		if strings.Contains(filePathStr, ".env") {
			return true, nil
		}
	}

	return false, nil
}
