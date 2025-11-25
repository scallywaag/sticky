# sticky

A lightweight, colorful sticky notes manager for your terminal. Built with Go and SQLite for fast, local note organization.

## Features

- Create and manage notes - Add, delete, and mutate notes with ease
- Multiple lists - Organize notes into separate lists for different projects or contexts
- Pin important notes - Keep critical notes at the top of your list
- Cross off completed items - Mark tasks as done (moves them to bottom)
- Color coding - Highlight notes in red, green, blue, or yellow
- Local storage - All data stored locally in SQLite database
- Fast and simple - No external dependencies, just a single binary

## Installation

```bash
go install github.com/scallywaag/sticky@latest
```

Or clone and build from source:

```bash
git clone https://github.com/scallywaag/sticky.git
cd sticky
go build
```

## Usage

### Note Operations

```bash
# List all notes from a list
sticky -l <listname>

# Add a new note
sticky -a "Buy groceries"

# Delete a note by ID
sticky -d 5

# Mutate an existing note
sticky -m 3
```

### Note Attributes (Toggles)

Use these flags with `-a` (add) or `-m` (mutate):

```bash
# Pin a note (moves to top)
sticky -a "Important task" -p

# Cross off a note (moves to bottom)
sticky -m 2 -c

# Color a note
sticky -a "Critical bug" -r    # red
sticky -a "Completed" -g        # green
sticky -a "In progress" -b      # blue
sticky -a "Reminder" -y         # yellow
```

### List Management

```bash
# Show all existing lists
sticky -ls

# Add a new list
sticky -la "Work Tasks"

# Delete a list by ID
sticky -ld 2
```

### Help

```bash
# Display help menu
sticky -h
```

## Examples

```bash
# Create a shopping list and add items
sticky -la "Shopping"
sticky -l "Shopping" -a "Milk"
sticky -l "Shopping" -a "Eggs" -y
sticky -l "Shopping" -a "Bread" -p

# Mark an item as done
sticky -l "Shopping" -m 1 -c

# Create a work todo list with priorities
sticky -la "Work"
sticky -l "Work" -a "Fix production bug" -r -p
sticky -l "Work" -a "Review PRs" -b
sticky -l "Work" -a "Update docs" -g
```

## Author

[@scallywaag](https://github.com/scallywaag)
