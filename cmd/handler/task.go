package handler

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	"github.com/spf13/cobra"

	"github.com/TaskManager/database"
	tasks "github.com/TaskManager/internal"
	UIAddForm "github.com/TaskManager/ui/add-form"
	UIDetailsList "github.com/TaskManager/ui/details-list"
	UIListTable "github.com/TaskManager/ui/list-table"
	UIMessageText "github.com/TaskManager/ui/message-text"
	UIUpdateForm "github.com/TaskManager/ui/update-form"
	"github.com/TaskManager/utils"
)

type FnCobra func(cmd *cobra.Command, args []string)

func AddTask() FnCobra {
	return func(_ *cobra.Command, _ []string) {
		formTask := UIAddForm.Create()
		newTask := tasks.Add(database.Db, formTask)
		UIMessageText.Create(fmt.Sprintf("Task added: %s (#%d)\n", newTask.Name, newTask.ID))
	}
}

func ListTasks() FnCobra {
	return func(_ *cobra.Command, _ []string) {
		allTasks := tasks.GetAll(database.Db)
		if len(allTasks) == 0 {
			UIMessageText.Create(fmt.Sprintf("No tasks found"))
			return
		}
		columns := []table.Column{
			{Title: "ID", Width: 4},
			{Title: "Name", Width: 20},
			{Title: "Description", Width: 30},
			{Title: "Completed", Width: 10},
		}
		rows := make([]table.Row, len(allTasks))

		for i, task := range allTasks {
			rows[i] = table.Row{
				strconv.Itoa(task.ID),
				task.Name,
				task.Description,
				fmt.Sprintf("%t", task.Completed),
			}
		}
		UIListTable.Create(columns, rows)
	}
}

func DetailsTask() FnCobra {
	return func(_ *cobra.Command, args []string) {
		id := utils.Must(strconv.Atoi(args[0]))
		task := tasks.GetById(database.Db, id)
		if task == nil {
			UIMessageText.Create(fmt.Sprintf("Task not found\n"))
			return
		}
		UIDetailsList.Create(*task)
	}
}

func UpdateTask() FnCobra {
	return func(_ *cobra.Command, args []string) {
		id := utils.Must(strconv.Atoi(args[0]))
		task := tasks.GetById(database.Db, id)
		if task == nil {
			UIMessageText.Create(fmt.Sprintf("Task not found\n"))
			return
		}
		taskForm := UIUpdateForm.Create(*task)
		taskUpdated := tasks.UpdateById(database.Db, id, taskForm)
		UIMessageText.Create(fmt.Sprintf("Task updated: %s\n", taskUpdated.Name))
	}
}

func CompletedTask() FnCobra {
	return func(_ *cobra.Command, args []string) {
		id := utils.Must(strconv.Atoi(args[0]))
		task := tasks.GetById(database.Db, id)
		if task == nil {
			UIMessageText.Create(fmt.Sprintf("Task not found\n"))
			return
		}
		task.Completed = true
		taskUpdated := tasks.UpdateById(database.Db, id, *task)
		UIMessageText.Create(fmt.Sprintf("Task completed: %s\n", taskUpdated.Name))
	}
}

func DeleteTask() FnCobra {
	return func(_ *cobra.Command, args []string) {
		id := utils.Must(strconv.Atoi(args[0]))
		task := tasks.GetById(database.Db, id)
		if task == nil {
			UIMessageText.Create(fmt.Sprintf("Task not found\n"))
			return
		}
		tasks.DeleteById(database.Db, id)
		UIMessageText.Create(fmt.Sprintf("Task deleted: %s\n", task.Name))
	}
}
