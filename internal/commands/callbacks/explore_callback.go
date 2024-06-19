package callbacks

import (
  "fmt"
  "github.com/stpiech/pokedexcli/internal/commands/callbacks/fetchers"
)

func ExploreCallback(args []string) {
  locationName := args[0]

  jsonLocation, err := fetchers.LocationJson(locationName)

  if err != nil {
    fmt.Println("Couldn't fetch location. Try later")
    return
  }
  
  fmt.Println("Found Pokemon:")

  for _, pokemonEncounter := range jsonLocation.PokemonEncounters {
    fmt.Println(fmt.Sprintf("- %s", pokemonEncounter.Pokemon.Name)) 
  }

}
