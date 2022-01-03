package handler

import (
	"github.com/gingama4/dotter/config"
	"github.com/gingama4/dotter/logger"
)

type Handler struct {
	Config   *config.Config
	OnlyExec string
}

func (h *Handler) Run() {
	ReplaceVariable(h.Config)
	for _, d := range h.Config.Dotfiles {
		if !h.isTarget(d.Name) {
			continue
		}
		logger.Log().StepReset().AddField("name", d.Name).Debug("Dotfile")
		handler(&d)
	}
}

func (h *Handler) isTarget(name string) bool {
	if h.OnlyExec == "" {
		return true
	}

	if name == h.OnlyExec {
		return true
	}

	return false
}

func handler(d *config.Dotfile) {
	for n, s := range d.Steps {
		logger.Log().SetStep(1)
		logger.Log().AddField("num", n).AddField("type", s.Type).Debug("Step")
		switch s.Type {
		case "ln":
			logger.Log().StepUp().AddField("src", s.Src).AddField("target", s.Target).Debug("ln")
			CreateLink(&s)
		case "mkdir":
			logger.Log().StepUp().AddField("target", s.Target).Debug("mkdir")
		case "cp":
			logger.Log().StepUp().AddField("src", s.Src).AddField("target", s.Target).Debug("cp")
		case "cmd":
			logger.Log().StepUp().AddField("cmd", s.Cmd).Debug("cmd")
		default:
			logger.Log().StepUp().Warn("Non defined type")
		}
	}
}
