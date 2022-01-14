package operator

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	l, data := ReadFile("F:\\huajianmodel\\data_model\\picture\\2.jpg")
	fmt.Println(l, data)

}
