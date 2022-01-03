package handler

import (
	"os/exec"
	"strings"

	"github.com/gingama4/dotter/config"
)

func ExecCmd(s *config.Step) error {

	if s.Cmd == "" {
		return nil
	}

	sp := strings.Split(s.Cmd, " ")

	err := exec.Command(sp[0], sp[1:]...).Run()
	if err != nil {
		return err
	}

	return nil
}
