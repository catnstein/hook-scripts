# Hook Scripts

Collection of helpful Claude Code specific hooks written in Go

## Block Env Read

Prevents tooling from reading the .env file

### Usage

```JSON
{
  "hooks": {
    "PreToolUse": [
      {
        "matcher": "Read",
        "hooks": [
          {
            "type": "command",
            "command": "block-env-read"
          }
        ]
      }
    ]
  }
}
```


