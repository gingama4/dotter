package main

import (
	"fmt"

	"github.com/gingama4/dotter/config"
)

func main() {
	c := config.LoadConfig("")
	fmt.Printf("%+v\n", c)
}
