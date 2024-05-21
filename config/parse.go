package config

import (
	"github.com/Netflix/go-env"
	_ "github.com/joho/godotenv/autoload"
)

var Parsed Config

func init() {
	if _, err := env.UnmarshalFromEnviron(&Parsed); err != nil {
		panic(err)
	}
}
