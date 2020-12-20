package di

import (
	"bytes"
	"fmt"
)

// この関数のテストは難しい
// 標準出力に表示されたテキストをテストフレームワークでキャプチャするのは困難
func GreetSample(name string) {
	fmt.Printf("Hello, %s", name)
}

func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
