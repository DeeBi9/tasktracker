package cmd

import (
	"fmt"
	"strconv"

	"github.com/DeeBi9/tasktracker/tasks"
	"github.com/spf13/cobra"
)

var add = &cobra.Command{
	Use:   "Add",
	Short: "Add JSON entry",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Entry added", tasks.AddJSON(args[0], args[1]))
	},
}

var update = &cobra.Command{
	Use:   "Update",
	Short: "Update JSON entry",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
			fmt.Println("Error: Expected at least 3 arguments")
			return
		}

		arg1 := args[0]
		arg2Str := args[1]
		arg3 := args[2]

		arg2, err := strconv.Atoi(arg2Str)
		if err != nil {
			fmt.Println("Error: Second argument must be an integer")
			return
		}
		fmt.Printf("Entry updated", tasks.UpdateJSON(arg1, arg2, arg3))
	},
}

var delete = &cobra.Command{
	Use:   "Delete",
	Short: "Delete JSON entry",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		arg1 := args[0]
		arg2Str := args[1]

		arg2, err := strconv.Atoi(arg2Str)
		if err != nil {
			fmt.Println("Error: Second argument must be an integer")
			return
		}
		fmt.Printf("Entry Deleted", tasks.DeleteJSON(arg1, arg2))
	},
}

func init() {
	rootCmd.AddCommand(add)
	rootCmd.AddCommand(update)
	rootCmd.AddCommand(delete)
}
