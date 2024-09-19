package main

import (
	"fmt"
	"github.com/artcurty/go-proxy-make/internal"
	"github.com/artcurty/go-proxy-make/pkg"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	configDir := pkg.GetEnv("CONFIG_DIR", "./setup")
	outputDir := pkg.GetEnv("OUTPUT_DIR", "./cmd/api/generated")

	files, err := ioutil.ReadDir(configDir)
	if err != nil {
		log.Fatalf("Error reading setup directory: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".yaml" {
			inputFile := filepath.Join(configDir, file.Name())
			err := internal.GenerateProxyFunctionForInput(inputFile, outputDir)
			if err != nil {
				fmt.Printf("Error generating proxy functions for %s: %v\n", inputFile, err)
				continue
			}
			fmt.Printf("File generation completed successfully for %s.\n", inputFile)
		}
	}
}
