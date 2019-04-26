package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Todo struct {
	Time time.Time
	Body string
}

func (self *Todo) save() ([]byte, error) {
	filename := fmt.Sprint("files/", time.Now().Unix(), ".txt")
	jsoned, _ := json.Marshal(self)
	err := ioutil.WriteFile(filename, jsoned, 0600)
	if err != nil {
		return []byte{}, err
	}

	return jsoned, nil
}
