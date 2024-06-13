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
    "exit": {
      Name: "exit",
      Description: "Exits the program",
    },
  }
}
