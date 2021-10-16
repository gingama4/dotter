package main

import (
	"flag"
	"fmt"

	"github.com/gingama4/dotter/config"
)

var onlyExec string
var configPath string

func init() {
	flag.StringVar(&onlyExec, "only", "", "Run specific dotfile")
	flag.StringVar(&configPath, "path", "", "path to config")
}

func main() {
	flag.Parse()

	c := config.LoadConfig(configPath)
	fmt.Printf("%+v\n", c)
}
