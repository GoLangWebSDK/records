package main

import (
	"fmt"

	"github.com/GoLangWebSDK/crud/database"
	"github.com/GoLangWebSDK/crud/database/adapters"
	"github.com/GoLangWebSDK/crud/orms/gorm"
)

type User struct {
	ID        uint32
	FirstName string
	LastName  string
}

func main() {
	adapter := adapters.NewSQLite(database.WithDSN("sqlite.db"))
	db := database.New(adapter)

	migrations := gorm.NewGormMigrator(db)
	migrations.Run()

	users := gorm.NewRepository[User](db, User{})

	allUsers, err := users.All()

	if err != nil {
		fmt.Println(err)
	}

	for _, user := range allUsers {
		fmt.Println(user)
	}
}
	