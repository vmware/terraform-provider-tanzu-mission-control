// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package helper

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	k8sYaml "sigs.k8s.io/yaml"
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

// ReadYamlFileAsJSON reads a yaml file from a given path.
func ReadYamlFileAsJSON(filePath string) (string, error) {
	bufString, err := ReadYamlFile(filePath)
	if err != nil {
		return "", err
	}

	jsonBytes, err := k8sYaml.YAMLToJSON([]byte(bufString))
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

// WriteYamlFile writes a yaml file to a given path.
func WriteYamlFile(filePath string, data interface{}) error {
	outputFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("Error opening or creating the %s file.", filePath))
	}

	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			log.Println("[ERROR] could not close file object.")
		}
	}(outputFile)

	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	_, err = outputFile.Write(yamlData)
	if err != nil {
		return err
	}

	return nil
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
