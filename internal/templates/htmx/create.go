package htmx

import (
	"os"
	"os/exec"
	"sea/internal/gomain"
)

func Create(name string) error {
	cmd := exec.Command("mkdir", "cmd", "cmd/"+name, "internal", "internal/app", "internal/config",
		"internal/logger", "internal/data", "internal/data/models", "internal/view", "internal/view/layout",
		"internal/web", "internal/web/src")
	err := cmd.Run()
	if err != nil {
		return err
	}
	file, err := os.Create("cmd/" + name + "/main.go")
	if err != nil {
		return err
	}
	err = gomain.Write(file, name)
	cmd = exec.Command("go", "mod", "init", name)
	err = cmd.Run()
	if err != nil {
		return err
	}

	apps, err := os.Create("internal/app/app.go")
	if err != nil {
		return err
	}
	_, err = apps.Write([]byte(`
	package app

	import(
			"net/http"
		  )

	type App struct{
		s *http.Server
	}
	func Init() App {
		return App{
			s : &http.Server{Addr: "localhost:8080"}
		}
	}

	func (a App) Run(){
		
	}
`))

	cmd = exec.Command("npm install --prefix ./internal/web htmx.org@2.0.4")
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
