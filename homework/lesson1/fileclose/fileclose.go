package fileclose

import (
	"os"
)


// FileClose description (for linter)
func FileClose(filename string) error {
	f,err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return f.Close()
}
