package parser

import (
	"gitlab.com/osamikoyo/sea/internal/tomltools"
	"io/ioutil"
	"os/exec"
	"strings"
)

func Pars(templ tomltools.TEMP, name string, par bool, git bool) error {
	cmd := exec.Command("mkdir", templ.Directories...)
	err := cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("touch", templ.Files...)
	err = cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("go", "mod", "init", name)
	if err := cmd.Run(); err != nil {
		return err
	}
	if par {
		var ch chan error
		for _, command := range templ.Commands {
			go func() {
				cms := strings.Split(command, " ")

				newcms := cms[1:]

				cmd = exec.Command(cms[0], newcms...)
				err = cmd.Run()
				if err != nil {
					ch <- err
				}
			}()
		}

		err := <-ch
		if err != nil {
			return err
		}
		select {}
	} else {
		for _, command := range templ.Commands {
			cms := strings.Split(command, " ")

			newcms := cms[1:]

			cmd = exec.Command(cms[0], newcms...)
			err = cmd.Run()
			if err != nil {
				return err
			}
		}
	}

	for _, content := range templ.Contents {

		cnt := strings.Replace(content.Data, "$", name, -1)

		err = ioutil.WriteFile(content.File, []byte(cnt), 0644)
		if err != nil {
			return err
		}
	}
	for _, d := range templ.Deps {
		cmd := exec.Command("go", "get", d)
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	if git {
		cmd := exec.Command("git", "init")
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}
