//go:build ignore

package main

import (
	"fmt"
)

type User struct { ID string; Name string }
type UserRepository interface {
	GetUser(id string) (User, error)
}
type UserUsecase struct { repo UserRepository }
func (uc *UserUsecase) FetchUser(id string) (User, error) {
	return uc.repo.GetUser(id)
}

type MockUserRepository struct {
	MockGetUser func(id string) (User, error)
}

func (m *MockUserRepository) GetUser(id string) (User, error) {
	return m.MockGetUser(id)
}

func main() {
	mockRepo := &MockUserRepository{
		MockGetUser: func(id string) (User, error) {
			if id == "1" {
				return User{ID: "1", Name: "Alice"}, nil
			}
			return User{}, fmt.Errorf("user not found")
		},
	}

	usecase := &UserUsecase{repo: mockRepo}

	user, err := usecase.FetchUser("1")
	fmt.Printf("Test ID '1': user=%+v, err=%v\n", user, err)

	user2, err2 := usecase.FetchUser("99")
	fmt.Printf("Test ID '99': user=%+v, err=%v\n", user2, err2)
}
