# mycli - Command Memory Tool

A simple CLI tool written in Go to save and search Linux commands you forget.

## Installation
```bash
git clone https://github.com/MM02-sudo/mycli.git
cd mycli
go build -o mycli
sudo mv mycli /usr/local/bin/
```

## Usage
```bash
# Add a command
mycli add "ls -la"

# List all commands
mycli list

# Search for commands
mycli search "find"

# Delete a command
mycli delete 3
```

## Why I built this

I kept forgetting Linux commands, so I built this tool to help me remember them.
Built this before discovering fuzzy finder, but it still serves a different purpose - organizing commands with descriptions!

## Features

- Save commands with descriptions
- Search through your command history
- Delete commands you no longer need
- Works from any directory
```
