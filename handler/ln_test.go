package handler

import (
	"fmt"
	"os"
	"testing"

	"github.com/gingama4/dotter/config"
)

func TestCreateLink(t *testing.T) {
	tmp := createTemp(t)
	defer tmp.remove(t)

	source := tmp.Name

	err := os.Mkdir(fmt.Sprintf("%s/testdir", tmp.Dir), os.ModeAppend)
	if err != nil {
		t.Errorf("\nfailed :%v", err)
	}

	type args struct {
		s config.Step
	}

	tests := []struct {
		name    string
		args    args
		acctual string
	}{
		{
			name: "Simple",
			args: args{
				s: config.Step{
					Name:   "",
					Force:  false,
					Backup: false,
					Src:    source,
					Target: fmt.Sprintf("%s/test_data1", tmp.Dir),
					Type:   "ln",
					Cmd:    "",
				},
			},
			acctual: fmt.Sprintf("%s/test_data1", tmp.Dir),
		},
		{
			name: "Force1",
			args: args{
				s: config.Step{
					Name:   "",
					Force:  true,
					Backup: false,
					Src:    source,
					Target: fmt.Sprintf("%s/test_data1", tmp.Dir),
					Type:   "ln",
					Cmd:    "",
				},
			},
			acctual: fmt.Sprintf("%s/test_data1", tmp.Dir),
		},
		{
			name: "Force2",
			args: args{
				s: config.Step{
					Name:   "",
					Force:  true,
					Backup: false,
					Src:    source,
					Target: fmt.Sprintf("%s/test_data2", tmp.Dir),
					Type:   "ln",
					Cmd:    "",
				},
			},
			acctual: fmt.Sprintf("%s/test_data2", tmp.Dir),
		},
		{
			name: "Directory",
			args: args{
				s: config.Step{
					Name:   "",
					Force:  false,
					Backup: false,
					Src:    fmt.Sprintf("%s/testdir", tmp.Dir),
					Target: fmt.Sprintf("%s/test_data3", tmp.Dir),
					Type:   "ln",
					Cmd:    "",
				},
			},
			acctual: fmt.Sprintf("%s/test_data3", tmp.Dir),
		},
		{
			name: "Force-Directory1",
			args: args{
				s: config.Step{
					Name:   "",
					Force:  true,
					Backup: false,
					Src:    fmt.Sprintf("%s/testdir", tmp.Dir),
					Target: fmt.Sprintf("%s/test_data3", tmp.Dir),
					Type:   "ln",
					Cmd:    "",
				},
			},
			acctual: fmt.Sprintf("%s/test_data3", tmp.Dir),
		},
		{
			name: "Force-Directory2",
			args: args{
				s: config.Step{
					Name:   "",
					Force:  true,
					Backup: false,
					Src:    fmt.Sprintf("%s/testdir", tmp.Dir),
					Target: fmt.Sprintf("%s/test_data4", tmp.Dir),
					Type:   "ln",
					Cmd:    "",
				},
			},
			acctual: fmt.Sprintf("%s/test_data4", tmp.Dir),
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			err := CreateLink(&v.args.s)
			if err != nil {
				t.Errorf("\nError is occured: \n%v", err)
			}

			_, err = os.Lstat(v.acctual)
			if err != nil {
				t.Errorf("\nLink not found: \n%v", err)
			}
		})
	}
}

func TestCreateLink_srcNotExists(t *testing.T) {
	tmp := createTemp(t)
	defer tmp.remove(t)

	source := fmt.Sprintf("%s/not_exists", tmp.Dir)

	s := config.Step{
		Name:   "",
		Force:  false,
		Backup: false,
		Src:    source,
		Target: fmt.Sprintf("%s/test_data1", tmp.Dir),
		Type:   "ln",
		Cmd:    "",
	}

	err := CreateLink(&s)
	if err != nil {
		t.Errorf("\nError is occured: \n%v", err)
	}

	i, err := os.Lstat(s.Target)
	if err == nil {
		t.Errorf("\nLink found: \n%v", i)
	}
}

func TestCreateLink_AlreadyExists(t *testing.T) {
	tmp := createTemp(t)
	defer tmp.remove(t)

	source := tmp.Name

	s := config.Step{
		Name:   "",
		Force:  false,
		Backup: false,
		Src:    source,
		Target: fmt.Sprintf("%s/test_data1", tmp.Dir),
		Type:   "ln",
		Cmd:    "",
	}

	err := CreateLink(&s)
	if err != nil {
		t.Errorf("\nError is occured: \n%v", err)
	}

	err = CreateLink(&s)
	if err != nil {
		t.Errorf("\nError is occured: \n%v", err)
	}

	_, err = os.Lstat(s.Target)
	if err != nil {
		t.Errorf("\nLink not found: \n%v", err)
	}
}
