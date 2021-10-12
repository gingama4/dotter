package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var fileName = [...]string{
	"config.toml",
	"config.tml",
}

func LoadConfig(configPath string) *Config {
	pwd, _ := os.Getwd()
	c := initConfig(pwd)
	data, err := getConfigFile(configPath, pwd)
	if err != nil {
		fmt.Println(err)
		return c
	}

	_, err = toml.Decode(data, c)
	if err != nil {
		fmt.Println(err)
	}

	return c
}

func getConfigFile(path, pwd string) (string, error) {
	var paths []string
	if path != "" {
		path = filepath.Clean(path)
		paths = append(paths, path)
	} else {
		for _, v := range fileName {
			tmp := filepath.Join(pwd, v)
			paths = append(paths, tmp)
		}
	}

	var err error
	for _, v := range paths {
		bytes, e := ioutil.ReadFile(v)
		if e == nil {
			return string(bytes), nil
		}
		err = e
	}

	return "", err
}
