package sql

type Config struct {
	Host     string `yaml:"host" json:"host" xml:"host"`
	Port     int    `yaml:"port" json:"port" xml:"port"`
	User     string `yaml:"user" json:"user" xml:"user"`
	Pass     string `yaml:"pass" json:"pass" xml:"pass"`
	Database string `yaml:"database" json:"database" xml:"database"`
}
