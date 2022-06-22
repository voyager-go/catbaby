package main

import (
	_ "catbaby/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"catbaby/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
