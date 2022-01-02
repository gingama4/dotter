package handler

import (
	"fmt"
	"os"
	"testing"

	"github.com/gingama4/dotter/config"
)

func TestExecCmd(t *testing.T) {
	tmp := createTemp(t)
	defer tmp.remove(t)

	tests := []struct {
		name    string
		arg     config.Step
		acctual func() error
	}{
		{
			name: "mkdir",
			arg: config.Step{
				Name:   "",
				Force:  false,
				Backup: false,
				Src:    "",
				Target: "",
				Type:   "cmd",
				Cmd:    fmt.Sprintf("mkdir %s/mkdir", tmp.Dir),
			},
			acctual: func() error {
				_, err := os.Stat(fmt.Sprintf("%s/mkdir", tmp.Dir))

				return err
			},
		},
		{
			name: "touch",
			arg: config.Step{
				Name:   "",
				Force:  false,
				Backup: false,
				Src:    "",
				Target: "",
				Type:   "cmd",
				Cmd:    fmt.Sprintf("touch %s/touch", tmp.Dir),
			},
			acctual: func() error {
				_, err := os.Stat(fmt.Sprintf("%s/touch", tmp.Dir))

				return err
			},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			err := ExecCmd(&v.arg)
			if err != nil {
				t.Errorf("\nError is occured: \n%v", err)
			}

			if err = v.acctual(); err != nil {
				t.Errorf("\nSomething went wrong: \n%v", err)
			}
		})
	}
}
