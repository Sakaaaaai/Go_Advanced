//go:build ignore

package main

import "fmt"

type Database struct{}
func NewDatabase() *Database { return &Database{} }

type Repository struct{ db *Database }
func NewRepository(db *Database) *Repository { return &Repository{db: db} }

type Service struct{ repo *Repository }
func NewService(r *Repository) *Service { return &Service{repo: r} }
func (s *Service) DoWork() { fmt.Println("Service is working with its injected repository!") }


func InitializeService() *Service {
	db := NewDatabase()
	repo := NewRepository(db)
	service := NewService(repo)
	
	return service
}

func main() {
	srv := InitializeService()
	srv.DoWork()
}
