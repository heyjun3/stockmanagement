package ikebe

import (
	"fmt"

	"github.com/BurntSushi/toml"
)



type Config struct {
	Psql Psql
	output string `toml:"output"`
}

type Psql struct {
	DBname string `toml:"dbname"`
	Host string `toml:"host"`
	Port string `toml:"port"`
	User string `toml:"user"`
	Pass string `toml:"pass"`
	SSLmode string `toml:"sslmode"`
	Blacklist []string `toml:"blacklist"`
}

var cfg Config
func init() {
	cfg = NewConfig()
}

func NewConfig() Config {
	var cfg Config
	_, err := toml.DecodeFile("sqlboiler.toml", &cfg)
	if err != nil {
		fmt.Println(err)
		return cfg
	}
	return cfg
}

func (c Config) dsn() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.Psql.User,
		c.Psql.Pass,
		c.Psql.Host,
		c.Psql.Port,
		c.Psql.DBname,
		c.Psql.SSLmode,
	)
}