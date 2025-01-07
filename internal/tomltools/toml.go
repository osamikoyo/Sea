package tomltools

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
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
	Deps        []string `toml:"deps"`
}

func Get(file *os.File) (TEMP, error) {
	var template TEMP

	body, err := ioutil.ReadAll(file)
	if err != nil{
		return TEMP{}, nil
	}
	if isValid(string(body)) == 0{
		fmt.Println(".toml file is no valid")
	}

	if _, err := toml.DecodeReader(file, &template); err != nil {
		return TEMP{}, err
	}

	return template, nil
}
