package config

import (
	"log"
	"os"
	"sync"

	"sigs.k8s.io/yaml"
)

var (
	once   sync.Once
	config *Config
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Logger   LoggerConfig
}

type ServerConfig struct {
	Port         string
	ReadTimeout  string // time.Duration
	WriteTimeout string // time.Duration
}

type LoggerConfig struct {
	Format string
	Level  string
}

type PostgresConfig struct {
	URL     string
	PoolMax int
}

type RedisConfig struct {
	Addr     string
	Database int
	TTL      string // time.Duration
}

func LoadConfigFromFile(filename string) *Config {
	once.Do(func() {
		content, err := os.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		content = expandEnvVars(content)

		var cfg Config
		err = yaml.Unmarshal(content, &cfg)
		if err != nil {
			log.Fatal(err)
		}
		config = &cfg
	})

	return config
}

func expandEnvVars(buf []byte) []byte {
	s := string(buf)
	e := os.ExpandEnv(s)
	return []byte(e)
}

func Cfg() *Config {
	if config == nil {
		log.Fatal("config was not loaded in main")
	}
	return config
}
