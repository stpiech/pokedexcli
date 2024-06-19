package pokedex

import (
	"fmt"
	"sync"

	"github.com/stpiech/pokedexcli/internal/commands/callbacks/fetchers"
)

type Pokedex struct {
  mu sync.Mutex
  data map[string]fetchers.FetchedPokemon   
}

var pokedexData Pokedex = Pokedex { sync.Mutex{}, make(map[string]fetchers.FetchedPokemon) }

func Add(pokemonName string, pokemonData fetchers.FetchedPokemon) {
  pokedexData.mu.Lock()
  defer pokedexData.mu.Unlock()
  pokedexData.data[pokemonName] = pokemonData
}

func Show() {
  pokedexData.mu.Lock()
  defer pokedexData.mu.Unlock()
  if len(pokedexData.data) > 0 {
    fmt.Println("Your Pokedex:")
    for _, pokemon := range pokedexData.data {
      fmt.Println(fmt.Sprintf("- %s", pokemon.Name))
    }
  } else {
    fmt.Println("Your Pokedex is empty") 
  }
}

func Inspect(pokemonName string) {
  pokedexData.mu.Lock()
  defer pokedexData.mu.Unlock()
  pokemon, ok := pokedexData.data[pokemonName]
  if !ok {
    fmt.Println("You have not cought that pokemon")
    return
  }
  fmt.Println(fmt.Sprintf("Name: %s", pokemon.Name))
  fmt.Println(fmt.Sprintf("Height: %d", pokemon.Height))
  fmt.Println(fmt.Sprintf("Weight: %d", pokemon.Weight))
  fmt.Println("Stats:")
  for _, stat := range pokemon.Stats {
    fmt.Println(fmt.Sprintf("-%s: %d", stat.Stat.Name, stat.BaseStat))
  }
  fmt.Println("Types:")
  for _, pokemonType := range pokemon.Types {
    fmt.Println(fmt.Sprintf("-%s", pokemonType.Type.Name))
  }
}

