package config

type SrvConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Mode string `yaml:"mode"`
}

func (s *SrvConfig) ParseUrl() string {
	return s.Host + ":" + s.Port
}
