# 問17｜gRPCクライアントの実装

## 目標
gRPCクライアントの接続確立と、スタブ（クライアント用のProxyオブジェクト）を使ったリモートメソッド呼び出しのパターンを学びます。

## 今回の演習で学ぶこと
この問題では、クライアント側からリモートのサーバーメソッドを「まるでローカルの関数」のように呼び出すための作法を、以下の3つの重要ポイントに注目して学びます。

**1. クライアント（スタブ）の生成**
サーバーとやり取りするための「専用の窓口（スタブ）」を生成する工程です。
```go
// 接続済みの通信経路(conn)を渡して、UserService専用の窓口を作る
client := NewUserServiceClient(conn)
```
- `NewUserServiceClient`: `.proto` ファイルから自動生成された強力な関数です。
- これを使うことで、「HTTPリクエストをパースして、ボディにJSONを詰めて…」といった泥臭いネットワーク処理を一切書く必要がなくなり、まるで手元の関数を呼ぶのと同じ感覚で通信できるようになります。

**2. リモートメソッドの呼び出し**
地球の裏側にあるかもしれないサーバーに対して、実際の処理をお願いする工程です。
```go
// サーバーのGetUser関数を呼び出す（ように見える）
resp, err := client.GetUser(ctx, &UserRequest{Id: 1})
if err != nil {
    // 通信エラーやサーバーエラーの処理
    log.Fatalf("失敗しました: %v", err)
}
```
- `client.GetUser`: ここが凄まじいポイントです。ネットワーク越しの呼び出しなのに、完全に型定義されたGoの関数として扱えます。
- もし `UserRequest` に存在しないフィールド（例えば `Age: 20`）を渡そうとすれば、実行前（コードを書いている時点）でコンパイラが「そんな項目はない！」と怒ってくれます。これがgRPCを導入する最大の理由「型破りな開発体験（Type Safety）」です。

**3. 取得データの利用**
サーバーから無事に返ってきた結果を利用する工程です。
```go
// レスポンスのフィールドに直接アクセス！
fmt.Printf("ユーザー名: %s\n", resp.Name)
```
- `resp` はただのGoの構造体（`*User`）として返ってくるため、`resp.Name` のようにドットでいつでも安全にデータへアクセスできます。




**スタブ（Stub）とは？**
「身代わり」や「切り株」という意味です。クライアント側のプログラムにとっては、あたかもリモートサーバーが「手近な関数」として存在しているかのように振る舞ってくれる、便利なプロキシ（代理人）の役割を果たします。

## 問題文
`main.go` の `___` の部分を埋めてプログラムを完成させてください。

`CallGetUser(conn grpcClientConn)` 関数を完成させてください：
1. `NewUserServiceClient(conn)` でクライアントを生成してください。
2. `client.GetUser(context.Background(), &UserRequest{Id: 1})` を呼んでください。
3. 取得した `resp.Name` をコンソールに出力してください。

### 実行方法
```bash
cd /workspaces/Go_Advanced/Week3/Q17
go run main.go
```

### テストでの確認
```bash
cd /workspaces/Go_Advanced/Week3/Q17
go test main.go main_test.go
```
## ヒント
- `NewUserServiceClient(conn)` でスタブを生成 → `client.GetUser(ctx, req)` で呼び出す
- エラー処理は `if err != nil { return }` の形で

## 達成基準
- [ ] gRPCクライアントの接続確立からメソッド呼び出しまでの流れを実装できる
- [ ] スタブ（Stub）の概念を説明できる
- [ ] gRPCのRPC種類（Unary, Streaming等）を説明できる
