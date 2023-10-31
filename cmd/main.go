package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "This is an example app",
	Long:  `This is a longer description of the example app`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		fmt.Printf("Hello, %s!\n", name)
	},
}

func main() {
	rootCmd.Flags().StringP("name", "n", "World", "a name to say hello to")
	rootCmd.Execute()
}
