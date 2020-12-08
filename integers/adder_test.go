package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// godoc -http=:6060 でdocに例を追加

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output:6
}
