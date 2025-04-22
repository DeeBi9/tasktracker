package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/DeeBi9/tasktracker/tasks"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to the JSON file",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		err := tasks.AddJSON(args[0], args[1])
		if err != nil {
			fmt.Printf("Error adding task: %v\n", err)
			return
		}
		fmt.Println("Task added successfully.")
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing task in the JSON file",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		arg2, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: Second argument must be an integer")
			return
		}

		err = tasks.UpdateJSON(args[0], arg2, args[2])
		if err != nil {
			fmt.Printf("Error updating task: %v\n", err)
			return
		}
		fmt.Println("Task updated successfully.")
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task from the JSON file",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		arg2, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: Second argument must be an integer")
			return
		}

		err = tasks.DeleteJSON(args[0], arg2)
		if err != nil {
			fmt.Printf("Error deleting task: %v\n", err)
			return
		}
		fmt.Println("Task deleted successfully.")
	},
}

var taskCmd_ = &cobra.Command{
	Use:   "changestatus",
	Short: "Change the status of task to to-do",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		arg1, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: Second argument must be an integer")
			return
		}
		err = tasks.Change_status(args[0], arg1)
		if err != nil {
			fmt.Printf("Error changing status: %v\n", err)
			return
		}
		fmt.Println("Status changed successfully.")
	},
}

var listAll = &cobra.Command{
	Use:   "listall",
	Short: "List all the tasks",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var data tasks.Data
		data, err = tasks.ListAll()
		if err != nil {
			fmt.Printf("Error Listing all tasks")
			return
		}
		fmt.Println("Tasks are :")
		for _, task := range data.Data {
			fmt.Printf("ID: %d\nDescription: %s\nStatus: %s\nCreated: %s\nUpdated: %s\n\n",
				task.Id, task.Description, task.Status, task.CreatedAt.Format(time.RFC822), task.UpdatedAt.Format(time.RFC822))
			fmt.Printf("\n")
		}
	},
}

var listStatusWise = &cobra.Command{
	Use:   "list",
	Short: "List all the tasks",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var data tasks.Data
		data, err = tasks.Liststatuswise(args[0])
		if err != nil {
			fmt.Printf("Error Listing all tasks")
			return
		}
		fmt.Printf("Tasks of %v\n :", args[0])
		for _, task := range data.Data {
			fmt.Printf("ID: %d\nDescription: %s\nStatus: %s\nCreated: %s\nUpdated: %s\n\n",
				task.Id, task.Description, task.Status, task.CreatedAt.Format(time.RFC822), task.UpdatedAt.Format(time.RFC822))
			fmt.Printf("\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(taskCmd_)
	rootCmd.AddCommand(listAll)
	rootCmd.AddCommand(listStatusWise)
}
