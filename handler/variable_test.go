package handler

import (
	"testing"

	"github.com/gingama4/dotter/config"
)

func TestReplaceVariable(t *testing.T) {
	variables := map[string]config.Variable{
		"var1": config.Variable{"test1"},
		"var2": config.Variable{"test2"},
		"var3": config.Variable{"test3"},
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
