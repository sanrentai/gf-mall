package main

import (
	_ "github.com/sanrentai/gf-mall/boot"
	_ "github.com/sanrentai/gf-mall/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
