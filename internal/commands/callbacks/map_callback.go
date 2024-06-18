package callbacks

import (
  "fmt"
  "github.com/stpiech/pokedexcli/internal/commands/callbacks/helpers"
)

func MapCallback() {
  jsonLocations, err := helpers.LocationsJson(true)

  if err != nil {
    fmt.Println("Couldn't fetch locations. Try later")
    return
  }

  for _, v := range jsonLocations.Results {
    fmt.Println(v.Name)
  }
}

