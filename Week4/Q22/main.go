package main

import "fmt"

// ======================================
// 1. ドメイン層 (Entities)
// コアとなるデータ構造
// ======================================
type ___ struct {
	ID   string
	Name string
}

// ======================================
// 2. リポジトリのインターフェース
// ユースケースが「外の世界（DB）」とやり取りするためのルールを定めたもの
// 重要なのは「ユースケース層の近くに定義する」ことです
// ======================================
type ___ interface {
	FindByID(id string) (User, error)
}

// ======================================
// 3. ユースケース層 (Use cases)
// ビジネスロジックを記述する
// ======================================
type UserUsecase interface {
	GetUser(id string) (User, error)
}

type userUsecase struct {
	// 具体的なDBの実装ではなく、インターフェースに依存する
	repo ___
}

func NewUserUsecase(r ___) UserUsecase {
	return &userUsecase{repo: r}
}

func (uc *userUsecase) GetUser(id string) (User, error) {
	// ビジネスロジック（例：入力値のチェックや特別な加工）
	if id == "" {
		return User{}, fmt.Errorf("id is required")
	}
	
	// リポジトリにデータ取得を委譲する
	return uc.repo.___(___)
}

// ======================================
// 4. インフラ層 (Infrastructure/Repository Implementation)
// 実際にDBへアクセスするコード（ここではモックの実装）
// ======================================
type mockUserRepository struct{}

func (m *mockUserRepository) FindByID(id string) (User, error) {
	return User{ID: id, Name: "Test User"}, nil
}

// ======================================

func main() {
	// 依存関係を外側から注入して組み上げる（DI）
	repo := &mockUserRepository{}
	usecase := NewUserUsecase(repo)

	user, err := usecase.GetUser("123")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("User struct: %+v\n", user)
}
