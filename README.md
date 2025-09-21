# Bubble Shopping List

```text
┌─ BUBBLE SHOPPING LIST ─┐
│                        │
│ ○ Milk Tea             │
│ ○ Tapioca Pearls       │
│ ○ Brown Sugar          │
│ ○ Green Tea            │
│ ○ Coconut Jelly        │
│ ○ Taro Powder          │
│ ○ Ice                  │
│ ○ Straws               │
│                        │
│    ○ ○ ○ ○ ○ ○         │
└────────────────────────┘
```

[![Go Reference](https://pkg.go.dev/badge/github.com/theantichris/bubble-shopping-list.svg)](https://pkg.go.dev/github.com/theantichris/bubble-shoping-list) [![CI](https://github.com/theantichris/bubble-shopping-list/actions/workflows/go.yml/badge.svg)](https://github.com/theantichris/bubble-shopping-list/actions/workflows/go.yml)

A shopping list app written in Go and using [Bubble Tea](https://github.com/charmbracelet/bubbletea). Created following their tutorial

## Debugging

Bubble Tea apps assume control of stdin and stdout so you'll need to run [delve](https://github.com/go-delve/delve) in headless mode and connect to it.

```bash
# Start the debugger
$ dlv debug --headless --api-version=2 --listen=127.0.0.1:43000 .
API server listening at: 127.0.0.1:43000

# Connect to it from another terminal
$ dlv connect 127.0.0.1:43000
```
