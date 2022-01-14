package operator

import "testing"

func TestWriteFile(t *testing.T) {
	destinationFilePath := "F:\\code\\go_code\\db1\\file\\result\\12.jpg"
	originFilePath := "F:\\huajianmodel\\data_model\\picture\\2.jpg"
	_, data := ReadFile(originFilePath)
	WriteFile(destinationFilePath, data)
}
