# Snake Game

A classic Snake game built with Go and Ebitengine.

## Requirements

- Go 1.25+
- Ebitengine v2

## Installation

```bash
go mod download
```

## Running

```bash
make run
```

Or build and run separately:

```bash
make build
./build/snake
```

## Controls

- **Arrow Keys** - Move snake
- **Enter** - Start game
- **R** - Restart after game over

## Project Structure

```
.
├── cmd/snake/main.go      # Entry point
├── internal/game/         # Game logic
│   ├── config.go         # Game constants
│   ├── direction.go      # Direction type
│   ├── food.go           # Food logic
│   ├── game.go           # Main game loop
│   ├── point.go          # Point type
│   └── snake.go          # Snake logic
├── Makefile              # Build commands
└── README.md
```

## Features

- Screen wrapping
- Self-collision detection
- Score tracking (food eaten)
- Start screen
- Game over screen with restart
