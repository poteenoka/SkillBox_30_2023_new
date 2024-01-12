package config

type config struct {
	HTTP  `yaml:"http"`
	MSSQL `yaml:"mssql"`
}

type HTTP struct {
	Port string `yaml:"port"`
}

type MSSQL struct {
}
