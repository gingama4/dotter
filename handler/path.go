package handler

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/gingama4/dotter/logger"
)

func expandPath(p string) string {
	u, _ := user.Current()
	ep := strings.Replace(p, "~", u.HomeDir, 1)

	ep = filepath.Clean(ep)
	ep, _ = filepath.Abs(ep)

	logger.Log().AddField("old", p).AddField("new", ep).Debug("expandPath")
	return ep
}

func isExists(p string) bool {
	ep := expandPath(p)
	_, err := os.Stat(ep)
	return err == nil
}
