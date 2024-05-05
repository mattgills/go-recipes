package main

import (
	"recipe-manager/api"
	"recipe-manager/db"

	_ "modernc.org/sqlite"
)

func main() {
	db.Init()
	api.Init()
}
