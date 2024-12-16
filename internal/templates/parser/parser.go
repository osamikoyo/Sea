package parser

import (
	"os/exec"
	"sea/internal/tomltools"
	"strings"
)

func Pars(templ tomltools.TEMP, name string, par bool) error {
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

	return nil
}
