package fileutils

import "os"

func WriteSliceToFile(path string, slice []string) error {

	writer := func(path string, data string) (err error) {
		file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}

		defer file.Close()

		_, err = file.WriteString(data + "\n")
		if err != nil {
			return err
		}

		return nil
	}

	for _, data := range slice {
		err := writer(path, data)
		if err != nil {
			return err
		}
	}

	return nil
}
