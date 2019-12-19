package tasks

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Tasks struct {
	Task []struct {
		Key           string `json:"key"`
		Name          string `json:"name"`
		Description   string `json:"description"`
		ExecutingTime string `json:"executing_time"`
	} `json:"task"`
}

func ParseTask(filename string) (*Tasks, error) {
	var err error
	task := new(Tasks)
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(byteValue, &task); err != nil {
		return nil, err
	}
	return task, nil

}
