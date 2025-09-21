package main

func main() {
	// model stores the application's state
	type model struct {
		choices  []string         // items on the shopping list
		cursor   int              // which list item the cursor is pointing at
		selected map[int]struct{} // which list items are selected
	}
}
