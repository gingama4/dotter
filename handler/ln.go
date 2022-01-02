package handler

import (
	"os"
	"os/exec"

	"github.com/gingama4/dotter/config"
	"github.com/gingama4/dotter/logger"
)

// CreateLink is create a symbolic link.
func CreateLink(s *config.Step) error {

	if !isExists(s.Src) {
		logger.Log().Debug("src is not exist")
		return nil
	}

	newname := s.Target

	if checkLink(newname) {
		if !s.Force {
			return nil
		}
		overwriteLink(s)
		return nil
	}

	oldname := s.Src

	err := os.Symlink(oldname, newname)
	if err != nil {
		return err
	}

	return nil
}

func checkLink(n string) bool {

	ep := expandPath(n)

	info, err := os.Lstat(ep)

	if err != nil {
		logger.Log().AddField("info", info).AddField("err", err).Debug("os.Lstat:%s", ep)
		return false
	}

	logger.Log().AddField("info", info).Debug("os.Lstat:%s", ep)

	return true
}

func overwriteLink(s *config.Step) error {
	ep := expandPath(s.Target)
	sf, _ := os.Stat(ep)
	logger.Log().AddField("dir", sf.IsDir()).AddField("mode", sf.Mode()).Debug("Stat")

	var cs string
	if sf.IsDir() {
		cs = "-nfs"
	} else {
		cs = "-fs"
	}

	src := expandPath(s.Src)
	target := expandPath(s.Target)

	logger.Log().Debug("exec command: ln %s %s %s", cs, src, target)
	err := exec.Command("ln", cs, src, target).Run()
	if err != nil {
		return err
	}

	return nil
}
