package callbacks

import (
  "fmt"
  "github.com/stpiech/pokedexcli/internal/commands/callbacks/fetchers"
)

func MapCallback() {
  jsonLocations, err := fetchers.LocationsJson(true)

  if err != nil {
    fmt.Println(err)
    return
  }

  for _, v := range jsonLocations.Results {
    fmt.Println(v.Name)
  }
}

