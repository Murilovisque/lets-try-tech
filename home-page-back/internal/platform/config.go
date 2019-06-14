package platform

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
)

func LoadConfigFromJSONFile(filename string, configType interface{}) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.Wrapf(err, "Error loadding file %s", filename)
	}
	return errors.Wrapf(json.Unmarshal(b, configType), "Error parsing json file %s to %v", filename, configType)
}
