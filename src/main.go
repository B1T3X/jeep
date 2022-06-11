package main

import (
	"os"
	"github.com/go-yaml/yaml"

)

type jeepConfig struct {
	PrinterPath string `yaml:"printerPath", envconfig:"JEEP_PRINTER_PATH"`
	JiraAddress string `yaml:"jiraAddress", envconfig:JEEP_JIRA_ADDRESS`
	HttpsConfig struct {
	Port string `yaml:"port", envconfig:"JEEP_HTTPS_PORT"`
	CertificatePath string `yaml:"certificatePath", envconfig:"JEEP_CERTIFICATE_PATH"`
	PrivateKeyPath string `yaml:"privateKeyPath", envconfig:"JEEP_PRIVATE_KEY_PATH"`
	} `yaml:"https"`
}

func readConfig(filePath string) (config jeepConfig, err error) {
	data, err := os.ReadFile(filePath)
	yaml.Unmarshal(data, &config)
	return
}


func main() {
	config, err := readConfig("./config/config.yaml")
	if err != nil {
		panic(err)
	}
	config.runServerWithConfig()
}
