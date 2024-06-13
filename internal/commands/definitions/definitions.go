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
    "exit": {
      Name: "exit",
      Description: "Exits the program",
    },
  }
}
