package main

import (
	"flag"
	"fmt"
	"log"
	"main/internal/app"

	"github.com/BurntSushi/toml"
)

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

	s := app.New(config)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
	// TODO: init config: cleanenv
	// TODO: init logger: slog
	// TODO: init storage: sqlite
	// TODO: init router: chi, "chi render"
	// TODO: run server

}
