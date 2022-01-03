package handler

import (
	"strings"
	"text/template"

	"github.com/gingama4/dotter/config"
)

func replaceVariable(v map[string]config.Variable, s string) (string, error) {
	tpl, err := template.New("").Parse(s)
	if err != nil {
		return s, err
	}

	w := new(strings.Builder)

	err = tpl.Execute(w, v)

	return w.String(), err
}

func ReplaceVariable(c *config.Config) {
	v := c.Variables

	for _, d := range c.Dotfiles {
		for _, s := range d.Steps {
			s.Src, _ = replaceVariable(v, s.Src)
			s.Target, _ = replaceVariable(v, s.Target)
			s.Cmd, _ = replaceVariable(v, s.Cmd)
		}
	}

}
