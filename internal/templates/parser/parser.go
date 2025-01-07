package parser

import (
	"fmt"
	"github.com/osamikoyo/sea/internal/loger"
	"github.com/osamikoyo/sea/internal/tomltools"
	"io/ioutil"
	"os/exec"
	"strings"
)


func getStrCommand(cmd string, arg ...string) string {
	result := cmd

	for _, a := range arg {
		result = fmt.Sprintf("%s %s", result, a)
	}

	return result
}

func Pars(templ tomltools.TEMP, name string, par bool, git bool, visual bool) error {



	logger := loger.New()

	logger.Info().Msg("Creating directories...")
	if visual {
		go func() {
			logger.Info().Msg(getStrCommand("mkdir", templ.Directories...))
		}()
	}
	cmd := exec.Command("mkdir", templ.Directories...)
	err := cmd.Run()
	if err != nil {
		return err
	}
	logger.Info().Msg("Success!")

	logger.Info().Msg("Creating files...")
	if visual {
		go func() {
			logger.Info().Msg(getStrCommand("touch", templ.Files...))
		}()
	}
	cmd = exec.Command("touch", templ.Files...)
	err = cmd.Run()
	if err != nil {
		return err
	}
	logger.Info().Msg("Success!")

	logger.Info().Msg("Creating go.mod")
	if visual {
		go func() {
			logger.Info().Msg(getStrCommand("go", "mod", "init", name))
		}()
	}
	cmd = exec.Command("go", "mod", "init", name)
	if err := cmd.Run(); err != nil {
		return err
	}
	logger.Info().Msg("Success!")
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
		logger.Info().Msg("execute commands")
		for _, command := range templ.Commands {
			cms := strings.Split(command, " ")

			newcms := cms[1:]

			cmd = exec.Command(cms[0], newcms...)
			err = cmd.Run()
			if err != nil {
				return err
			}
		}
		logger.Info().Msg("Success!")
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
