package utils

import (
	"bytes"
	"os"
	"testing"
)

func TestSaveDataToFile(t *testing.T) {
	dataCheck := "Test data!!!\nTest data!!!\nTest data!!!\n"
	data := []byte("Test data!!!")
	fileName := "test.txt"

	if err := SaveDataToFile(fileName, data, 3); err != nil {
		t.Error(err.Error())
	}

	file, err := os.OpenFile(fileName, os.O_RDONLY, 0755)
	if err != nil {
		t.Error(err)
	}

	defer file.Close()
	defer os.Remove(fileName)

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(file); err != nil {
		t.Error(err)
	}

	if buf.String() != dataCheck {
		t.Error("Invalid save data")
	}
}
