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

## 14章 同期
- go vetでバグに関する警告を受け取れる
- sync.Mutexは最初の使用後にコピーしてはならない https://golang.org/pkg/sync/#Mutex
- Mutexを使用するとデータにロックを追加できる
- Waitgroupはゴルーチンがジョブを完了するのを待つ
- channelとmutexの使い分け https://github.com/golang/go/wiki/MutexOrChannel

## 15章 コンテキスト
- [go blog context](https://blog.golang.org/context)
- [context.Valueは使うべきではない](https://faiface.github.io/post/context-should-go-away-go2/)
- [context.Valueは制御ではなく通知](https://medium.com/@cep21/how-to-correctly-use-context-context-in-go-1-7-8f2c0fafdf39)
- [Go2ではコンテキストはなくなるはず](https://faiface.github.io/post/context-should-go-away-go2/)

### 内容
- リクエストがクライアントによってキャンセルされたHTTPハンドラをテストする方法。
- キャンセルを管理するためのコンテキストの使用方法。
- contextを受け入れ、それを使ってゴルーチン、select、およびチャネルを使用してそれ自体をキャンセルする関数の作成方法。
- コールスタックを通じてリクエストスコープのコンテキストを伝播してキャンセルを管理する方法については、Googleのガイドラインに従ってください。
- 必要に応じて、http.ResponseWriterの独自のスパイをロールする方法。

## 16章プロパティベースのテスト
- [string.Builder](https://golang.org/pkg/strings/#Builder)
  - Builderは、Writeメソッドを使用して文字列を効果的に構築するために使用される。メモリのコピーを最小限に抑える。
- プロパティベースのテストは、コードにランダムなデータを投げ、記述したルールが常に正しいことを確認する
  - このテストの本質は、ドメインをよく理解してこれらのプロパティを記述できるようにすること

## 17章 数学
- [Cos](https://golang.org/pkg/math/#Cos)
  - Cosはラジアン引数xの余弦を返す
    - ラジアン： 360度 = 2πラジアン （[rad])
- [不動小数点演算の不正確さ](https://0.30000000000000004.com/)
  - しかし、分割して乗算しないようにするだけで精度を保てる

```go
// NG
numberOfSeconds * π　/ 30

// OK
π / (30 / numberOfSeconds)
```
- [XML構造体の自動生成 zek](https://github.com/miku/zek)
  - [オンラインバージョン](https://www.onlinetool.io/xmltogo/)

- mathTest/clockface内で `go build` を実行しSVGを出力


## 18章 HTTPサーバー
- [ListenAndServe](https://golang.org/pkg/net/http/#ListenAndServe)
-[HandlerFunc](https://golang.org/pkg/net/http/#HandlerFunc)

## 19章 JSON、ルーティング、embedded
- [effective go: embedding](https://golang.org/doc/effective_go.html#embedding)
- JSON文字列のテストをする場合の問題点
  - テストが脆くなる（データモデルを変更するたびにテストが失敗する）
  - デバックが難しい（2つのJSON文字列を比較する時に実際の問題が何なのか理解するのが難しい場合がある）
  - 出力はJSONである必要はあるが、重要なのは出力のエンコード方法ではなく、出力結果である
  - 標準ライブラリの再テストになってしまう。標準ライブラリはすでにテスト済みであり、他の人のコードをテストする必要は無い

- JSONエンコードとデコード
  - この章において、 `Encoder` を作成するには、 `http.ResponseWriter` が作成する `io.Writer` が必要
  - この章において、`Decoder` を作成するには、 レスポンススパイの `Body` フィールドが実装する `io.Writer` が必要