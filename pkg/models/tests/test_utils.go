package models

import "os"

// GetSampleJsonFiles reads JSON files from a predefined directory and returns their contents.
//
// The function does not take any parameters. It reads all JSON files located in the
// "../../samples/attom/" directory, which are organized in folders that match API endpoints.
//
// It returns a map where the keys are the file names and the values are the contents of the files
// as strings. If there is an error reading the directory or any file, the function will panic.

func GetSampleJsonFiles() map[string]string {

	/// The path where the sample JSON files are located.
	const basePath = "../../samples/attom/"

	// the base path contains json files organized in folders
	// that match the api endpoints, roughly.
	folders, err := os.ReadDir(basePath)

	// We aint gong no where, we need these!
	if err != nil {
		panic(err)
	}

	var jsonFiles = make(map[string]string)

	for _, file := range folders {

		if file.IsDir() {
			continue
		}

		// Read the file and store its contents in a map.
		contents, err := os.ReadFile(basePath + file.Name())

		// this is an all-or-non situation here
		if err != nil {
			panic(err)
		}

		jsonFiles[file.Name()] = string(contents)
	}

	return jsonFiles
}
