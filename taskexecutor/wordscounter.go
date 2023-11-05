package taskexecutor

import (
	"container/heap"
	"fmt"
	"os"
	"sync"
	"txtfile-reader/utils"
)

type FileWordExecutor struct {
	IFileProcessor
	dataChan chan processData
	wg       sync.WaitGroup
}

func (exec *FileWordExecutor) Execute(files []string) error {
	wordFreqMap := make(map[string]int)
	for _, filePath := range files {
		exec.wg.Add(1)
		go exec.readAndProcessFile(filePath)
	}
	go exec.asyncHandleChan()
	for data := range exec.dataChan {
		wordFreqMap[data.name] += data.count
	}
	arr := &heapData{}
	heap.Init(arr)
	for key, val := range wordFreqMap {

		heap.Push(arr, processData{key, val})
		if arr.Len() > 10 {
			heap.Pop(arr)
		}
	}
	exec.print(arr)
	return nil
}

func (exec *FileWordExecutor) print(data *heapData) {
	var curr processData
	fmt.Printf("%s Top 10 occured words in this directory. %s\n", utils.Cyellow, utils.Creset)
	for data.Len() > 0 {
		curr = heap.Pop(data).(processData)
		fmt.Printf("Word %s%v%s is occured %s%v%s times.\n", utils.Cblue, curr.name, utils.Creset, utils.Cgreen, curr.count, utils.Creset)
	}
}
func (exec *FileWordExecutor) asyncHandleChan() {
	exec.wg.Wait()
	close(exec.dataChan)
}
func (exec *FileWordExecutor) readAndProcessFile(path string) {
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
	data, ok := result.(map[string]int)
	if !ok {

	}
	for key, val := range data {
		exec.dataChan <- processData{
			name:  key,
			count: val,
		}
	}
	// return nil
}
