# CLI CHESS

CLI GUI for playing chess with an engine. Default engine to play against is stockfish. You need to have Stockfish engine installed.

## Install Stockfish

for example

```bash
brew install stockfish
```

or from official [github page](https://github.com/official-stockfish/Stockfish).

## Run

```bash
go run main.go
```

## Build

```bash
go build -o cli-chess main.go

# Add exec rights to the new build
chmod +x cli-chess
# run
./cli-chess
```

## Control

| Action                                  | Key    |
| --------------------------------------- | ------ |
| Move cursor                             | Arrows |
| Select                                  | Space  |
| Deselect/Exit app (if nothing selected) | Esc    |
| Restart the game                        | "r"    |
| Hint                                    | "h"    |
| Previous move                           | "b"    |
| Next move                               | "n"    |

