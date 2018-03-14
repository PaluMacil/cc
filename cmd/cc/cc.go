package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/PaluMacil/cc/config"
)

func main() {
	conf, err := config.Load(config.Path())
	if err != nil {
		log.Fatalf("loading config: %s", err)
	}
	selected, compiler := conf.SelectedCompiler()
	if !selected {
		log.Fatalf("no compiler selected")
	}

	exec.Command(compiler.Path, os.Args[1:]...)
}
