package app

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"sea/internal/directory"
	"sea/internal/loger"
	"sea/internal/templates/parser"
	"sea/internal/tomltools"
)

func Run(args []string) {
	loggers := loger.New()
	switch args[1] {
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
	}
}
