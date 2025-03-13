package configuration

// NOTE(wholesomeow): Used https://zhwt.github.io/yaml-to-go/ to auto-convert lol
type Config struct {
	Server struct {
		Port     int    `yaml:"port"`
		Host     string `yaml:"host"`
		Mode     string `yaml:"mode"`
		LogLevel string `yaml:"loglevel"`
		Network  string `yaml:"network"`
	} `yaml:"server"`
	Database struct {
		DBName        string   `yaml:"dbname"`
		Hostname      string   `yaml:"hostname"`
		Username      string   `yaml:"user"`
		Password      string   `yaml:"password"`
		Port          int      `yaml:"port"`
		SSLMode       string   `yaml:"sslmode"`
		CSVPath       string   `yaml:"csvpath"`
		JSONPath      string   `yaml:"jsonpath"`
		MigrationPath string   `yaml:"migrationpath"`
		RequiredFiles []string `yaml:"reqfiles"`
		OptionalFiles []string `yaml:"optionalfiles"`
	} `yaml:"database"`
}

// NOTE(wholesomeow): Use https://mholt.github.io/json-to-go/ to auto-convert too lol
