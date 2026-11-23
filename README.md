# Pokedex CLI

A small command-line Pokedex written in Go. It uses [PokeAPI](https://pokeapi.co/)
to browse Pokemon location areas from an interactive prompt.

## Requirements

- Go 1.27.2 or newer
- Internet access for PokeAPI requests

## Run

Clone the repository, then start the REPL:

```bash
go run .
```

You can also build a binary:

```bash
go build -o pokedexcli .
./pokedexcli
```

## Commands

At the `Pokedex >` prompt:

| Command | Description |
| --- | --- |
| `help` | Show available commands |
| `map` | Show the next page of location areas |
| `mapb` | Show the previous page of location areas |
| `explore <area-name>` | Show the Pokemon encountered in a location area |
| `exit` | Exit the Pokedex |

Example:

```text
Pokedex > map
canalave-city-area
...
Pokedex > mapb
...
Pokedex > explore canalave-city-area
Exploring canalave-city-area...
Found Pokemon:
- tentacruel
- starmie
...
Pokedex > exit
Closing the Pokedex... Goodbye!
```

The `map` and `mapb` commands request pages from PokeAPI and keep track of
the next and previous page URLs during the current session.

The `explore` command requests the Pokemon encounters for a location area.
Area names can be copied from the output of `map` or `mapb`. PokeAPI responses
are cached in memory for five minutes.

## Test

Run the test suite with:

```bash
go test ./...
```

## Project layout

```text
.
├── main.go                 # Application entry point
├── repl.go                 # Interactive command loop
├── command_*.go            # CLI command handlers
├── pokeapi/                # PokeAPI client and response types
├── pokecache/              # In-memory response cache
└── pokedexcli              # Prebuilt executable
```
