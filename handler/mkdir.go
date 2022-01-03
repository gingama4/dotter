package handler

import (
	"os"

	"github.com/gingama4/dotter/config"
)

func CreateDir(s *config.Step) error {

	et := expandPath(s.Target)

	err := os.MkdirAll(et, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
