package main

import "fmt"

// GraphQLのスキーマ定義（.graphqls ファイルの中身のイメージ）
const GraphQLSchema = `
# 1. User型の定義
___ User {
  id: ___!
  name: ___!
  // ageはオプショナル（!をつけない）
  age: ___
}

# 2. Query型の定義（データの取得用）
___ Query {
  # 3. 引数に 必須のid を取り、User型を返すクエリ
  ___(id: ___!): ___
}
`

func main() {
	fmt.Println("=== GraphQL Schema Definition ===")
	fmt.Println(GraphQLSchema)
	fmt.Println("=================================")
	fmt.Println("実際の開発では、このスキーマファイルをもとに 'gqlgen' を使ってGoのコードを生成します。")
}
