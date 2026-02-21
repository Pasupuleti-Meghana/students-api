package config

import (
	"fmt"
	"log"
	"os"
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string `yaml:"address" env-required:"true"`
}

type Config struct {
	Env string `yaml:"env" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer HTTPServer `yaml:"http_server"`
}

func MustLoad() *Config {
	var configpath string 
	configpath = os.Getenv("CONFIG_PATH")
	fmt.Println("config path:", configpath)

	if configpath == "" {
		flags := flag.String("config", "", "path to config file")
		flag.Parse()

		configpath = *flags

		
		if configpath == "" {
			log.Fatal("config path is not set")
		}
	}

	if _, err := os.Stat(configpath); err != nil{
		log.Fatalf("config file does not exist: %s", configpath)
	}

	var cfg Config 

	err := cleanenv.ReadConfig(configpath, &cfg)
	if err != nil {
		log.Fatalf("failed to read config file: %s", err.Error())
	}

	return &cfg
}