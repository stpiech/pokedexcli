package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/stpiech/pokedexcli/internal/pokecache"
  "github.com/stpiech/pokedexcli/internal/commontypes"
)

var page int = -1
var cache *pokecache.Cache = pokecache.NewCache(10 * time.Second)

func LocationsJson(forward bool) (commontypes.FetchedLocations, error) {
  controllPage(forward)
 
  cacheObj, err := cache.Get(locationsUrl())
  if err == nil {
    return cacheObj, nil
  }

  res, err := http.Get(locationsUrl())
  if err != nil {
    return commontypes.FetchedLocations{}, err 
  }
  defer res.Body.Close()

  decoder := json.NewDecoder(res.Body)
  var data commontypes.FetchedLocations 
  err = decoder.Decode(&data)
  if err != nil {
    return commontypes.FetchedLocations{}, err
  }
  cache.Add(locationsUrl(), data)
  return data, nil
}

func locationsUrl() string {
  baseUrl := "https://pokeapi.co/api/v2/location-area"  
  if page > 0 {
    return fmt.Sprintf("%s?offset=%d&limit=20", baseUrl, page * 20)
  }
  return baseUrl 
}

func controllPage(forward bool) {
  if forward == true {
    page++
  } else {
    if page != 0 {
      page--
    }
  }
}
