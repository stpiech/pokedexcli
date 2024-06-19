package definitions

type CommandDefinition struct {
  Name string
  Description string
}

func CommandsDefinition() map[string]CommandDefinition {
  return map[string]CommandDefinition {
    "help": {
      Name: "help",
      Description: "Displays help",
    },
    "map": {
      Name: "map",
      Description: "Shows 20 location areas. Each next usage will load next 20 locations.",
    },
    "mapr": {
      Name: "mapr",
      Description: "Shows previous 20 locations. If no locations were initialized, initializes them.",
    },
    "explore": {
      Name: "explore",
      Description: "Explores location and show its Pokemons. Usage: explore <area_name>",
    },
    "catch": {
      Name: "catch",
      Description: "Tries to catch a pokemon. Usage: explore <pokemon_name>",
    },
    "inspect": {
      Name: "inspect",
      Description: "Inspects already caught pokemon",
    },
    "pokedex": {
      Name: "pokedex",
      Description: "Shows your Pokedex",
    },
    "exit": {
      Name: "exit",
      Description: "Exits the program",
    },
  }
}
