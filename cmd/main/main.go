package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/stpiech/pokedexcli/internal/commands"
)

func main() {
  cliCommands := commands.Commands()

  for {
    fmt.Print("pokedex > ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()

    err := scanner.Err()
    if err != nil {
      panic(err)
    }

    commandKeywords := strings.Fields(scanner.Text()) 
    commandArgs := commandKeywords[1:]
    v, ok := cliCommands[commandKeywords[0]]
      

    if !ok {
      fmt.Println("Command not recognized. Try again")
      continue
    }

    if v.ArgCallback != nil && len(commandArgs) > 0 {
      v.ArgCallback(commandArgs)
      continue 
    }

    if v.Callback != nil && len(commandArgs) == 0 {
      v.Callback()
      continue
    }
    fmt.Println("Could not run a command. Make sure if args are set/unset")
  }
}

