package handler

import (
	"fmt"
	"os"
	"testing"

	"github.com/gingama4/dotter/config"
)

func TestCopy(t *testing.T) {
	tmp := createTemp(t)
	defer tmp.remove(t)

	tests := []struct {
		name string
		arg  config.Step
	}{
		{
			name: "Simple",
			arg: config.Step{
				Name:   "",
				Force:  false,
				Backup: false,
				Src:    tmp.Name,
				Target: fmt.Sprintf("%s/test-copy", tmp.Dir),
				Type:   "cp",
				Cmd:    "",
			},
		},
		{
			name: "Force1",
			arg: config.Step{
				Name:   "",
				Force:  true,
				Backup: false,
				Src:    tmp.Name,
				Target: fmt.Sprintf("%s/test-copy", tmp.Dir),
				Type:   "cp",
				Cmd:    "",
			},
		},
		{
			name: "Force2",
			arg: config.Step{
				Name:   "",
				Force:  true,
				Backup: false,
				Src:    tmp.Name,
				Target: fmt.Sprintf("%s/test-copy", tmp.Dir),
				Type:   "cp",
				Cmd:    "",
			},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			err := Copy(&v.arg)
			if err != nil {
				t.Errorf("\nError is occured: \n%v", err)
			}

			_, err = os.Stat(v.arg.Target)
			if err != nil {
				t.Errorf("\n%s not found: \n:%v", v.arg.Target, err)
			}
		})
	}

}
