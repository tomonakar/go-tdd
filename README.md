# Go-TDD

https://andmorefine.gitbook.io/learn-go-with-tests/

## TableDrivenTests 
- [TableDrivenTests](https://github.com/golang/go/wiki/TableDrivenTests)
同じ方法でテストできるテストケースのリストを作成するのに役立つ。


## Pointers
- [ポインタの逆参照の仕様](https://golang.org/ref/spec#Method_values)

## 未チェックのエラーを検出
```go
go get -u github.com/kisielk/errcheck
```

```sh
errcheck .
```

## エラーはチェックするだけでなく適切に処理する
https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully