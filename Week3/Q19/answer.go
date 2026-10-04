//go:build ignore

package main

import "fmt"

const GraphQLSchema = `
type User {
  id: ID!
  name: String!
  age: Int
}

type Query {
  user(id: ID!): User
}
`

func main() {
	fmt.Println("=== GraphQL Schema Definition ===")
	fmt.Print(GraphQLSchema)
	fmt.Println("=================================")
	fmt.Println("実際の開発では、このスキーマファイルをもとに 'gqlgen' を使ってGoのコードを生成します。")
}
