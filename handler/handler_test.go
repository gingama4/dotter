package handler

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/gingama4/dotter/logger"
)

func TestMain(m *testing.M) {
	logger.InitLog(false)

	m.Run()
}

type temp struct {
	Dir  string
	Name string
}

func createTemp(t *testing.T) *temp {
	t.Helper()
	tmp := t.TempDir()

	tmpFile, err := ioutil.TempFile(tmp, "testfile-")
	if err != nil {
		t.Fatalf("\nfatal: %v", err)
	}

	temp := temp{
		Dir:  tmp,
		Name: tmpFile.Name(),
	}

	return &temp
}

func (tp *temp) remove(t *testing.T) {
	t.Helper()
	err := os.RemoveAll(tp.Dir)
	if err != nil {
		t.Fatalf("\nfatal: %v", err)
	}
}
