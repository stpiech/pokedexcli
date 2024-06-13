package callbacks

import (
  "os"
  "os/exec"
)

func ExitCallback() {
  cmd := exec.Command("clear")
  cmd.Stdout = os.Stdout
  cmd.Run()

  os.Exit(3)
}
