package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// この関数のテストは難しい
// 標準出力に表示されたテキストをテストフレームワークでキャプチャするのは困難
func GreetSample(name string) {
	fmt.Printf("Hello, %s", name)
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	Greet(os.Stdout, "tomo")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
