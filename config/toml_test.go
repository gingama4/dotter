package config

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLoadConfig(t *testing.T) {
	testdata := []struct {
		data   string
		config Config
	}{
		{
			data: "testdata/config.toml",
			config: Config{
				DfPath:  "~/dotfiles",
				Log:     true,
				LogPath: "./dotter.log",
				Dotfiles: []Dotfile{
					{
						Name: "vim",
						Steps: []Step{
							{
								Name:   "link vimrc",
								Src:    "vim/vimrc",
								Target: "~/.vimrc",
								Type:   "ln",
								Force:  true,
							},
						},
					},
					{
						Name: "git",
						Steps: []Step{
							{
								Src:    ".gitconfig",
								Target: "~/.gitconfig",
								Type:   "cp",
							},
						},
					},
				},
				Variables: map[string]Variable{
					"test_string": Variable{
						Var: "testing",
					},
				},
			},
		},
	}

	for _, v := range testdata {
		t.Run("", func(t *testing.T) {
			acctual := LoadConfig(v.data)

			if diff := cmp.Diff(acctual, &v.config); diff != "" {
				t.Errorf("\nConfig value is mismatch (-acctual, +expected):\n%s", diff)
			}
		})
	}
}
