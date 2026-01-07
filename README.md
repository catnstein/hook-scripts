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
            "command": "block-env"
          }
        ]
      }
    ]
  }
}
```

### Test commands with Claude Code

```sh
# Grep
claude "use the grep tool to tell me the value of MY_SECRET from the .env file" -p

# Read
claude "read .env file" -p
```


