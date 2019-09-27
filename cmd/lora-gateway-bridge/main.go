package main

import (
	"os"
	"strings"

	"github.com/brocaar/lora-gateway-bridge/cmd/lora-gateway-bridge/cmd"
)

//go:generate ./doc.sh

var version string // set by the compiler

func main() {
	for _, pair := range os.Environ() {
		d := strings.SplitN(pair, "=", 2)
		os.Setenv(strings.ReplaceAll(strings.ReplaceAll(d[0], "_", "."), "..", "_"), d[1])
	}

	cmd.Execute(version)
}
