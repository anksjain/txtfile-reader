package processor

import (
	"bufio"
	"os"
)

type LineCounter struct{}
type WordCounter struct{}

func (c *LineCounter) Process(file *os.File) (interface{}, error) {
	lineCount := 0
	var err error
	newLineScanner := bufio.NewScanner(file)
	for newLineScanner.Scan() {
		lineCount++
	}
	err = newLineScanner.Err()
	if err != nil {
		return 0, err
	}
	return lineCount, err
}
func (c *WordCounter) Process(file *os.File) (interface{}, error) {
	mp := make(map[string]int)
	var err error
	newWordScanner := bufio.NewScanner(file)
	newWordScanner.Split(bufio.ScanWords)
	for newWordScanner.Scan() {
		mp[newWordScanner.Text()]++
	}
	err = newWordScanner.Err()
	if err != nil {
		return nil, err
	}
	return mp, err
}
