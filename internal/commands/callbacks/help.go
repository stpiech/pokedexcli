package callbacks

import (
  "fmt"
  "github.com/stpiech/pokedexcli/internal/commands/definitions"
)

func HelpCallback() {
  fmt.Println("\nHelp for PokedexCLI") 
  fmt.Println("Available commands:") 

  for _, v := range definitions.CommandsDefinition() {
    fmt.Printf("\n> %s\n", v.Name)
    fmt.Println("  Description:", v.Description)
  }
  fmt.Println()
}

