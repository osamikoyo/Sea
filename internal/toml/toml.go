package toml

import (
	"github.com/BurntSushi/toml"
	"os"
)

type TEMPLATE struct {
	Directories []string `toml:"directories"`
	Files       []string `toml:"files"`
	Commands    []string `toml:"commands"`
}

func Get(file *os.File, name string) (TEMPLATE, error) {
	var template TEMPLATE

	if _, err := toml.DecodeReader(file, &template); err != nil {
		return TEMPLATE{}, err
	}

	return template, nil
}
