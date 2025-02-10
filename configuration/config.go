package configuration

// NOTE(wholesomeow): Used https://zhwt.github.io/yaml-to-go/ to auto-convert lol
type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		DBName        string   `yaml:"db-name"`
		Username      string   `yaml:"user"`
		Password      string   `yaml:"password"`
		SSLMode       string   `yaml:"sslmode"`
		CSVPath       string   `yaml:"csv-path"`
		JSONPath      string   `yaml:"json-path"`
		MigrationPath string   `yaml:"migration-path"`
		RequiredFiles []string `yaml:"req-files"`
		OptionalFiles []string `yaml:"optional-files"`
	} `yaml:"database"`
}

// NOTE(wholesomeow): Use https://mholt.github.io/json-to-go/ to auto-convert too lol
