package callbacks

import (
	"fmt"
	"math/rand"

	"github.com/stpiech/pokedexcli/internal/commands/callbacks/fetchers"
	"github.com/stpiech/pokedexcli/internal/pokedex"
)

func CatchCallback(args []string) {
  pokemonName := args[0]

  fmt.Println(fmt.Sprintf("Throwing a Pokeball at %s...", pokemonName))

  pokemon, err := fetchers.PokemonJson(pokemonName)
  if err != nil {
    fmt.Println(err)
    return
  }

  if 70 > rand.Intn(pokemon.BaseExperience) {
    fmt.Println(fmt.Sprintf("%s was caught!", pokemonName))
    pokedex.Add(pokemonName, pokemon)
  } else {
    fmt.Println(fmt.Sprintf("%s has escaped", pokemonName))
  }
}

