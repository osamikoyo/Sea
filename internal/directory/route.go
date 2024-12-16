package directory

import (
	"os"
	"os/exec"
)

func GetTempl(name string) (*os.File, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(homedir + "/.sea/" + name + ".tomltools")
	if err != nil {
		return nil, err
	}

	return file, nil
}
func Create() error {
	cmd := exec.Command("mkdir", "~/.sea")
	return cmd.Run()
}
