package config

import "fmt"

type Config struct {
}

var Conf Config

func init() {
	// config
	fmt.Println("config init!!!")
	//if _, err := toml.DecodeFile("test.toml", &Conf); err != nil {
	//	panic(err)
	//}
}
