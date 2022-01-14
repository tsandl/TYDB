package operator

import (
	"fmt"
	"testing"
)

func TestFile(t *testing.T) {
	folderPath := "F:\\研究生\\研一下"
	//list, err := getDirList("F:\\研究生\\研一下")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for _, v := range list {
	//	fmt.Println(v)
	//}
	var s []string
	//str := getFiles(folderPath,s)
	for index, value := range getFiles(folderPath, s) {
		fmt.Println(index, value)
	}

}
