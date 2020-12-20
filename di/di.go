package di

import (
	"fmt"
	"io"
)

// この関数のテストは難しい
// 標準出力に表示されたテキストをテストフレームワークでキャプチャするのは困難
func GreetSample(name string) {
	fmt.Printf("Hello, %s", name)
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
