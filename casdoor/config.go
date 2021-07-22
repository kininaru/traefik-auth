package casdoor

type Config struct {
	Endpoint     string `long:"endpoint" env:"ENDPOINT" default:"http://localhost:8000"`
	Organization string `long:"organization" env:"ORGANIZATION"`
	Application  string `long:"application" env:"APPLICATION"`
	ClientId     string `long:"client-id" env:"CLIENT_ID" default:""`
	ClientSecret string `long:"client-secret" env:"CLIENT_SECRET" default:""`
}
