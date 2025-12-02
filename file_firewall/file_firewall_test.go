package filefirewall

import (
	"testing"
)

const READ_PAYLOAD = `{
    "session_id": "REDACTED_UUID",
    "transcript_path": "/home/user/.claude/projects/REDACTED_PROJECT_PATH/REDACTED_UUID.jsonl",
    "cwd": "/home/user/work/project",
    "permission_mode": "default",
    "hook_event_name": "PreToolUse",
    "tool_name": "Read",
    "tool_input": {
      "file_path": "/home/user/work/project/.env"
    },
    "tool_use_id": "REDACTED_TOOL_USE_ID"
  }`

func TestFileFirewall_FilePathContainsEnv(t *testing.T) {
	// res, _ := AnalyzeContainsEnv(READ_PAYLOAD)

}
