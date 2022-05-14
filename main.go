package main

import (
	_ "ProxyGet/internal/packed"

	"ProxyGet/internal/cmd"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
