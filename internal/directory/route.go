package directory

import (
	"os"
)

func GetTempl(name string) (*os.File, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(homedir + "/.sea/" + name + ".toml")
	if err != nil {
		return nil, err
	}

	return file, nil
}
func Create() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	err = os.Mkdir(home+"/.sea", 0755)
	if err != nil {
		return err
	}

	return nil
}
