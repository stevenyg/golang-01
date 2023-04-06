package config

func (m *module) ApplyConfig(c Config) {
	m.config = c
}

func (m *module) GetConfig() Config {
	return m.config
}
