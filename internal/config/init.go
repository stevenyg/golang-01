package config

type Config struct {
	Env EnvConfig
	API APIConfig
}

type module struct {
	config Config
}

type EnvConfig struct {
	Env         string
	ServiceName string
}

type APIConfig struct {
	Port string
}

func New() *module {
	return &module{}
}
