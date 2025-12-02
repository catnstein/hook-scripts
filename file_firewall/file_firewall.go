package filefirewall

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func Analyze() (bool, error) {
	input, err := io.ReadAll(os.Stdin)

	// LogDebug(string(input))
	if err != nil {
		return false, err
	}

	var jsonMap map[string]any
	if err := json.Unmarshal(input, &jsonMap); err != nil {
		return false, err
	}

	toolInput, ok := jsonMap["tool_input"].(map[string]any)
	if ok {
		filePath := toolInput["file_path"]
		filePathStr := fmt.Sprintf("%v", filePath)

		if strings.Contains(filePathStr, ".env") {
			return false, nil
		}
	}

	return true, nil
}
