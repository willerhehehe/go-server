package config

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func InitFromFile(filePath string, fileType string, resultStruct interface{}) error {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("read local config fail: %v", err.Error())
	}
	return InitFromContent(fileType, fileContent, resultStruct)
}

func InitFromContent(confType string, content []byte, resultStruct interface{}) error {
	switch confType {
	case "xml":
		return xml.Unmarshal(content, resultStruct)
	case "json":
		return json.Unmarshal(content, resultStruct)
	case "yaml":
		return yaml.Unmarshal(content, resultStruct)
	default:
		return fmt.Errorf("config type must be one of: xml, json, yaml, '%s' can't recognize", confType)
	}
}
