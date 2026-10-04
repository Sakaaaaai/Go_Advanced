package main

import "fmt"

// コンポーネント群
type Database struct{}
func NewDatabase() *Database { return &Database{} }

type Repository struct{ db *Database }
func NewRepository(db *Database) *Repository { return &Repository{db: db} }

type Service struct{ repo *Repository }
func NewService(r *Repository) *Service { return &Service{repo: r} }
func (s *Service) DoWork() { fmt.Println("Service is working with its injected repository!") }


// 手動DI（またはWireが自動生成するコードのイメージ）
func InitializeService() *___ {
	// 1. 下位の依存から順に生成し、上位コンポーネントに注入していく
	db := ___()
	repo := ___(db)
	service := ___(repo)
	
	return ___
}

/* （参考）実際の Wire 用の定義を書くとしたらこんなイメージになります（今回は実行しません）
// +build wireinject

func InitializeServiceWire() *Service {
    wire.Build(NewDatabase, NewRepository, NewService)
    return &Service{} // wireがここを上で書いたコードで埋めてくれる
}
*/

func main() {
	srv := InitializeService()
	srv.DoWork()
}
