package util

import (
	"os"
)

func CreatFile(name string) (*os.File, error) {
	return os.Create(name)
}
