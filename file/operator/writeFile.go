package operator

import (
	"fmt"
	"os"
)

func WriteFile(filePath string, data []byte) error {
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	n2, err := f.Write(data)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return err
	}
	fmt.Println(n2, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
