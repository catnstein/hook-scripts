package filefirewall

import (
	"testing"
)

var READ_PAYLOAD = map[string]any{
	"session_id":      "REDACTED_UUID",
	"transcript_path": "/home/user/.claude/projects/REDACTED_PROJECT_PATH/REDACTED_UUID.jsonl",
	"cwd":             "/home/user/work/project",
	"permission_mode": "default",
	"hook_event_name": "PreToolUse",
	"tool_name":       "Read",
	"tool_input": map[string]any{
		"file_path": "/home/user/work/project/.env",
	},
	"tool_use_id": "REDACTED_TOOL_USE_ID",
}

func TestFileFirewall_FilePathContainsEnv(t *testing.T) {
	res, err := AnalyzeContainsEnv(READ_PAYLOAD)
	if !res || err != nil {
		t.Errorf(`AnalyzeContainsEnv should return true and no error for file containing ".env". Returned %t and %v`, res, err)
	}
}
