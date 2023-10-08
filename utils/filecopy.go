package utils

import (
	"io"
	"os"
)

func CopyFile(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return &os.PathError{Op: "CopyFile", Path: src, Err: os.ErrInvalid}
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destinition, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinition.Close()

	_, err = io.Copy(destinition, source)
	return err
}
