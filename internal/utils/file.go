package utils

import (
	"bytes"
	"os"
)

func SaveDataToFile(filePath string, data []byte, repeat uint) error  {
	buf := new(bytes.Buffer)
	for ; repeat > 0; repeat-- {
		buf.Write(data)
		buf.WriteByte('\n')
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			return err
		}

		defer func(file *os.File) {
			file.Close()
		}(file)

		if _, err := file.Write(buf.Bytes()); err != nil {
			return err
		}
	}

	return nil
}