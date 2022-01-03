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

	for i, d := range c.Dotfiles {
		for j, s := range d.Steps {
			c.Dotfiles[i].Steps[j].Src, _ = replaceVariable(v, s.Src)
			c.Dotfiles[i].Steps[j].Target, _ = replaceVariable(v, s.Target)
			c.Dotfiles[i].Steps[j].Cmd, _ = replaceVariable(v, s.Cmd)
		}
	}

}
