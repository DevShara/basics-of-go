package fileutils

import "os"

func WriteToFile(filename string, content string) error {
	 os.WriteFile(filename, []byte(content), 0644)
	 return nil
}