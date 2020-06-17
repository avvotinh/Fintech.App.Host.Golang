package migrations

import (
	"github.com/avvotinh/Fintech.App.Host.Golang/helpers"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance uint
	UserID  uint
}

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=fintechapp password=postgres sslmode=disable")
	helpers.HandleError(err)

	return db
}

func CreateAccounts() {
	db := ConnectDB()

	users := [2]User{
		{Username: "Hople", Email: "test01@gmail.com"},
		{Username: "Rumdev", Email: "test02@gmail.com"},
	}

	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

		account := Account{Type: "Daily Account", Name: string(users[i].Username + "s" + " account"), Balance: uint(10000 * int(i+1)), UserID: user.ID}
		db.Create(&account)
	}

	defer db.Close()
}

func Migtate() {
	db := ConnectDB()
	db.AutoMigrate(&User{}, &Account{})
	defer db.Close()

	CreateAccounts()
}
