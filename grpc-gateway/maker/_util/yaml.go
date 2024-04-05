package _util

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

func YamlReader(path string, data any) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Failed to read YAML file: %v", err)
		return
	}
	err = yaml.Unmarshal(yamlFile, data)
	if err != nil {
		log.Printf("Failed to unmarshal YAML data: %v", err)
		return
	}
}

func YamlWriter(filename string, data any) error {
	yamlData, err := yaml.Marshal(&data)
	if err != nil {
		return fmt.Errorf("failed to marshal YAML data: %v", err)
	}

	err = ioutil.WriteFile(filename, yamlData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write YAML file: %v", err)
	}
	fmt.Println("Yaml written to", filename)
	return nil
}
