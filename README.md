# Go言語 上級者向けコース 演習リポジトリ

[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://codespaces.new/Sakaaaaai/Go_Advanced?quickstart=1)

上のボタンを押すと、ブラウザ上に Go 1.24 と VS Code の Go 拡張が入った演習環境が立ち上がります。ローカルへのインストールは不要です。

## 使い方

1. 上の「Open in GitHub Codespaces」ボタンを押し、「Create codespace」を選びます（初回は1〜2分かかります）。
2. 左のエクスプローラーから課題フォルダ（例：`Week1/Q1`）を開き、`README.md` を読みます。
3. `main.go` の `___` や TODO を埋めます。
4. 下部のターミナルで課題フォルダに移動し、実行・テストします。

```bash
cd /workspaces/Go_Advanced/Week1/Q1
go run main.go
go test main.go main_test.go
```

すべて `PASS`（`ok`）になればクリアです。

## 解答例について

各フォルダの `answer.go` が解答例です。`main.go` と同時にビルドされないよう `//go:build ignore` を付けてあります。解答例の動作は次のように確認できます。

```bash
go run answer.go
go test answer.go main_test.go
```

## 構成

| 週 | テーマ | 問題 |
|---|---|---|
| Week1 | 高度な並行処理とパフォーマンス・チューニング | Q1〜Q7 |
| Week2 | データベース応用とORM・キャッシュ | Q8〜Q14 |
| Week3 | gRPC・API設計・メッセージキュー | Q15〜Q21 |
| Week4 | アーキテクチャ設計・テスト・CI/CD・総合演習 | Q22〜Q30 |

詳しい学習の進め方は [guide.md](guide.md) を参照してください。

## 注意

- 演習はすべて Go 標準ライブラリだけで動きます（Redis・DB・gRPC はモックで再現しています）。外部サービスの起動は不要です。
- Codespaces は無料枠を超えると課金されます。使い終わったら codespace を停止してください。
