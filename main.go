package main

import (
	"github.com/TaskManager/cmd"
	"github.com/TaskManager/database"
	"github.com/TaskManager/utils"
)

func main() {
	utils.Throw(database.NewConnectionPostgres())
	utils.Throw(database.Migration())
	cmd.Execute(database.Db)
}
