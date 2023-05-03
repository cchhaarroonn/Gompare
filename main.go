package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
)

type files struct {
	Name string
	Hash string
}

func main() {
	var sourceDir string = "SOURCE DIR PATH"
	var targetDir string = "TARGET DIR PATH"

	var directories []files

	errSource := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("[X] Failed walking path: " + path)
		}

		fileName := info.Name()
		fileHash := md5.Sum([]byte(fileName))
		fileEncoded := hex.EncodeToString(fileHash[:])

		directories = append(directories, files{Name: fileName, Hash: fileEncoded})

		return nil
	})
	if errSource != nil {
		fmt.Println("[X] Can't read directory: ", sourceDir)
	}

	errTarget := filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("[X] Failed walking path: " + path)
			return nil
		}

		fileName := info.Name()
		fileHash := md5.Sum([]byte(fileName))
		fileEncoded := hex.EncodeToString(fileHash[:])

		found := false
		for _, element := range directories {
			if element.Name == fileName {
				found = true

				if element.Hash != fileEncoded {
					fmt.Println("[X] File mismatch: ", fileName)
				}

				break
			}
		}

		if !found {
			fmt.Println("[X] Missing file: ", fileName)
		}

		return nil
	})
	if errTarget != nil {
		fmt.Println("[X] Can't read directory: ", sourceDir)
	}

	fmt.Print("Successfully compared!")
}
