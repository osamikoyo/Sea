package tomltools

import (
	"github.com/BurntSushi/toml"
	"os"
)

type Data struct {
	File string `toml:"file"`
	Data string `toml:"data"`
}

type TEMP struct {
	Directories []string `toml:"directories"`
	Files       []string `toml:"files"`
	Commands    []string `toml:"commands"`
	Contents    []Data   `toml:"contents"`
}

func Get(file *os.File) (TEMP, error) {
	var template TEMP

	if _, err := toml.DecodeReader(file, &template); err != nil {
		return TEMP{}, err
	}

	return template, nil
}
