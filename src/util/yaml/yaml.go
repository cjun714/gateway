package yaml

import (
	yaml "gopkg.in/yaml.v2"
)

// ToYaml is used to serialize object to yaml byte[]
func ToYaml(obj interface{}) ([]byte, error) {
	bts, e := yaml.Marshal(&obj)
	if e != nil {
		return nil, e
	}
	return bts, nil
}

// ToObj is used to unserialize object from yaml byte[]
func ToObj(yamlBytes []byte, obj interface{}) error {
	e := yaml.Unmarshal(yamlBytes, &obj)
	if e != nil {
		return e
	}
	return nil
}
