package config

import (
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

type Config struct {
	// NOTE(wholesomeow): Used https://zhwt.github.io/yaml-to-go/ to auto-convert lol
	Server struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Mode     string `yaml:"mode"`
		Loglevel string `yaml:"loglevel"`
		Network  string `yaml:"network"`
	} `yaml:"server"`
	Database struct {
		DBName        string `yaml:"dbname"`
		Hostname      string `yaml:"hostname"`
		Username      string `yaml:"user"`
		Password      string `yaml:"password"`
		Port          int    `yaml:"port"`
		SSLMode       string `yaml:"sslmode"`
		CSVPath       string `yaml:"csvpath"`
		JSONPath      string `yaml:"jsonpath"`
		MigrationPath string `yaml:"migrationpath"`
		Files         []struct {
			Filename  string   `yaml:"filename"`
			Required  bool     `yaml:"required"`
			Header    bool     `yaml:"header"`
			Tablename string   `yaml:"tablename"`
			Schema    []string `yaml:"schema"`
		} `yaml:"files"`
	} `yaml:"database"`
}

var (
	instance *Config
	once     sync.Once
)

// Returns the singleton instance of the configuration file
func ReadConfig(path string) (*Config, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Open files and decode into the singleton instance
	log.Printf("reading %s file", path)
	yaml_decoder := yaml.NewDecoder(f)
	once.Do(func() {
		instance = &Config{}
		yaml_decoder.Decode(instance)
	})
	return instance, nil
}
