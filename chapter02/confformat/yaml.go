package confformat

import (
	"bytes"

	"github.com/go-yaml/yaml"
)

type YAMLData struct {
	Name string `yaml:"name"`
	Age int `yaml:"age"`
}

// ToYAML dumps the YAMLData struct to a YAML format bytes.Buffer
func (t *YAMLData) ToYAML() (*bytes.Buffer, error) {
	d, err := yaml.Marshal(t)
	if err != nil {
		return nil, err
	}
	b := bytes.NewBuffer(d)
	return b, nil
}

// Decode will decode to TOMLData
func (t *YAMLData) Decode(data []byte) error {
	return yaml.Unmarshal(data, t)
}