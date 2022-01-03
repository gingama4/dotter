package handler

import (
	"io"
	"os"

	"github.com/gingama4/dotter/config"
	"github.com/gingama4/dotter/logger"
)

func Copy(s *config.Step) error {
	es := expandPath(s.Src)
	et := expandPath(s.Target)

	if !isExists(es) {
		logger.Log().Debug("src is not exist")
		return nil
	}

	if isExists(et) {
		if !s.Force {
			return nil
		}

		os.Remove(et)
	}

	src, err := os.Open(es)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(et)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}
