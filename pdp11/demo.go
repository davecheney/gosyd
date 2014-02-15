package main

import (
	"go/build"
	"path/filepath"

	"github.com/davecheney/pdp11"
)

func stdin(c chan uint8) {
	for _, v := range []byte("unix\n") {
		c <- v
	}
}

func main() {
	pdp := pdp11.New()
	pdp.LoadMemory(pdp11.BOOTRK05)
	pdp.SetPC(002000)
	pdp.Attach(0, filepath.Join(build.Default.GOPATH, "src/github.com/davecheney/pdp11/rk0"))
	go stdin(pdp.Input)
	pdp.Run()
}
