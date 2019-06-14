package platform

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfigFromJSONFileShouldWorks(t *testing.T) {
	const json = "{ \"f1\": \"val1\", \"f2\": 1 }"
	file, err := ioutil.TempFile("", t.Name())
	assert.Nil(t, err, "Error to create tempfile")
	defer os.Remove(file.Name())
	_, err = file.WriteString(json)
	assert.Nil(t, err, "Error to create tempfile")
	configType := struct {
		F1 string
		F2 int
	}{}
	err = LoadConfigFromJSONFile(file.Name(), &configType)
	assert.Nil(t, err, "Error loading config")
	assert.Equal(t, configType.F1, "val1")
	assert.Equal(t, configType.F2, 1)
}

func TestLoadConfigFromJSONFileShouldFailWhenFileDoesNotExists(t *testing.T) {
	tempDir, err := ioutil.TempDir("", t.Name())
	assert.Nil(t, err, "Error to create tempDir")
	configType := struct {
		F1 string
		F2 int
	}{}
	err = LoadConfigFromJSONFile(tempDir+string(os.PathSeparator)+t.Name(), &configType)
	assert.NotNil(t, err, "Test must failed")
}

func TestLoadConfigFromJSONFileShouldFailWhenJSONParsingFailed(t *testing.T) {
	const json = "{ \"f1\": \"val1\", \"f2\": \"val2\" }"
	file, err := ioutil.TempFile("", t.Name())
	assert.Nil(t, err, "Error to create tempfile")
	defer os.Remove(file.Name())
	_, err = file.WriteString(json)
	assert.Nil(t, err, "Error to create tempfile")
	configType := struct {
		F1 string
		F2 int
	}{}
	err = LoadConfigFromJSONFile(file.Name(), &configType)
	assert.NotNil(t, err, "Test must failed")
}
