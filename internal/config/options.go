package config

type Options struct {
	Env string `yaml:"Env"`

	SecKey string `yaml:"SecKey"`

	Server struct {
		Port int `yaml:"Port"`
	} `yaml:"Server"`

	DataBase struct {
		Host     string `yaml:"Host"`
		Port     int    `yaml:"Port"`
		DBName   string `yaml:"DBName"`
		Username string `yaml:"Username"`
		Password string `yaml:"Password"`
	} `yaml:"Database"`

	Logger struct {
		LogPath string `yaml:"LogPath"`
	} `yaml:"Logger"`
}
