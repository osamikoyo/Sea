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

func GenerateToml(name string) error {
	file, err := os.Create(name + ".toml")
	if err != nil {
		return err
	}

	if _, err := file.Write([]byte(`
directories = ["dir1", "dir2", "dir3"]
files = ["file1.txt", "file2.txt"]
comands = ["command1", "command2"]

deps = ["gorm.io/gorm", "github.com/go-chi/chi/v5"]

[[contents]]
file = "file1.txt"
data = "hello"

[[contents]]
file = "file2.txt"
data = "fd"


		
`)); err != nil {
		return err
	}
	return nil
}
