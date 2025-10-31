# PokedexCLI

A simple command-line interface (CLI) tool built in Go for exploring the Pokémon universe. Inspired by the classic Pokédex, this tool lets you explore locations, catch Pokémon, inspect your collection, and more – all using the free [PokeAPI](https://pokeapi.co/).

---

## Features

- `map` / `mapb` – paginated list of Pokémon locations
- `explore <location>` – discover wild Pokémon in a specific area
- `catch` – try to catch the Pokémon you just encountered
- `inspect <name>` – view full details of a caught Pokémon
- `pokedex` – list all Pokémon you have caught
- `help` – show command help
- `exit` – quit the program

---

## Quick Start (Download & Run)


   | Platform | File |
   |----------|------|
   | Windows  | `pokedexcli.exe` |
   | Linux    | `pokedexcli` |

2. **Make it executable** (macOS / Linux):

   ```bash
   chmod +x pokedexcli   # or pokedexcli-arm64
   ./pokedexcli          # macOS / Linux

### or double-click pokedexcli.exe on Windows