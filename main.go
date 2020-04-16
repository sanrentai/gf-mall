package main

import (
	_ "gf-mall/boot"
	_ "gf-mall/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
