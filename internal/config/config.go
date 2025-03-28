package config

type Config struct {
	PostgresDSN string
}

func New(file string) Config {
	// TODO распарсить yaml файл, при его наличии
	return Config{}
}
