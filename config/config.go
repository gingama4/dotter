package config

// Config is user configuration
type Config struct {
	DfPath    string              `toml:"dotfiles_path"`
	Log       bool                `toml:"log"`
	LogPath   string              `toml:"log_path"`
	Dotfiles  []Dotfile           `toml:"dotfile"`
	Variables map[string]Variable `toml:"variable"`
}

type Dotfile struct {
	Name  string `toml:"name"`
	Steps []Step `toml:"step"`
}

type Step struct {
	Name   string `toml:"name"`
	Force  bool   `toml:"force"`
	Backup bool   `toml:"backup"`
	Src    string `toml:"src"`
	Target string `toml:"target"`
	Type   string `toml:"type"`
	Cmd    string `toml:"cmd"`
}

type Variable struct {
	Var string `toml:"var"`
}

func (v Variable) String() string {
	return v.Var
}

func initConfig(pwd string) *Config {
	c := Config{
		DfPath:    pwd,
		Log:       true,
		LogPath:   pwd,
		Dotfiles:  []Dotfile{},
		Variables: map[string]Variable{},
	}

	return &c
}
