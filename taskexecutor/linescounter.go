package taskexecutor

import (
	"fmt"
	"os"
	"path"
	"sync"
	"txtfile-reader/utils"
)

type FileLineExecutor struct {
	IFileProcessor
	dataChan chan processData
	wg       sync.WaitGroup
}

func (exec *FileLineExecutor) Execute(files []string) error {
	for _, filePath := range files {
		exec.wg.Add(1)
		go exec.readAndProcessFile(filePath)
	}
	go exec.asyncHandleChan()
	for data := range exec.dataChan {
		exec.print(data)
	}
	return nil
}
func (exec *FileLineExecutor) print(data processData) {
	fmt.Printf("%s%v%s contains %s%v%s lines.(Dir Path: %v)\n", utils.Cblue, path.Base(data.name), utils.Creset, utils.Cgreen, data.count, utils.Creset, path.Dir(data.name))
}
func (exec *FileLineExecutor) asyncHandleChan() {
	exec.wg.Wait()
	close(exec.dataChan)
}
func (exec *FileLineExecutor) readAndProcessFile(path string) {
	defer exec.wg.Done()
	openFile, err := os.Open(path)
	if err != nil {
		// return err
	}
	defer openFile.Close()
	result, err := exec.Process(openFile)
	if err != nil {
		// return err
	}
	data, ok := result.(int)
	if !ok {

	}
	exec.dataChan <- processData{
		name:  path,
		count: data,
	}
	// return nil
}
