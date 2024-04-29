package main

import (
	"flag"
	"fmt"
	"log"
	"main/internal/app"

	"github.com/BurntSushi/toml"
)

// curl -X POST -H "Content-Type: application/json" -d '{"email":"u@e1.ru","password":"password"}' http://localhost:8080/sessions
// curl -X POST -H "Content-Type: application/json" -d '{"email":"user@mail.com","password":"password"}' http://localhost:8080/sessions
// curl -X POST -H "Content-Type: application/json" -d '{"email":"user@mail.com","password":"password"}' http://localhost:8080/users

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/appConfig.toml", "config file path")
}
func main() {
	fmt.Println("Start Server")

	flag.Parse()
	config := app.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	// s := app.New(config)
	if err := app.Start(config); err != nil { // сделать на Run
		log.Fatal(err)
	}
	// TODO: init config: cleanenv
	// TODO: init logger: slog
	// TODO: init storage: sqlite
	// TODO: init router: chi, "chi render"
	// TODO: run server

}
