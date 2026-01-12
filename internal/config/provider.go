package config

func ProvideConfig() (*Config, error) {
	return LoadConfig()
}
