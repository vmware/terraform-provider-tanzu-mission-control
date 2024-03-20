/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package helper

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// ReadYamlFile reads a yaml file from a given path.
func ReadYamlFile(filePath string) (string, error) {
	inputFile, err := os.Open(filePath)
	if err != nil {
		return "", errors.WithMessage(err, fmt.Sprintf("Error opening the %s file.", filePath))
	}

	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			log.Println("[ERROR] could not close file object.")
		}
	}(inputFile)

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, inputFile)

	if err != nil {
		return "", err
	}

	_, err = yaml.Marshal(buf.String())
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// FileExists checks if a file exists in a given path.
func FileExists(filePath string) bool {
	fileInfo, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		log.Println("[ERROR] file does not exists.")
		return false
	}

	return !fileInfo.IsDir()
}
