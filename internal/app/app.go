package app

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"gitlab.com/osamikoyo/sea/internal/directory"
	"gitlab.com/osamikoyo/sea/internal/loger"
	"gitlab.com/osamikoyo/sea/internal/templates/parser"
	"gitlab.com/osamikoyo/sea/internal/tomltools"
	"os"
)

func Run(args []string) {
	loggers := loger.New()
	if len(os.Args) < 2 {
		loggers.Info().Msg("to little arguments, use sea info")
		return
	}

	switch args[1] {
	case "info":
		directory.InfoPrintln()
	case "search":
		body, err := directory.GetTempl(os.Args[2])
		if err != nil {
			loggers.Error().Err(err)
		}

		var git bool
		git = false
		for _, flag := range os.Args {
			if flag == "-g" {
				git = true
			}
		}

		template, err := tomltools.Get(body)
		if err != nil {
			fmt.Println(err)
		}

		if err = parser.Pars(template, os.Args[3], false, git); err != nil {
			loggers.Error().Err(err)
		}
		fmt.Println(color.CyanString(`
  ####   ######    ##
 #       #        #  #
  ####   #####   #    #
      #  #       ######
 #    #  #       #    #
  ####   ######  #    #

`))
		loggers.Info().Msg("Template " + os.Args[2] + " parsed successfully")
	case "create":
		if err := directory.Create(); err != nil {
			loggers.Error().Err(err)
		}
		loggers.Info().Msg("Created successfully")
	case "install":
		if len(os.Args) < 3 {
			err := errors.New("to low arguments, use sea info")
			loggers.Error().Err(err)
		}
		if err := directory.Install(os.Args[2]); err != nil {
			loggers.Error().Err(err)
		}
		loggers.Info().Msg("Template installed successfully")

	case "generate":
		if len(os.Args) < 3 {
			err := errors.New("to low arguments, use sea info")
			loggers.Error().Err(err)
		}

		if err := directory.GenerateToml(os.Args[2]); err != nil {
			loggers.Error().Err(err)
		}
		loggers.Info().Msg("Template Generated successfully!")
	default:
		loggers.Info().Msg("arguments are not right,please, use sea info")
	}
}
