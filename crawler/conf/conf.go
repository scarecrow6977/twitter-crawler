package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

var lock sync.Mutex
var configFile string

type PostgresAccessConfig struct {
	Host     string  `yaml:"host"`
	Port     *uint16 `yaml:"port,omitempty"`
	Dbname   string  `yaml:"dbname"`
	User     string  `yaml:"user"`
	Password *string `yaml:"password,omitempty"`
}

type Neo4jAccessConfig struct {
	Uri      string `yaml:"uri"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type MasterConfig struct {
	NumOfWorkers       int                  `yaml:"num_of_workers"`
	QueueSize          int                  `yaml:"queue_size"`
	QueueNoRefillLimit int                  `yaml:"queue_no_refill_limit"`
	ApiLimitTimeout    int                  `yaml:"api_limit_timeout"`
	PostgresAccess     PostgresAccessConfig `yaml:"pg_access"`
	Neo4jAccess        Neo4jAccessConfig    `yaml:"neo4j_access"`
	Cookies            map[string]string    `yaml:"cookies"`
	Headers            map[string]string    `yaml:"headers"`
}

func Init(configFilePath string) {
	configFile = configFilePath
}

func LoadConfig() (*MasterConfig, error) {
	lock.Lock()
	defer lock.Unlock()
	yamlBytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	config := &MasterConfig{}
	err = yaml.Unmarshal(yamlBytes, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
