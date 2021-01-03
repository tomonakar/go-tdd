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
- コードをテストする関数を簡単にテストできない場合は、通常、依存関係が関数またはグローバルな状態に組み込まれているためです。たとえば、ある種のサービス層で使用されているグローバルデータベース接続プールがある場合、テストが困難になる可能性が高く、実行が遅くなります。DIは、（インターフェイスを介して）データベースの依存関係を挿入するように動機付けし、テストで制御できるものでモックアウトできます。
- 懸念事項を分離して、「データの移動先」と「生成方法」を分離します。メソッド/関数の責任が多すぎると感じた場合は、（データの生成、およびデータベースへの書き込み、HTTPリクエストの処理、およびドメインレベルのロジックの実行）おそらくDIが必要なツールになるでしょう。
- コードをさまざまなコンテキストで再利用できるようにするコードを使用できる最初の「新しい」コンテキストは、テスト内です。しかし、さらに誰かがあなたの関数で何か新しいことを試したい場合、彼らは彼ら自身の依存関係を注入することができます。

## 10章 スタブ・モック
### TDDアプローチの詳細
- ささいな例に直面した場合は、問題を「薄いスライス」に分解してください。ウサギの穴に入り込み、「ビッグバン」アプローチをとらないように、できるだけ早く_testsで動作するソフトウェアを使用できるようにしてください。
- 動作するソフトウェアを入手したら、必要なソフトウェアにたどり着くまで小さなステップで繰り返すのが簡単です。

### モック
- コードの重要な領域をモックしないと、テストされません。私たちのケースでは、各表示の間にコードが一時停止することをテストすることはできませんが、他にも数え切れないほどの例があります。失敗する可能性のあるサービスを呼び出していますか？特定の状態でシステムをテストしたいですか？モックなしでこれらのシナリオをテストすることは非常に困難です。
- モックがないと、単純なビジネスルールをテストするためだけに、データベースや他のサードパーティの設定が必要になる場合があります。テストが遅くなり、フィードバックループが遅くなる可能性があります。
- 何かをテストするためにデータベースまたはWebサービスをスピンアップする必要があるため、そのようなサービスの信頼性が低いために、壊れやすいテストを受ける可能性があります。
- 開発者がモックについて学ぶと、システムの 機能 ではなく 機能 の観点から、システムのあらゆる側面を過剰にテストすることが非常に簡単になります。 テストの価値と、それらが将来のリファクタリングに与える影響について常に注意してください。
- モックに関するこの投稿では、モックの一種であるスパイのみを取り上げました。 モックにはさまざまな種類があります。 ボブおじさんは非常に読みやすい記事でタイプについて説明しています。
- [When to Mock](https://blog.cleancoder.com/uncle-bob/2014/05/10/WhenToMock.html)
- [ボブおじさんの記事](https://blog.cleancoder.com/uncle-bob/2014/05/14/TheLittleMocker.html)

## 11章 並行性
- ベンチマークの実行 `go test -bench=.`
- **blocking**: 通常、Go関数で `doSomething()`を呼び出すと関数が返されるのを待つ（返される値が無い場合でも、関数の終了を待つ）
- [Make It Work Make It Right Make It Fast](http://wiki.c2.com/?MakeItWorkMakeItRightMakeItFast)
- [早期最適化は諸悪の根元](http://wiki.c2.com/?PrematureOptimization)


## 12章 Select
- select
- httptest

## 13章 reflection
- [The Laws of Reflection (go blog)](https://blog.golang.org/laws-of-reflection)

