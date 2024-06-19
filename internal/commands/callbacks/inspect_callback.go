package callbacks

import "github.com/stpiech/pokedexcli/internal/pokedex"

func InspectCallback(args []string) {
  pokemonName := args[0]

  pokedex.Inspect(pokemonName)
}
