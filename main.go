package main

import tea "github.com/charmbracelet/bubbletea"

func main() {

}

// model stores the application's state.
type model struct {
	choices  []string         // items on the shopping list
	cursor   int              // which list item the cursor is pointing at
	selected map[int]struct{} // which list items are selected
}

// Init returns an initial Cmd to perform I/O.
func (model model) Init() tea.Cmd {
	return nil // no I/O right now, please
}

// initialModel returns the initial application state.
func initialModel() model {
	return model{
		// the shopping list
		choices:  []string{"buy carrots", "buy celery", "buy kohlrabi"},
		selected: make(map[int]struct{}),
	}
}
