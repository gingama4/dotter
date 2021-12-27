package main

import (
	"flag"

	"github.com/gingama4/dotter/config"
	"github.com/gingama4/dotter/handler"
	"github.com/gingama4/dotter/logger"
)

var onlyExec string
var configPath string

func init() {
	flag.StringVar(&onlyExec, "only", "", "Run specific dotfile")
	flag.StringVar(&configPath, "path", "", "path to config")
}

func main() {
	flag.Parse()

	logger.InitLog(true)
	logger.SetLevel(logger.DEBUG)

	c := config.LoadConfig(configPath)
	h := handler.Handler{
		Config:   c,
		OnlyExec: onlyExec,
	}
	h.Run()
}
