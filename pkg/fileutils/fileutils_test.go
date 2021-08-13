package fileutils

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteSliceToFile(t *testing.T) {
	file, err := ioutil.TempFile("", "test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())

	WriteSliceToFile(file.Name(), []string{"test1", "test2"})

	data, err := ioutil.ReadFile(file.Name())
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, "test1\ntest2\n", string(data))
}

func TestWriteSliceToFileWithWrongPath(t *testing.T) {
	file, err := ioutil.TempFile("", "test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())

	err = WriteSliceToFile("wrongPath", []string{"test1", "test2"})

	assert.Error(t, err)

	data, err := ioutil.ReadFile(file.Name())
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, "", string(data))
}
