# Fishpot Simulator

## Overview

Fishpot Simulator is a terminal-based simulation written in Go that mimics a fishpot with moving fish and bubbles. The fish move horizontally and vertically within the terminal window, and bubbles rise from the bottom to the top of the screen.

## Project Structure

The project is organized into the following packages:

- `main.go`: The entry point of the application.
- `fish/`: Contains the logic for fish movement.
- `fishpot/`: Manages the fishpot environment, including drawing and updating the screen.
- `bubble/`: Contains the logic for bubble movement.

## Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/yourusername/fishpotSimulator.git
   cd fishpotSimulator
   ```

2. **Build the project:**
   ```sh
   go build -o fishpotSimulator main.go
   ```

3. **Run the executable:**
   ```sh
   ./fishpotSimulator
   ```

## Usage

Simply run the executable in your terminal. The fishpot simulator will start, displaying fish moving around and bubbles rising.

## Contributing

Feel free to fork this repository and submit pull requests.

Check out the todo.md file for ideas on how to contribute.

For major changes, please open an issue first to discuss what you would like to change.

You can run the code and tests using the following commands:

```sh
go run .
```

```sh
go test ./...
```
