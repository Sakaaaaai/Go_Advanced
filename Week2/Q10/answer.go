//go:build ignore

package main

import "fmt"

type gormDB struct{}
type gormModel struct {
	ID        uint
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

func (db *gormDB) Create(value interface{}) *gormDB { return db }
func (db *gormDB) Where(query interface{}, args ...interface{}) *gormDB { return db }
func (db *gormDB) First(dest interface{}, conds ...interface{}) *gormDB { return db }
func (db *gormDB) Error() error { return nil }

type User struct {
	gormModel
	Name string
	Age  int
}

func CreateUser(db *gormDB, name string, age int) error {
	user := User{Name: name, Age: age}
	result := db.Create(&user)
	return result.Error()
}

func GetUserByName(db *gormDB, name string) (User, error) {
	var user User
	result := db.Where("name = ?", name).First(&user)
	return user, result.Error()
}

func main() {
	fmt.Println("This is a package for GORM logic.")
}
