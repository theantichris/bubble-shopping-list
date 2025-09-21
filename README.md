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

A shopping list TUI app written in Go and using [Bubble Tea](https://github.com/charmbracelet/bubbletea). Created following their [tutorial](https://github.com/charmbracelet/bubbletea/tree/master/tutorials/basics).

## Features

The application implements a basic interactive shopping list with:

- Keyboard navigation (arrow keys or w/s)
- Item selection (space or enter)
- Clean terminal UI
- Model-View-Update architecture using Bubble Tea

## Debugging

Bubble Tea apps assume control of stdin and stdout so you'll need to run [delve](https://github.com/go-delve/delve) in headless mode and connect to it.

```bash
# Start the debugger
$ dlv debug --headless --api-version=2 --listen=127.0.0.1:43000 .
API server listening at: 127.0.0.1:43000

# Connect to it from another terminal
$ dlv connect 127.0.0.1:43000
```

## Logging

You can't log to stdout with Bubble Tea. You can log to a file before starting the Bubble Tea program. To see what’s being logged in real time, run `tail -f debug.log` while you run your program in another window.

```go
if len(os.Getenv("DEBUG")) > 0 {
  f, err := tea.LogToFile("debug.log", "debug")
  if err != nil {
    fmt.Println("fatal:", err)
    os.Exit(1)
  }

  defer f.Close()
}
```
