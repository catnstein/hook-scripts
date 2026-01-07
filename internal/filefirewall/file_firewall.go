package filefirewall

import (
	"fmt"
	"strings"
)

func AnalyzeReadContainsEnv(jsonMap *ReadPayload) (bool, error) {
	filePathStr := fmt.Sprintf("%v", jsonMap.ToolInput.FilePath)

	if strings.Contains(filePathStr, ".env") {
		return true, nil
	}

	return false, nil
}

func AnalyzeGrepContainsEnv(jsonMap *ReadPayload) (bool, error) {
	filePathStr := fmt.Sprintf("%v", jsonMap.ToolInput.Path)

	if strings.Contains(filePathStr, ".env") {
		return true, nil
	}

	return false, nil
}

func AnalyzeContainsEnv(jsonMap *ReadPayload) (bool, error) {
	if jsonMap.ToolName == "Read" {
		return AnalyzeReadContainsEnv(jsonMap)
	}
	if jsonMap.ToolName == "Grep" {
		return AnalyzeGrepContainsEnv(jsonMap)
	}
	return false, nil
}
