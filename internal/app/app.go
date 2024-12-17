package app

import (
	"fmt"
	"os"
	"sea/internal/directory"
	"sea/internal/loger"
	"sea/internal/templates/parser"
	"sea/internal/tomltools"
)

func Run(args []string) {
	logger := loger.New()
	switch args[1] {
	case "search":
		body, err := directory.GetTempl(os.Args[2])
		if err != nil {
			logger.Error().Err(err)
		}

		templ, err := tomltools.Get(body)
		if err != nil {
			fmt.Println(err)
		}

		if err = parser.Pars(templ, os.Args[3], false); err != nil {
			logger.Error().Err(err)
		}
		fmt.Println(`
  ####   ######    ##
 #       #        #  #
  ####   #####   #    #
      #  #       ######
 #    #  #       #    #
  ####   ######  #    #

`)
		logger.Info().Msg("Template " + os.Args[2] + " parsed successfully")
	case "create":
		if err := directory.Create(); err != nil {
			logger.Error().Err(err)
		}
	}
}
