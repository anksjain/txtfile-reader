package taskexecutor

import (
	"os"
	"sync"
	processor "txtfile-reader/fileProcessor"
)

const (
	CLineCounter string = "line"
	CWordCounter string = "word"
)

type IExecutor interface {
	Execute(files []string) error
}
type IFileProcessor interface {
	Process(file *os.File) (interface{}, error)
}
type processData struct {
	name  string
	count int
}

func NewTaskExecutor(processType string, chanSize int) IExecutor {
	if processType == CLineCounter {
		return &FileLineExecutor{
			&processor.LineCounter{},
			make(chan processData, chanSize),
			sync.WaitGroup{},
		}
	}
	return &FileWordExecutor{
		&processor.WordCounter{},
		make(chan processData, chanSize),
		sync.WaitGroup{},
	}
}
