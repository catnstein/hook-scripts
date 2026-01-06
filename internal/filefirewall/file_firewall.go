package filefirewall

import (
	"fmt"
	"strings"
)

// TODO: implementation for Grep missing
func AnalyzeContainsEnv(jsonMap *ReadPayload) (bool, error) {
	filePathStr := fmt.Sprintf("%v", jsonMap.ToolInput.FilePath)

	if strings.Contains(filePathStr, ".env") {
		return true, nil
	}

	return false, nil
}
