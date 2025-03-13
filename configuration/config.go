package configuration

// NOTE(wholesomeow): Used https://zhwt.github.io/yaml-to-go/ to auto-convert lol
type Config struct {
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

// NOTE(wholesomeow): Use https://mholt.github.io/json-to-go/ to auto-convert too lol
