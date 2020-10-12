package config

// Configuration settings
type Configuration struct {
	Environment string `yaml:"environment" default:"local"`
	Database    string `yaml:"database" default:"memorydb"`

	Debug struct {
		Enable bool `yaml:"enable" default:"false" comment:"allow debug mode"`
	} `yaml:"debug"`

	Logger struct {
		LogLevel string `yaml:"loglevel" default:"info"`
		Encoding string `yaml:"encoding" default:"console"`
		Sentry   string `yaml:"sentry" default:"dsn"`
	} `yaml:"logger"`



	Psql struct {
		AutoMigrate bool   `yaml:"-" default:"false"`
		Host        string `yaml:"host" default:"localhost:5432"`
		Username    string `yaml:"username" default:"golang"`
		Database    string `yaml:"database" default:"db"`
		Password    string `yaml:"password" default:""`
	} `yaml:"psql"`

	GraphQL struct {
		Host string `yaml:"host" default:"0.0.0.0" comment:"Listen address for server"`
		Port string `yaml:"port" default:"8080" comment:"Listen port for server"`
	} `yaml:"server"`
}
