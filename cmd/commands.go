package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"gorm.io/gorm"

	tasks "github.com/TaskManager/internal"
	"github.com/TaskManager/utils"
)

func Execute(db *gorm.DB) {
	var cmdAdd = &cobra.Command{
		Use:   "add [string]",
		Short: "Add a new task to your task list",
		Long:  "Add a new task to your task list. The task will be marked as not completed by default.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var task = tasks.MTask{
				Name: strings.Join(args, " "),
			}
			newTask := tasks.Add(db, task)
			fmt.Printf("▓ Task added: %s (#%d)\n", newTask.Name, newTask.ID)
		},
	}
	var cmdList = &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		Long:  "List all tasks in your task list.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("▓ Tasks list")
			allTasks := tasks.GetAll(db)
			if len(allTasks) == 0 {
				fmt.Println("▓ No tasks found")
				return
			}
			for i, task := range allTasks {
				ascii := "├"
				if i == 0 {
					ascii = "┌"
				} else if i == len(allTasks)-1 {
					ascii = "└"
				}
				if len(allTasks) == 1 {
					ascii = "-"
				}
				fmt.Printf("%s (#%d) %s [%t]\n", ascii, task.ID, task.Name, task.Completed)
			}
		},
	}
	var cmdDetails = &cobra.Command{
		Use:   "details [id]",
		Short: "Show details of a task",
		Long:  "Show details of a task by its ID.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := utils.Must(strconv.Atoi(args[0]))
			task := tasks.GetById(db, id)
			if task == nil {
				fmt.Println("▓ Task not found")
				return
			}
			fmt.Printf("▓ Task details: (#%d) \n- Name: %s\n- Description: %s\n- Estado: %t\n- Created: %s",
				task.ID, task.Name, task.Description, task.Completed, task.CreatedAt)
		},
	}
	var cmdUpdate = &cobra.Command{
		Use:   "update [id] [string] [string]",
		Short: "Update a task",
		Long:  "Update a task by its ID. You can update the name and description.",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			id := utils.Must(strconv.Atoi(args[0]))
			taskToEdit := tasks.GetById(db, id)
			if taskToEdit == nil {
				fmt.Println("▓ Task not found")
				return
			}
			taskToEdit.Name = args[1]
			taskToEdit.Description = args[2]
			taskUpdated := tasks.UpdateById(db, id, *taskToEdit)
			fmt.Printf("▓ Task updated: %s\n", taskUpdated.Name)
		},
	}
	var cmdCompleted = &cobra.Command{
		Use:   "completed [id]",
		Short: "Mark a task as completed",
		Long:  "Mark a task as completed by its ID.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := utils.Must(strconv.Atoi(args[0]))
			task := tasks.GetById(db, id)
			if task == nil {
				fmt.Println("▓ Task not found")
				return
			}
			task.Completed = true
			taskUpdated := tasks.UpdateById(db, id, *task)
			fmt.Printf("▓ Task completed: %s\n", taskUpdated.Name)
		},
	}
	var cmdDelete = &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete a task",
		Long:  "Delete a task by its ID.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := utils.Must(strconv.Atoi(args[0]))
			task := tasks.GetById(db, id)
			if task == nil {
				fmt.Println("▓ Task not found")
				return
			}
			tasks.DeleteById(db, id)
			fmt.Printf("▓ Task deleted: %s\n", task.Name)
		},
	}
	var rootCmd = &cobra.Command{Use: "task"}
	rootCmd.AddCommand(
		cmdAdd,
		cmdList,
		cmdDetails,
		cmdUpdate,
		cmdCompleted,
		cmdDelete,
	)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
