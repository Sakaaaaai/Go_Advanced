# 問15｜Protocol Buffers（Protobuf）基礎

## 目標
gRPCの基盤となるデータシリアライズフォーマット「Protocol Buffers（Protobuf）」の `.proto` 定義ファイルの書き方を学びます。

## 今回の演習で学ぶこと
この問題では、gRPCの通信で使われるデータ定義フォーマット「Protocol Buffers」の `.proto` ファイルの書き方を、以下の3つの重要ポイントに注目して学びます。

**1. シンタックス（バージョン）の指定**
Protobufにはいくつかバージョンがありますが、現在は「プロト3（proto3）」が世界的なスタンダードです。
```protobuf
syntax = "proto3";
```
- この1行をファイルの冒頭に書くことで、「このファイルは最新のproto3のルールで解釈してください」とコンパイラに宣言します。
- これを書き忘れると古い proto2 として扱われ、デフォルト値の挙動などが変わって予期せぬバグを生む原因になります。

**2. メッセージ（データ構造）とフィールド番号**
サーバーとクライアント間でやり取りする「データの設計図（型）」を定義する工程です。
```protobuf
message User {
    int32 id = 1;      // 【超重要】1番の荷物は「ID」
    string name = 2;   // 【超重要】2番の荷物は「名前」
}
```
- `message`: Go言語の `struct`（構造体）やJSONのオブジェクトに相当するものです。
- `= 1`, `= 2`: これは代入ではなく「フィールド番号」と呼ばれる**絶対不変の背番号**です。Protobufはデータを極限まで圧縮して送るため、「`name`という文字」ではなく「背番号`2`」として送受信します。この番号を後から勝手に変えると、通信データが完全に壊れて大事故に繋がります。

**3. サービス（API）の定義**
どんな機能（関数）を外部に公開するのかを取り決める「メニュー表」を作る工程です。
```protobuf
service UserService {
    rpc GetUser (UserRequest) returns (User);
}
```
- `service`: `UserService` という、機能のまとまり（インターフェース）を定義します。
- `rpc`: 「Remote Procedure Call（遠隔手続き呼び出し）」の略で、サーバー側にある `GetUser` 関数を定義しています。「`UserRequest` 型のデータを受け取って、`User` 型のデータを返す」という厳格な契約をここで結びます。




**なぜ Protobuf なのか？**
JSONなどのテキストデータに比べ、Protobufはバイナリ形式でデータを送るため、サイズが圧倒的に小さく、通信速度が非常に高速です。マイクロサービス間の通信や、高パフォーマンスが求められるシステムでは標準的な技術となっています。

## 問題文
定数 `ProtoFileContent` に以下の内容の `.proto` ファイルの文字列を完成させてください：

1. `syntax = "proto3"` を記述してください。
2. `id(int32, field 1)` と `name(string, field 2)` を持つ `User` メッセージを定義してください。
3. `GetUser(UserRequest) → User` の RPC を持つ `UserService` サービスを定義してください。

### 実行方法
```bash
cd /workspaces/Go_Advanced/Week3/Q15
go run main.go
```

### テストでの確認
```bash
cd /workspaces/Go_Advanced/Week3/Q15
go test main.go main_test.go
```
## ヒント
- `message User { int32 id = 1; string name = 2; }` の形
- `service UserService { rpc GetUser (UserRequest) returns (User); }` の形

## 達成基準
- [ ] Protobuf と JSON の主な違いを説明できる
- [ ] `message` と `service` の `.proto` 定義が書ける
- [ ] フィールド番号の役割を説明できる
