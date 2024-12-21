package app

import (
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

	case "search":
		body, err := directory.GetTempl(os.Args[2])
		if err != nil {
			loggers.Error().Err(err)
		}

		templ, err := tomltools.Get(body)
		if err != nil {
			fmt.Println(err)
		}

		if err = parser.Pars(templ, os.Args[3], false); err != nil {
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
	case "install":
		if err := directory.Install(os.Args[2]); err != nil {
			loggers.Error().Err(err)
		}
		loggers.Info().Msg("Template installed successfully")

	case "generate":
		if err := directory.GenerateToml(os.Args[2]); err != nil {
			loggers.Error().Err(err)
		}
		loggers.Info().Msg("Templ Generated successfully!")
	default:
		loggers.Info().Msg("arguments are not right,please, use sea info")
	}
}
