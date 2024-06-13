package main

import (
  "fmt"
  "github.com/stpiech/pokedexcli/internal/commands"
)

func main() {
  cliCommands := commands.Commands()

  for {
    var userInput string
    fmt.Print("pokedex > ")
    fmt.Scan(&userInput)
    v, ok := cliCommands[userInput]

    if !ok {
      fmt.Println("Command not recognized. Try again")
      continue
    }

    v.Callback()
  }
}

