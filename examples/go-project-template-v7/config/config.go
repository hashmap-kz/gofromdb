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
	Server   ServerConfig   `json:"server"`
	Postgres PostgresConfig `json:"postgres"`
	Logger   LoggerConfig   `json:"logger"`
}

type ServerConfig struct {
	Port         string `json:"port,omitempty"`
	ReadTimeout  string `json:"read_timeout,omitempty"`  // time.Duration
	WriteTimeout string `json:"write_timeout,omitempty"` // time.Duration
}

type LoggerConfig struct {
	Format string `json:"format,omitempty"`
	Level  string `json:"level,omitempty"`
}

type PostgresConfig struct {
	ConnStr string `json:"conn_str,omitempty"`
	PoolMax int    `json:"pool_max,omitempty"`
}

func FromFile(filename string) *Config {
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
