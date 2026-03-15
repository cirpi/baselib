package sql

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Pass     string `yaml:"pass"`
	Database string `yaml:"database"`
}
