package tools

import (
	"encoding/hex"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	err := os.ErrNotExist
	assert.False(t, Check(err))

	noErr := error(nil)
	assert.True(t, Check(noErr))
}

func TestExists(t *testing.T) {
	tempFile, _ := os.CreateTemp("", "example")
	defer os.Remove(tempFile.Name())

	assert.True(t, Exists(tempFile.Name()))
	assert.False(t, Exists("nonexistent.file"))
}

func TestReadFile(t *testing.T) {
	type Sample struct {
		Name string `json:"name"`
	}

	content := Sample{Name: "Test"}
	contentBytes, _ := json.Marshal(content)

	tempFile, _ := os.CreateTemp("", "example")
	defer os.Remove(tempFile.Name())
	err := os.WriteFile(tempFile.Name(), []byte(hex.EncodeToString(contentBytes)), 0644)
	assert.Nil(t, err)

	var decodedContent Sample
	err = ReadFile(tempFile.Name(), &decodedContent)
	assert.Nil(t, err)
	assert.Equal(t, content, decodedContent)
}

func TestWriteFile(t *testing.T) {
	type Sample struct {
		Name string `json:"name"`
	}

	content := Sample{Name: "WriteTest"}
	tempFile, _ := os.CreateTemp("", "write_example")
	defer os.Remove(tempFile.Name())

	err := WriteFile(tempFile.Name(), content)
	assert.Nil(t, err)

	fileContent, _ := os.ReadFile(tempFile.Name())
	decodedContentBytes, _ := hex.DecodeString(string(fileContent))

	var decodedContent Sample
	err = json.Unmarshal(decodedContentBytes, &decodedContent)
	assert.Nil(t, err)

	assert.Equal(t, content, decodedContent)
}
