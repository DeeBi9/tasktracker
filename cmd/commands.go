package cmd

import (
	"fmt"
	"strconv"

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

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)
}
