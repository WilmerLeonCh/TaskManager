package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	cmd "github.com/TaskManager/cmd/handler"
)

func Execute() {
	var cmdAdd = &cobra.Command{
		Use:   "add",
		Short: "Add a new task to your task list",
		Long:  "Add a new task to your task list. The task will be marked as not completed by default.",
		Run:   cmd.AddTask(),
	}
	var cmdList = &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		Long:  "List all tasks in your task list.",
		Run:   cmd.ListTasks(),
	}
	var cmdDetails = &cobra.Command{
		Use:   "details [id]",
		Short: "Show details of a task",
		Long:  "Show details of a task by its ID.",
		Args:  cobra.ExactArgs(1),
		Run:   cmd.DetailsTask(),
	}
	var cmdUpdate = &cobra.Command{
		Use:   "update [id]",
		Short: "Update a task",
		Long:  "Update a task by its ID. You can update the name and description.",
		Args:  cobra.ExactArgs(1),
		Run:   cmd.UpdateTask(),
	}
	var cmdCompleted = &cobra.Command{
		Use:   "completed [id]",
		Short: "Mark a task as completed",
		Long:  "Mark a task as completed by its ID.",
		Args:  cobra.ExactArgs(1),
		Run:   cmd.CompletedTask(),
	}
	var cmdDelete = &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete a task",
		Long:  "Delete a task by its ID.",
		Args:  cobra.ExactArgs(1),
		Run:   cmd.DeleteTask(),
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
