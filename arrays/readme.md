- カバレッジの取得
```sh
$ go test -cover
```

- 値を比較（型安全を考慮しない）`reflect.DeepEqual`
```go
func TestSumAll(t *testing.T) {

    got := SumAll([]int{1, 2}, []int{0, 9})
    want := "bob"

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want)
    }
}
```

- [The Go Blog Go Slice](https://blog.golang.org/slices-intro)

