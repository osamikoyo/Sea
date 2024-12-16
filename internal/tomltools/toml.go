package tomltools

import (
	"github.com/BurntSushi/toml"
	"os"
)

type TEMPLATE struct {
	Directories []string `tomltools:"directories"`
	Files       []string `tomltools:"files"`
	Commands    []string `tomltools:"commands"`
}

func Get(file *os.File) (TEMPLATE, error) {
	var template TEMPLATE

	if _, err := toml.DecodeReader(file, &template); err != nil {
		return TEMPLATE{}, err
	}

	return template, nil
}
