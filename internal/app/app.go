package app

import (
	"fmt"
	"os"
	"sea/internal/templates/htmx"
)

func Run(args []string) {
	switch args[1] {
	case "htmx":
		if err := htmx.Create(os.Args[2]); err != nil {
			fmt.Println(err)
		}
	}
}
