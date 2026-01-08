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

type ReadPayload struct {
	SessionID      string `json:"session_id"`
	TranscriptPath string `json:"transcript_path"`
	Cwd            string `json:"cwd"`
	PermissionMode string `json:"permission_mode"`
	HookEventName  string `json:"hook_event_name"`
	ToolName       string `json:"tool_name"`
	ToolInput      struct {
		FilePath string `json:"file_path"`
		Path     string `json:"path"`
	} `json:"tool_input"`
	ToolUseID string `json:"tool_use_id"`
}

func ParsePayload() (*ReadPayload, error) {
	input, err := io.ReadAll(os.Stdin)

	Debug(string(input))
	if err != nil {
		return nil, err
	}

	return DeserialzeReadPayload(input)
}

func DeserialzeReadPayload(input []byte) (*ReadPayload, error) {
	var jsonMap ReadPayload
	if err := json.Unmarshal(input, &jsonMap); err != nil {
		return nil, err
	}

	return &jsonMap, nil
}
