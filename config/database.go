package config

type DBConfig struct {
	Dialect   string `yaml:"dialect"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	Name      string `yaml:"name"`
	Charset   string `yaml:"charset"`
	ParseTime string `yaml:"parse_time"`
	Loc       string `yaml:"loc"`
}

func (d *DBConfig) ParseUrl() string {
	return d.Username + ":" + d.Password +
		"@tcp(" + d.Host + ":" + d.Port + ")/" +
		d.Name + "?charset=" + d.Charset +
		"&parseTime=" + d.ParseTime + "&loc=" + d.Loc
}
