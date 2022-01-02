package handler

import (
	"os"

	"github.com/gingama4/dotter/config"
)

func CreateDir(s *config.Step) error {

	err := os.MkdirAll(s.Target, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
