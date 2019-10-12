package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// add adds a and b, and returns the result
func Add(a, b int) int {
	return a + b
}

type config struct {
	Host struct {
		Address string `mapstructure:"address"`
		Port    int    `mapstructure:"port"`
	} `mapstructure:"host"`
}

func main() {
	fmt.Println(Add(1, 1))

	viper.SetDefault("host.address", "default")
	viper.SetDefault("host.port", 1234)

	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	var c config
	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c.Host.Address, c.Host.Port)
}
