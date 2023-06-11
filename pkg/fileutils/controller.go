package fileutils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// ReadJsonFile - Loads a json file and returns a map of the contents.
func ReadJsonFile(filePath string) (interface{}, error) {

	// Open our jsonFile
	jsonFile, err := os.Open(filePath)

	// if we os.Open returns an error then handle it
	if err != nil {
		log.Println("fileutils.ReadJsonFile - Error opening configuraiton file.")
		return nil, err
	}
	defer jsonFile.Close()

	// Read the file into a map
	byteResult, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println("fileutils.ReadJsonFile - Error parsing json file.")
		return nil, err
	}

	var jsonMap interface{}
	json.Unmarshal(byteResult, &jsonMap)

	return jsonMap, nil
}

// ReadFileToBytes - Loads a file and returns a byte array with its contents.
func ReadFileToBytes(filePath string) ([]byte, error) {

	// Open our jsonFile
	jsonFile, err := os.Open(filePath)

	// if we os.Open returns an error then handle it
	if err != nil {
		log.Println("fileutils.ReadJsonFile - Error opening configuraiton file.")
		return nil, err
	}
	defer jsonFile.Close()

	// Read the file into a map
	byteResult, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println("fileutils.ReadJsonFile - Error parsing json file.")
		return nil, err
	}

	return byteResult, nil
}
