package model

var DATABASE string
var CONFIG string

type Status struct {
	IsMaster   bool
	IsWritable bool
	Error      error
}

type MysqlRows struct {
	SlaveIORunning      string
	SlaveSQLRunning     string
	SecondsBehindMaster string
}

type Config struct {

	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`

	Pgsql struct {
		Host           string `yaml:"host"`
		User           string `yaml:"user"`
		Password       string `yaml:"password"`
		Port           int    `yaml:"port"`
		Database       string `yaml:"database"`
		ConnectTimeout string `yaml:"timeout"`
	} `yaml:"pgsql"`

	Mysql struct {
		Host           string `yaml:"host"`
		User           string `yaml:"user"`
		Password       string `yaml:"password"`
		Port           int    `yaml:"port"`
		Database       string `yaml:"database"`
		ConnectTimeout string `yaml:"timeout"`
	} `yaml:"mysql"`
}