package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
	"txtfile-reader/taskexecutor"
	"txtfile-reader/utils"
)

var taskPtr = flag.String("task", "line", "Choose task 'line' or 'word'. \n Choosing 'line' will read files and number of lines in each file.\n Choosing 'word' will read files and print top 10 frequent word in files count. ")
var dirPathPtr = flag.String("d", ".", "The directory to scan for text files")
var fileExtPtr = flag.String("ext", ".txt", "Provide the comma seprated file extension to read data")

var textFileMap = map[string]struct{}{}

func main() {
	flag.Parse()
	loadFileExt(*fileExtPtr)
	txtFilesPath, err := getTextFilesPath(*dirPathPtr)
	if err != nil {
		panic(err)
	}
	if len(txtFilesPath) == 0 {
		fmt.Printf("%s There no text files found in the directory", utils.Cyellow)
		return
	}
	taskExecutor := taskexecutor.NewTaskExecutor(*taskPtr, 10000)
	err = taskExecutor.Execute(txtFilesPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Task Completed SuccessFully......")
}
func loadFileExt(fileExts string) {
	exts := strings.Split(fileExts, ",")
	for _, ext := range exts {
		textFileMap[ext] = struct{}{}
	}
}
func getTextFilesPath(dirPath string) ([]string, error) {
	res := make([]string, 0)
	all, err := os.ReadDir(dirPath)
	if err != nil {
		return res, err
	}
	for _, kind := range all {
		if kind.IsDir() {
			data, err := getTextFilesPath(dirPath + "/" + kind.Name())
			if err != nil {
				return res, err
			}
			res = append(res, data...)
			continue
		}
		if _, exist := textFileMap[path.Ext(kind.Name())]; exist {
			res = append(res, dirPath+"/"+kind.Name())
		}
	}
	return res, nil
}
