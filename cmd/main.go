package main

import (
	"log"

	"github.com/asaberwd/users-api/api"
	storage "github.com/asaberwd/users-api/internal/db"
	"github.com/asaberwd/users-api/internal/user"
	"github.com/boltdb/bolt"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db, err := bolt.Open("users.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	boltDB := getBoltDB(db, storage.UserBucket)
	userManager := user.NewManager(boltDB, nil)
	userHandler := api.NewUserHandler(*userManager)
	api.Router(e, userHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

func getBoltDB(db *bolt.DB, bucket storage.Bucket) *storage.Bolt {
	bolt, err := storage.NewBoltDB(db, bucket)
	if err != nil {
		log.Fatal(err)
	}
	return bolt
}
