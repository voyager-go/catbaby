package main

import (
	_ "threebody/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"threebody/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
