package commands
import (
  "github.com/stpiech/pokedexcli/internal/commands/definitions"
  "github.com/stpiech/pokedexcli/internal/commands/callbacks"
)

type Command struct {
  Name string
  Description string
  Callback func() 
}

func Commands() map[string]Command {
  definitionsWithCallback := make(map[string]Command)

  for k, v := range definitions.CommandsDefinition() { 
    var value Command
    value.Callback = callbacksMap()[k]
    value.Name = v.Name
    value.Description = v.Description
    definitionsWithCallback[k] = value
  }

  return definitionsWithCallback
}

func callbacksMap() map[string]func() {
  return map[string]func() {
    "exit": callbacks.ExitCallback,
    "help": callbacks.HelpCallback, 
    "map": callbacks.MapCallback,
    "mapr": callbacks.MaprCallback,
  }

}
