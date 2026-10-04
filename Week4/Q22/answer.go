//go:build ignore

package main

import "fmt"

type User struct {
	ID   string
	Name string
}

type UserRepository interface {
	FindByID(id string) (User, error)
}

type UserUsecase interface {
	GetUser(id string) (User, error)
}

type userUsecase struct {
	repo UserRepository
}

func NewUserUsecase(r UserRepository) UserUsecase {
	return &userUsecase{repo: r}
}

func (uc *userUsecase) GetUser(id string) (User, error) {
	if id == "" {
		return User{}, fmt.Errorf("id is required")
	}
	
	return uc.repo.FindByID(id)
}

type mockUserRepository struct{}

func (m *mockUserRepository) FindByID(id string) (User, error) {
	return User{ID: id, Name: "Test User"}, nil
}

func main() {
	repo := &mockUserRepository{}
	usecase := NewUserUsecase(repo)

	user, err := usecase.GetUser("123")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("User struct: %+v\n", user)
}
