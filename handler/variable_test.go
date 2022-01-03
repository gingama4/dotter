package handler

import (
	"testing"

	"github.com/gingama4/dotter/config"
	"github.com/google/go-cmp/cmp"
)

func TestReplaceVariable(t *testing.T) {
	variables := map[string]config.Variable{
		"var1":  config.Variable{"test1"},
		"var2":  config.Variable{"test2"},
		"var3":  config.Variable{"test3"},
		"var_4": config.Variable{"test4"},
	}

	tests := []struct {
		name   string
		data   string
		expect string
	}{
		{
			name:   "Basic",
			data:   "replace {{.var1}}",
			expect: "replace test1",
		},
		{
			name:   "Double",
			data:   "replace {{.var1}} and {{.var2}}",
			expect: "replace test1 and test2",
		},
		{
			name:   "Trim",
			data:   "replace {{- .var1 -}} space",
			expect: "replacetest1space",
		},
		{
			name:   "Non template",
			data:   "not replace string",
			expect: "not replace string",
		},
		{
			name:   "Under score var name",
			data:   "Under score {{.var_4}}",
			expect: "Under score test4",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			s, err := replaceVariable(variables, v.data)

			if err != nil {
				t.Errorf("\nError is occured: \n%v", err)
			}

			if s != v.expect {
				t.Errorf("\nnot much string: \nacctual: %s\nexpected: %s\n", s, v.expect)
			}
		})
	}
}

func TestReplaceVar(t *testing.T) {
	tests := []struct {
		name   string
		data   config.Config
		expect config.Config
	}{
		{
			data: config.Config{
				Variables: map[string]config.Variable{
					"var1": {Var: "test"},
				},
				Dotfiles: []config.Dotfile{
					{
						Steps: []config.Step{
							{
								Name:   "Test1 {{.var1}}",
								Force:  true,
								Backup: true,
								Src:    "Test2 {{.var1}}",
								Target: "Test3 {{.var1}}",
								Type:   "Test4 {{.var1}}",
								Cmd:    "Test5 {{.var1}}",
							},
						},
					},
				},
			},
			expect: config.Config{
				Dotfiles: []config.Dotfile{
					{
						Steps: []config.Step{
							{
								Name:   "Test1 {{.var1}}",
								Force:  true,
								Backup: true,
								Src:    "Test2 test",
								Target: "Test3 test",
								Type:   "Test4 {{.var1}}",
								Cmd:    "Test5 test",
							},
						},
					},
				},
			},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			ReplaceVariable(&v.data)

			if diff := cmp.Diff(v.data.Dotfiles, v.expect.Dotfiles); diff != "" {
				t.Errorf("\nConfig value is mismatch (-acctual, +expected):\n%s", diff)
			}
		})
	}
}
