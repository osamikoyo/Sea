package app

import (
	"fmt"
	"os"
	"sea/internal/directory"
	"sea/internal/templates/htmx"
	"sea/internal/templates/parser"
	"sea/internal/tomltools"
)

func Run(args []string) {
	switch args[1] {
	case "htmx":
		if err := htmx.Create(os.Args[2]); err != nil {
			fmt.Println(err)
		}
	case "search":
		body, err := directory.GetTempl(os.Args[2])
		if err != nil {
			fmt.Println(err)
		}

		templ, err := tomltools.Get(body)
		if err != nil {
			fmt.Println(err)
		}

		if err = parser.Pars(templ, os.Args[3], false); err != nil {
			fmt.Println(err)
		}
	case "create":
		if err := directory.Create(); err != nil {
			fmt.Println(err)
		}
	}
}
