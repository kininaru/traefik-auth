package casdoor

import (
	"fmt"
	"os"
)

type Rule struct {
	Action      string
	Rule        string
	Application string
	Whitelist   []string
	Domains     []string
}

type Config struct {
	Endpoint     string `long:"endpoint" env:"ENDPOINT" default:"http://localhost:8000"`
	Organization string `long:"organization" env:"ORGANIZATION"`
	Application  string `long:"application" env:"APPLICATION"`
	ClientId     string `long:"client-id" env:"CLIENT_ID" default:""`
	ClientSecret string `long:"client-secret" env:"CLIENT_SECRET" default:""`
	Port         int    `long:"port" env:"PORT" default:"8001"`

	Rules map[string]*Rule `long:"rule.<name>.<param>" description:"Rule definitions, param can be: \"action\", \"rule\" or \"provider\""`
}

func NewConfig() Config {
	fmt.Println(os.Args)
	return Config{}
}

func (c Config) Validate() {

}
