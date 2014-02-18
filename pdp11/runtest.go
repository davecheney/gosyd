package main

import "os"
import "os/exec"

func main() {
	cmd := exec.Command("go", "test", "-v", "-run=PDP", "github.com/davecheney/pdp11")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
