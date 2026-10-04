package main

import (
	"fmt"
)

// ========= テスト対象のコード =========
type User struct { ID string; Name string }
type UserRepository interface {
	GetUser(id string) (User, error)
}
type UserUsecase struct { repo UserRepository }
func (uc *UserUsecase) FetchUser(id string) (User, error) {
	// ビジネスロジック（そのまま返すだけ）
	return uc.repo.GetUser(id)
}
// ===================================

// 1. テスト用のモック構造体
type MockUserRepository struct {
	// テスト時に好きな振る舞いを注入できるフィールド
	MockGetUser func(id string) (User, error)
}

// 2. モックのメソッド実装
func (m *MockUserRepository) ___(id string) (User, error) {
	// 注入された関数を実行して結果を返す
	return m.___(___)
}

func main() {
	// --------- テストコードのシミュレーション ---------

	// 3. モックの振る舞いを定義する
	mockRepo := &MockUserRepository{
		MockGetUser: func(id string) (User, error) {
			if id == "1" {
				return User{ID: "1", Name: "Alice"}, nil
			}
			return User{}, fmt.Errorf("user not found")
		},
	}

	// 4. モックをユースケースに注入（DI）してテスト実行
	usecase := &UserUsecase{repo: ___}

	// 正常系のテスト
	user, err := usecase.FetchUser("1")
	fmt.Printf("Test ID '1': user=%+v, err=%v\n", user, err)

	// 異常系のテスト
	user2, err2 := usecase.FetchUser("99")
	fmt.Printf("Test ID '99': user=%+v, err=%v\n", user2, err2)
}
