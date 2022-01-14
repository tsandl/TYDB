package operator

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 读取文件夹下文件目录信息

func getDirList(dirpath string) ([]string, error) {
	var dirList []string
	dirErr := filepath.Walk(dirpath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				dirList = append(dirList, path)
				return nil
			}

			return nil
		})
	return dirList, dirErr
}

// 获取目录下所有文件（包括文件夹中的文件）

func getFiles(folder string, s []string) []string {
	files, _ := ioutil.ReadDir(folder)
	for _, file := range files {
		if file.IsDir() {
			s = getFiles(folder+"\\"+file.Name(), s) // linux是 ‘/’
		} else {
			filePath := folder + "\\" + file.Name()
			//fmt.Println(filePath)
			s = append(s, filePath)
		}
	}
	return s
}

// 获取根目录下直属所有文件（不包括文件夹及其中的文件）

func getAllFile(pathname string, s []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}

	for _, fi := range rd {
		if !fi.IsDir() {
			fullName := pathname + "\\" + fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}
