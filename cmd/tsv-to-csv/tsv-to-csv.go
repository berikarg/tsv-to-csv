package main

import (
	"io/ioutil"
	"os"
	"strings"

	"go.uber.org/zap"
)

func main() {

	logger := zap.NewExample().Sugar()
	defer logger.Sync()

	inputDir := "../../assets/in/"
	outputDir := "../../assets/out/"
	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		logger.Fatal("cannot read dir")
	}
	for _, file := range files {
		data, err := os.ReadFile(inputDir + file.Name())
		if err != nil {
			logger.Error("cannot read file")
		}
		for i, symbol := range data {
			if symbol == 9 {
				data[i] = 44
			}
		}
		err = os.WriteFile(outputDir + strings.TrimRight(file.Name(), ".tsv") + ".csv", data, 0644)
		if err != nil {
			logger.Error("cannot write file")
		}
	}
}

