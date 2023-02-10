package configurer

import (
	"bytes"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func Load(configPath string, to interface{}) error {
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	content = []byte(os.ExpandEnv(string(content)))

	d := yaml.NewDecoder(bytes.NewReader(content))
	return d.Decode(to)
}
