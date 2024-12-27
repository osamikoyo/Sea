package saver

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"gitlab.com/osamikoyo/sea/internal/tomltools"
	"os"
)

func Save(template tomltools.TEMP, name string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	file, err := os.Create(fmt.Sprintf("%s/%s.toml", home, name))
	if err != nil {
		return err
	}

	body, err := toml.Marshal(template)
	if err != nil {
		return err
	}

	if _, err = file.Write(body); err != nil {
		return err
	}

	return nil
}
