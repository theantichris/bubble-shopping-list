package main

import tea "github.com/charmbracelet/bubbletea"

func main() {

}

// model stores the application's state.
type model struct {
	choices  []string         // items on the shopping list
	cursor   int              // which list item the cursor is pointing at (map index)
	selected map[int]struct{} // which list items are selected
}

// Init returns an initial Cmd to perform I/O.
func (model model) Init() tea.Cmd {
	return nil // no I/O right now, please
}

// Update handles incoming events (tea.Msg) and updates the model, it can also
// return a tea.Cmd.
func (model model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// is it a key press?
	case tea.KeyMsg:

		// what key was pressed?
		switch msg.String() {

		// quit the program
		case "ctrl+c", "q":
			return model, tea.Quit

		// move the cursor up
		case "up", "w":
			if model.cursor > 0 {
				model.cursor--
			}

		// move the cursor down
		case "down", "s":
			if model.cursor < len(model.choices)-1 {
				model.cursor++
			}

		// select the item
		case "enter", " ":
			_, ok := model.selected[model.cursor]
			if ok {
				delete(model.selected, model.cursor)
			} else {
				model.selected[model.cursor] = struct{}{}
			}
		}
	}

	// return the update model to the Bubble Tea runtime for processing.
	return model, nil
}

// initialModel returns the initial application state.
func initialModel() model {
	return model{
		// the shopping list
		choices:  []string{"buy carrots", "buy celery", "buy kohlrabi"},
		selected: make(map[int]struct{}),
	}
}
