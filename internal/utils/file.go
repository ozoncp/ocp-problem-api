package utils

import (
	"bytes"
	"os"
)

func SaveDataToFile(filePath string, data []byte, repeat uint) error  {
	buf := new(bytes.Buffer)
	saveData := func(filePath string, data []byte) (err error) {
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			return
		}

		defer func(file *os.File, err error) {
			err = file.Close()
		}(file, err)

		if _, err = file.Write(buf.Bytes()); err != nil {
			return
		}

		return
	}

	for ; repeat > 0; repeat-- {
		buf.Write(data)
		buf.WriteByte('\n')

		if err := saveData(filePath, buf.Bytes()); err != nil {
			return err
		}
	}

	return nil
}
