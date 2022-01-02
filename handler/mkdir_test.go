package handler

import (
	"fmt"
	"os"
	"testing"

	"github.com/gingama4/dotter/config"
)

func TestCreateDir(t *testing.T) {
	tmp := createTemp(t)
	defer tmp.remove(t)

	s := config.Step{
		Name:   "",
		Force:  false,
		Backup: false,
		Src:    "",
		Target: fmt.Sprintf("%s/test-create-dir/subdir", tmp.Dir),
		Type:   "mkdir",
		Cmd:    "",
	}

	err := CreateDir(&s)
	if err != nil {
		t.Errorf("\nError is occured: \n%v", err)
	}

	sf, err := os.Stat(s.Target)
	if err != nil {
		t.Errorf("\n%s not found: \n%v", s.Target, err)
	}

	if sf != nil && !sf.IsDir() {
		t.Errorf("\n%s is not directory \n", s.Target)
	}
}
