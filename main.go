package main

import "github.com/TaskManager/database"

func main() {
	err := database.NewConnectionPostgres("host=localhost user=r00t password=passw0rd dbname=task_manager_db port=5410 sslmode=disable TimeZone=America/Lima")
	if err != nil {
		return
	}
}
