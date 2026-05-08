package config

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
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
	Redis    RedisConfig
	Auth     KeycloakConfig
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

type KeycloakConfig struct {
	URL          string
	Realm        string
	ClientID     string
	ClientSecret string
}

type RedisConfig struct {
	Addr     string
	Database int
	TTL      string // time.Duration
}

// LoadConfigFromFile: unmarshal file into config struct
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

		errors := validateStruct(config)
		if len(errors) > 0 {
			for _, e := range errors {
				log.Printf("%v", e)
			}
			// log.Fatalln("config-file is not completely set")
		}
	})

	return config
}

// LoadConfig: unmarshal raw data into config struct
func LoadConfig(content []byte) *Config {
	once.Do(func() {
		content = expandEnvVars(content)

		var cfg Config
		err := yaml.Unmarshal(content, &cfg)
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

// validateStruct collects all unset fields in a struct
func validateStruct(s interface{}) []error {
	var errors []error

	v := reflect.ValueOf(s)

	// Ensure the input is a struct or pointer to a struct
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return []error{
			fmt.Errorf("validateStruct: input must be a struct or pointer to a struct: %v", v),
		}
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)

		// Recursively validate nested structs
		if field.Kind() == reflect.Struct {
			nestedErrors := validateStruct(field.Interface())
			for _, err := range nestedErrors {
				errors = append(errors, fmt.Errorf("%s.%s", fieldType.Name, strings.TrimPrefix(err.Error(), ".")))
			}
			continue
		}

		// Check if field is zero value
		if reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()) {
			errors = append(errors, fmt.Errorf("%s is not set (zero value)", fieldType.Name))
		}
	}

	return errors
}
