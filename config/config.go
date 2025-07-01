package config

import "github.com/caarlos0/env/v11"

type Config struct {
	Port     uint16   `env:"PORT"`
	Database Database `           prefix:"DB_"`
}

func Load() (Config, error) {
	return env.ParseAsWithOptions[Config](env.Options{
		PrefixTagName:       "prefix",
		DefaultValueTagName: "default",
	})
}
