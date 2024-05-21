package main

import "github.com/TaskManager/database"

func main() {
	err := database.NewConnectionPostgres()
	if err != nil {
		return
	}
	errMigration := database.Migration()
	if errMigration != nil {
		return
	}
}
