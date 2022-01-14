package operator

import (
	"fmt"
	"io/ioutil"
)

func ReadFile(filePath string) (int, []byte) {
	fileFullPath := filePath
	fmt.Println(fileFullPath)
	data, err := ioutil.ReadFile(fileFullPath)
	//fmt.Println(len(data))
	if err != nil {
		return 0, nil
	}
	return len(data), data
}
