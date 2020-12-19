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

## 8章 Maps
- mapは参照型なのでnil値になる可能性がある
  - nilのマップは読み取り時には正常に動くが、書き込み時にpanicとなる

```go
// nilのマップとなるので原則利用しないのが良さそう
var m map[string] string

// 以下のいずれかの方法で、初期化する場合は空のマップを作成する
var dictionary = map[string]string{}
var dictionary = make(map[string]string)
```

- mapの詳細 https://blog.golang.org/maps

- [エラーの定数化について](https://dave.cheney.net/2016/04/07/constant-errors)

## 9章 依存性注入
