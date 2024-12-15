package gomain

import (
	"fmt"
	"os"
)

func Write(file *os.File, name string) error {
	_, err := file.Write([]byte(fmt.Sprintf(`
		package main

		import "%s/internal/app"
		
		func main(){
			apps := app.Init()
			apps.Run()
		}
		
`, name)))
	if err != nil {
		return err
	}
	return nil
}
