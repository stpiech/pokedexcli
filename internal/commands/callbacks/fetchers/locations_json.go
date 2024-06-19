package fetchers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/stpiech/pokedexcli/internal/pokecache"
)

type FetchedLocations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var page int = -1
var cacheData = pokecache.NewCache(5 * time.Second)

func LocationsJson(forward bool) (FetchedLocations, error) {
  controllPage(forward)

  res, err := http.Get(locationsUrl())
  if err != nil {
    return FetchedLocations{}, err 
  }
  defer res.Body.Close()

  decoder := json.NewDecoder(res.Body)
  var data FetchedLocations 
  err = decoder.Decode(&data)
  if err != nil {
    return FetchedLocations{}, err
  }

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
